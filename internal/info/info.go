package info

var (
	sha1 string
	buildTime string
	version string
)

func BuildTime() string {
	return buildTime
}

func Revision() string {
	return sha1
}

func Version() string {
	return version
}