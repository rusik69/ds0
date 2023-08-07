package env

import "os"

// Parse parses the environment of a web server
func Parse() Env {
	listenPort := os.Getenv("WEB_LISTEN_PORT")
	if listenPort == "" {
		listenPort = "80"
	}
	listenHost := os.Getenv("WEB_LISTEN_HOST")
	if listenHost == "" {
		listenHost = "0.0.0.0"
	}
	nsHost := os.Getenv("WEB_NS_HOST")
	if nsHost == "" {
		nsHost = "localhost"
	}
	nsPort := os.Getenv("WEB_NS_PORT")
	if nsPort == "" {
		nsPort = "6969"
	}
	envInstance := Env{
		ListenPort: listenPort,
		ListenHost: listenHost,
		NSHost:     nsHost,
		NSPort:     nsPort,
	}
	return envInstance
}
