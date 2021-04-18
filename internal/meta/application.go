package meta

var (
	applicationName string
	buildTime       string
	sha1            string
	version         string
)

// ApplicationName returns the name of the application
func ApplicationName() string {
	return applicationName
}

// BuildTime returns the date and time the application binaries were originally created
func BuildTime() string {
	return buildTime
}

// Revision returns the GIT SHA of the commit from which the application binaries were originally created
func Revision() string {
	return sha1
}

// Version returns the application version
func Version() string {
	return version
}
