package database

import (
	"fmt"
	"os"
	"osv-detector/detector"
	"osv-detector/detector/parsers"
	"osv-detector/detector/semver"
	"regexp"
	"strings"
	"time"
)

type AffectsRangeType string

const (
	TypeSemver    AffectsRangeType = "SEMVER"
	TypeEcosystem AffectsRangeType = "ECOSYSTEM"
	TypeGit       AffectsRangeType = "GIT"
)

type Ecosystem = detector.Ecosystem

type Package struct {
	Name      string    `json:"name"`
	Ecosystem Ecosystem `json:"ecosystem"`
}

// NormalizedName ensures that the package name is normalized based on ecosystem
// in accordance to the OSV specification.
//
// This is required because currently both GitHub and Pip seem to be a bit
// inconsistent in their package name handling, so we normalize them
// to be on the safe side.
//
// In the future, it's hoped that this can be improved.
func (p Package) NormalizedName() string {
	if p.Ecosystem != parsers.PipEcosystem {
		return p.Name
	}

	// per https://www.python.org/dev/peps/pep-0503/#normalized-names
	name := regexp.MustCompile(`[-_.]+`).ReplaceAllString(p.Name, "-")

	return strings.ToLower(name)
}

type RangeEvent struct {
	Introduced string `json:"introduced,omitempty"`
	Fixed      string `json:"fixed,omitempty"`
}

type AffectsRange struct {
	Type   AffectsRangeType `json:"type"`
	Events []RangeEvent     `json:"events"`
}

func (ar AffectsRange) containsEcosystem(v string) bool {
	if ar.Type != TypeEcosystem {
		return false
	}
	// todo: we should probably warn here
	if len(ar.Events) == 0 {
		return false
	}

	vp := semver.Parse(v)

	var affected bool
	for _, e := range ar.Events {
		if !affected && e.Introduced != "" {
			affected = e.Introduced == "0" || vp.CompareStr(e.Introduced) >= 0
		} else if affected && e.Fixed != "" {
			affected = vp.CompareStr(e.Fixed) < 0
		}
	}

	return affected
}

type Affects []AffectsRange

// AffectsEcosystem checks if the given version is within the range
// specified by the events of any "Ecosystem" type ranges
func (a Affects) AffectsEcosystem(v string) bool {
	for _, r := range a {
		if r.Type != TypeEcosystem {
			continue
		}
		if r.containsEcosystem(v) {
			return true
		}
	}

	return false
}

type Reference struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Affected struct {
	Package Package `json:"package"`
	Ranges  Affects `json:"ranges,omitempty"`
}

// OSV represents an OSV style JSON vulnerability database entry
type OSV struct {
	ID        string     `json:"id"`
	Summary   string     `json:"summary"`
	Published time.Time  `json:"published"`
	Modified  time.Time  `json:"modified"`
	Withdrawn *time.Time `json:"withdrawn,omitempty"`
	Details   string     `json:"details"`
	Affected  []Affected `json:"affected"`
}

func (osv *OSV) AffectsEcosystem(ecosystem detector.Ecosystem) bool {
	if osv.Affected == nil {
		fmt.Printf("Ignoring %s as it does not have an 'affected' property\n", osv.ID)

		return false
	}

	for _, affected := range osv.Affected {
		if affected.Package.Ecosystem == ecosystem {
			return true
		}
	}

	return false
}

func (osv *OSV) Link() string {
	return "https://github.com/advisories/" + osv.ID
}

func (osv *OSV) IsAffected(pkg detector.PackageDetails) bool {
	if osv.Affected == nil {
		fmt.Printf("Ignoring %s as it does not have an 'affected' property\n", osv.ID)

		return false
	}

	for _, affected := range osv.Affected {
		if affected.Package.Ecosystem == pkg.Ecosystem &&
			affected.Package.NormalizedName() == pkg.Name {
			if len(affected.Ranges) == 0 {
				_, _ = fmt.Fprintf(
					os.Stderr,
					"%s does not have any ranges - this is probably a mistake!\n",
					osv.ID,
				)

				continue
			}

			if affected.Ranges.AffectsEcosystem(pkg.Version) {
				return true
			}
		}
	}

	return false
}
