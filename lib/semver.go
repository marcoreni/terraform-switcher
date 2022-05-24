package lib

import (
	"fmt"

	semver "github.com/hashicorp/go-version"
)

// GetSemver : returns version that will be installed based on server constraint provided
func GetSemver(tfconstraint *string, mirrorURL *string) (*Release, error) {

	tfReleases, err := GetTFReleases(*mirrorURL, true)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Reading required version from constraint: %s\n", *tfconstraint)
	release, err := SemVerParser(tfconstraint, tfReleases)
	return release, err
}

// ValidateSemVer : Goes through the list of terraform version, return a valid tf version for contraint provided
func SemVerParser(tfconstraint *string, tfReleases []*Release) (*Release, error) {
	constraints, err := semver.NewConstraint(*tfconstraint) //NewConstraint returns a Constraints instance that a Version instance can be checked against
	if err != nil {
		return nil, fmt.Errorf("error parsing constraint: %q", err)
	}
	for _, release := range tfReleases {
		if constraints.Check(release.Version) {
			fmt.Printf("Matched version: %s\n", release.Version)
			return release, nil
		}
	}

	PrintInvalidTFVersion()
	return nil, fmt.Errorf("error parsing constraint: %s", *tfconstraint)
}

// Print invalid TF version
func PrintInvalidTFVersion() {
	fmt.Println("Version does not exist or invalid terraform version format.\n Format should be #.#.# or #.#.#-@# where # are numbers and @ are word characters.\n For example, 0.11.7 and 0.11.9-beta1 are valid versions")
}

// Print invalid TF version
func PrintInvalidMinorTFVersion() {
	fmt.Println("Invalid minor terraform version format. Format should be #.# where # are numbers. For example, 0.11 is valid version")
}