package lockfile

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"io/ioutil"
	"regexp"
	"strings"
)

const GoEcosystem Ecosystem = "Go"

func parseGoModVersion(version string) (string, string, string) {
	re := regexp.MustCompile(`^v([\d.]+(?:-[\w.]+)?)[-.](\d{14})-(\w{12})$`)

	matched := re.FindStringSubmatch(version)

	if matched == nil {
		return version, "", ""
	}

	return matched[1], matched[2], matched[3]
}

func ParseGoLock(pathToLockfile string) ([]PackageDetails, error) {
	lockfileContents, err := ioutil.ReadFile(pathToLockfile)

	if err != nil {
		return []PackageDetails{}, fmt.Errorf("could not read %s: %w", pathToLockfile, err)
	}

	parsedLockfile, err := modfile.Parse(pathToLockfile, lockfileContents, nil)

	if err != nil {
		return []PackageDetails{}, fmt.Errorf("could not parse %s: %w", pathToLockfile, err)
	}

	packages := make([]PackageDetails, 0, len(parsedLockfile.Require))

	for _, require := range parsedLockfile.Require {
		version, _, commit := parseGoModVersion(require.Mod.Version)

		packages = append(packages, PackageDetails{
			Name:      require.Mod.Path,
			Version:   strings.TrimPrefix(version, "v"),
			Ecosystem: GoEcosystem,
			Commit:    commit,
		})
	}

	return packages, nil
}
