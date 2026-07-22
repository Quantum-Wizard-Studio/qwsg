package collector

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/inventory"
)

func TestRegistryRejectsDuplicates(t *testing.T) {
	r := NewRegistry()
	c := testCollector("one", nil)
	if err := r.Register(c); err != nil {
		t.Fatal(err)
	}
	if err := r.Register(c); !errors.Is(err, ErrDuplicateCollector) {
		t.Fatalf("got %v", err)
	}
}

func TestRegistryExecutesDeterministically(t *testing.T) {
	r := NewRegistry()
	for _, name := range []string{"zeta", "alpha", "middle"} {
		if err := r.Register(testCollector(name, nil)); err != nil {
			t.Fatal(err)
		}
	}
	want := []string{"alpha", "middle", "zeta"}
	if got := r.Names(); !reflect.DeepEqual(got, want) {
		t.Fatalf("names %v, want %v", got, want)
	}
	results := r.Execute(context.Background(), "request")
	got := make([]string, 0, len(results))
	for _, result := range results {
		got = append(got, result.CollectorName)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("execution %v, want %v", got, want)
	}
}

func TestDependencyAndAvailabilityChecks(t *testing.T) {
	r := NewRegistry()
	unavailable := testCollector("alpha", nil)
	unavailable.Check = func(context.Context) Availability {
		return Availability{Available: false, Code: "unsupported", MessageKey: "collector.unsupported"}
	}
	dependent := testCollector("beta", nil)
	dependent.Dependencies = []string{"alpha"}
	if err := r.Register(unavailable); err != nil {
		t.Fatal(err)
	}
	if err := r.Register(dependent); err != nil {
		t.Fatal(err)
	}
	results := r.Execute(context.Background(), "request")
	if results[0].HealthStatus != inventory.Unsupported || results[1].HealthStatus != inventory.Unavailable || results[1].Metadata["dependency"] != "alpha" {
		t.Fatalf("unexpected results: %#v", results)
	}
}

func TestDependencyRunsBeforeAlphabeticallyEarlierDependent(t *testing.T) {
	r := NewRegistry()
	dependent := testCollector("alpha", nil)
	dependent.Dependencies = []string{"zeta"}
	for _, c := range []Collector{dependent, testCollector("zeta", nil)} {
		if err := r.Register(c); err != nil {
			t.Fatal(err)
		}
	}
	results := r.Execute(context.Background(), "request")
	if results[0].CollectorName != "zeta" || results[1].CollectorName != "alpha" || results[1].HealthStatus != inventory.Available {
		t.Fatalf("unexpected dependency plan: %#v", results)
	}
}

func TestTimeoutCancellationAndPanicAreIsolated(t *testing.T) {
	cases := []struct {
		name string
		ctx  func() (context.Context, context.CancelFunc)
		run  func(context.Context) ([]inventory.Item, []string, error)
		want inventory.CategoryStatus
	}{
		{name: "timeout", ctx: func() (context.Context, context.CancelFunc) { return context.Background(), func() {} }, run: func(ctx context.Context) ([]inventory.Item, []string, error) {
			<-ctx.Done()
			return nil, nil, ctx.Err()
		}, want: inventory.Timeout},
		{name: "cancelled", ctx: func() (context.Context, context.CancelFunc) { return context.WithCancel(context.Background()) }, run: func(ctx context.Context) ([]inventory.Item, []string, error) {
			<-ctx.Done()
			return nil, nil, ctx.Err()
		}, want: inventory.Cancelled},
		{name: "panic", ctx: func() (context.Context, context.CancelFunc) { return context.Background(), func() {} }, run: func(context.Context) ([]inventory.Item, []string, error) { panic("boom") }, want: inventory.Error},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := testCollector(tc.name, tc.run)
			if tc.name == "timeout" {
				c.Timeout = time.Millisecond
			}
			r := NewRegistry()
			if err := r.Register(c); err != nil {
				t.Fatal(err)
			}
			ctx, cancel := tc.ctx()
			if tc.name == "cancelled" {
				cancel()
			} else {
				defer cancel()
			}
			results := r.Execute(ctx, "request")
			if results[0].HealthStatus != tc.want {
				t.Fatalf("got %s, want %s", results[0].HealthStatus, tc.want)
			}
		})
	}
}

func TestOutputLimitIsEnforced(t *testing.T) {
	c := limitedCollector{Func: testCollector("limited", func(context.Context) ([]inventory.Item, []string, error) {
		return []inventory.Item{{ID: "large", Kind: "test", Facts: map[string]inventory.Fact{"value": {Value: string(make([]byte, 1024)), Quality: "observed", Sensitivity: "operational", Provenance: inventory.Provenance{SourceType: "fixture"}}}}}, []string{"fixture"}, nil
	}), limit: 128}
	r := NewRegistry()
	if err := r.Register(c); err != nil {
		t.Fatal(err)
	}
	result := r.Execute(context.Background(), "request")[0]
	if result.HealthStatus != inventory.Error || len(result.Errors) != 1 || result.Errors[0].Code != "resource_limit" {
		t.Fatalf("output limit not enforced: %#v", result)
	}
}

type limitedCollector struct {
	Func
	limit int64
}

func (c limitedCollector) Descriptor() Descriptor {
	d := c.Func.Descriptor()
	d.OutputLimitBytes = c.limit
	return d
}

func testCollector(name string, run func(context.Context) ([]inventory.Item, []string, error)) Func {
	if run == nil {
		run = func(context.Context) ([]inventory.Item, []string, error) {
			return []inventory.Item{{ID: name, Kind: "test", Facts: map[string]inventory.Fact{}}}, []string{"fixture"}, nil
		}
	}
	return Func{Name: name, Run: run}
}
