package releasediscovery

import (
	"context"
	"sort"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/update"
)

type Discoverer struct {
	source    ReleaseSource
	verifier  Verifier
	evaluator Evaluator
}

func NewDiscoverer(source ReleaseSource, verifier Verifier, evaluator Evaluator) (Discoverer, error) {
	if source == nil || evaluator.classify == nil {
		return Discoverer{}, fail(SourceAuthority)
	}
	return Discoverer{source: source, verifier: verifier, evaluator: evaluator}, nil
}

type CheckResult struct {
	NotModified       bool
	Source            SourceEvidence
	IndexGeneratedAt  string
	WithdrawnVersions []string
	Authenticity      AuthenticityEvidence
	Evaluation        Evaluation
}

func (d Discoverer) Check(ctx context.Context, request FetchRequest, platform string, allowPrerelease bool) (CheckResult, error) {
	fetched, err := d.source.Fetch(ctx, request)
	if err != nil {
		return CheckResult{}, err
	}
	if fetched.NotModified {
		return CheckResult{NotModified: true, Source: fetched.Evidence}, nil
	}
	index, err := Parse(fetched.Manifest)
	if err != nil {
		return CheckResult{}, err
	}
	authenticated, err := d.verifier.Verify(index)
	if err != nil {
		return CheckResult{}, err
	}
	result := CheckResult{Source: fetched.Evidence, IndexGeneratedAt: index.GeneratedAt, Authenticity: authenticated.Authenticity()}
	for _, channel := range index.Channels {
		if channel.Name != request.Channel {
			continue
		}
		for _, release := range channel.Releases {
			if release.Status == "withdrawn" {
				for _, artifact := range release.Artifacts {
					if artifact.Platform == platform {
						result.WithdrawnVersions = append(result.WithdrawnVersions, release.Version)
					}
				}
			}
		}
	}
	evaluation, err := d.evaluator.Evaluate(authenticated, request.Channel, platform, allowPrerelease)
	if err != nil {
		return result, err
	}
	result.Evaluation = evaluation
	return result, nil
}

type InstalledClassifier func(candidateVersion string) installation.Result

type Evaluator struct{ classify InstalledClassifier }

func NewEvaluator(classifier InstalledClassifier) (Evaluator, error) {
	if classifier == nil {
		return Evaluator{}, fail(InstalledUnverified)
	}
	return Evaluator{classify: classifier}, nil
}

func LocalEvaluator() Evaluator {
	evaluator, _ := NewEvaluator(func(candidateVersion string) installation.Result {
		return installation.Classify(installation.Options{Root: "/", CandidateVersion: candidateVersion})
	})
	return evaluator
}

type Compatibility string

const (
	CompatibilityNotApplicable Compatibility = "not_applicable"
	CompatibilitySupported     Compatibility = "supported"
	CompatibilityUnsupported   Compatibility = "unsupported"
)

type Evaluation struct {
	InstalledVersion string
	Channel          string
	Platform         string
	Release          Release
	Artifact         Artifact
	Relation         update.Relation
	Compatibility    Compatibility
	MigrationID      string
	Authenticity     AuthenticityEvidence
}

func (e Evaluator) Evaluate(authenticated AuthenticatedIndex, channelName, platform string, allowPrerelease bool) (Evaluation, error) {
	installed := e.classify("")
	installedParsed, installedErr := update.ParseVersion(installed.Version)
	if installed.State != installation.VerifiedSupported || installedErr != nil || installedParsed.Major != 1 {
		return Evaluation{}, fail(InstalledUnverified)
	}
	index := authenticated.Index()
	var selected *Channel
	for position := range index.Channels {
		if index.Channels[position].Name == channelName {
			selected = &index.Channels[position]
			break
		}
	}
	if selected == nil || !validChannel(channelName) || platform != "linux-amd64" {
		return Evaluation{}, fail(NoEligibleRelease)
	}
	type candidate struct {
		release  Release
		artifact Artifact
		version  update.Version
	}
	candidates := make([]candidate, 0, len(selected.Releases))
	for _, release := range selected.Releases {
		parsed, _ := update.ParseVersion(release.Version)
		if release.Status != "active" || (!allowPrerelease && len(parsed.Prerelease) != 0) {
			continue
		}
		for _, artifact := range release.Artifacts {
			if artifact.Platform == platform {
				candidates = append(candidates, candidate{release: release, artifact: artifact, version: parsed})
			}
		}
	}
	if len(candidates) == 0 {
		return Evaluation{}, fail(NoEligibleRelease)
	}
	sort.Slice(candidates, func(i, j int) bool { return update.Compare(candidates[i].version, candidates[j].version) > 0 })
	best := candidates[0]
	result := Evaluation{InstalledVersion: installed.Version, Channel: channelName, Platform: platform, Release: best.release, Artifact: best.artifact, Relation: update.Classify(installed.Version, best.release.Version), Compatibility: CompatibilityNotApplicable, Authenticity: authenticated.Authenticity()}
	if result.Relation != update.Newer {
		return result, nil
	}
	result.Compatibility = CompatibilityUnsupported
	minimum, _ := update.ParseVersion(best.release.MinimumSourceVersion)
	installedVersion, _ := update.ParseVersion(installed.Version)
	if update.Compare(installedVersion, minimum) < 0 {
		return result, nil
	}
	upgrade := e.classify(best.release.Version)
	localPlan, planErr := update.PlanMigration(installed.Version, best.release.Version)
	if upgrade.State != installation.SupportedUpgradeSource || upgrade.Version != installed.Version || planErr != nil || localPlan.Validate() != nil || upgrade.MigrationID != localPlan.ID || !contains(best.release.MigrationRoutes, localPlan.ID) {
		return result, nil
	}
	result.Compatibility = CompatibilitySupported
	result.MigrationID = localPlan.ID
	return result, nil
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
