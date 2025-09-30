package service

var (
	appMode = "prod"
	version = "alpha-0.0.8"
)

type VersionService struct {
}

func NewVersionService() *VersionService {
	return &VersionService{}
}

// GetVersion returns the current version of the service.
// This function provides a way to retrieve the version information
// that is stored in the version variable.
//
// Returns:
//   - string: The current version as a string value.
func (s VersionService) GetVersion() string {
	return version
}

func IsProd() bool {
	return appMode == "prod"
}
