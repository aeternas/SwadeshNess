package configuration

type Security struct {
	NeedsHTTPS     bool
	ServerKeyPath  string
	ServerCertPath string
}
