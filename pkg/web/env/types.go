package env

// Env is the environment of a web server
type Env struct {
	ListenPort string
	ListenHost string
	NSHost     string
	NSPort     string
}

// EnvInstance is the instance of Env
var EnvInstance Env
