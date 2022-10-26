package environment

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

// Environment variables used to configure the environment.
var (
	EnvSecretsPath = Variable{Key: "SECRETS_PATH", Default: "/run/secrets", Description: "Path where the secrets are stored."}
	EnvDevelopment = Variable{Key: "OPENSLIDES_DEVELOPMENT", Default: "false", Description: "If set, the service uses the default secrets."}
)

// Variable represents a environment variable. It can be used by the packages
// for configuration.
//
// It is only allowed to use an environment variable at startup time.
type Variable struct {
	Key         string
	Default     string
	Description string
}

// Value returns the value for an environment.Variable using a Getenver.
func (v Variable) Value(lookup Getenver) string {
	if lookup == nil {
		return v.Default
	}

	val := lookup.Getenv(v.Key)
	if val == "" {
		return v.Default
	}
	return val
}

// Secret returns the value as secret.
//
// It uses the environment varialbe SECRETS_PATH to find the secrets. The
// defaults are only (and allways) used if OPENSLIDES_DEVELOPMENT is set to
// true. If no Default is set, then "openslides" is used as default.
func (v Variable) Secret(lookup Getenver) string {
	useDev, _ := strconv.ParseBool(EnvDevelopment.Value(lookup))

	if useDev {
		defaultVal := v.Default
		if defaultVal == "" {
			defaultVal = "openslides"
		}

		return defaultVal
	}

	path := path.Join(EnvSecretsPath.Value(lookup), v.Key)
	secret, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Can not read secret in %s: %v", path, err))
	}

	return string(secret)
}

// Getenver is an type, that can return environment variables with a function like
// os.Getenver().
type Getenver interface {
	Getenv(key string) string
}

// Getenvfunc is a function that implements the Getenver interface.
//
// Example: lookup := Getenvfunc(os.Getenv)
type Getenvfunc func(key string) string

// Getenv calls the function.
func (f Getenvfunc) Getenv(key string) string {
	return f(key)
}

// ForTests is a map that simulates environment variables.
type ForTests map[string]string

// Getenv returns the fake enironment variable for a key.
func (e ForTests) Getenv(key string) string {
	v := e[key]

	if key == EnvDevelopment.Key && v == "" {
		return "true"
	}

	return v
}
