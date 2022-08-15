package lockfile

import "github.com/another-rex/osv-detector/internal"

type Ecosystem = internal.Ecosystem
type PackageDetails = internal.PackageDetails
type PackageDetailsParser = func(pathToLockfile string) ([]PackageDetails, error)
