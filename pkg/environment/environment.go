// Package environment implements helpers to handle environment varialbes and
// other function for startup.
package environment

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"

	"golang.org/x/sys/unix"
)

// Environment variables used to configure the environment.
var (
	envSecretsPath = NewVariable("SECRETS_PATH", "/run/secrets", "Path where the secrets are stored.")
	envDevelopment = NewVariable("OPENSLIDES_DEVELOPMENT", "false", "If set, the service uses the default secrets.")
)

// Variable represents a environment variable. It can be used by the packages
// for configuration.
//
// It is only allowed to use an environment variable at startup time.
type Variable struct {
	Key         string
	Default     string
	Description string
	isSecret    bool
}

// NewVariable initializes a environment.Variable
func NewVariable(key, defaultValue, description string) Variable {
	return Variable{
		Key:         key,
		Default:     defaultValue,
		Description: description,
		isSecret:    false,
	}
}

// NewSecret initializes a secret.
func NewSecret(key, description string) Variable {
	return Variable{
		Key:         key,
		Default:     "openslides",
		Description: description,
		isSecret:    true,
	}
}

// NewSecretWithDefault initializes a secret with a secial default value.
//
// Try not to use this. The default value for all secrets should be 'openslides'.
func NewSecretWithDefault(key, defaultValue, description string) Variable {
	return Variable{
		Key:         key,
		Default:     defaultValue,
		Description: description,
		isSecret:    true,
	}
}

// Value returns the value for an environment.Variable using a Getenver.
func (v Variable) Value(lookup Environmenter) string {
	lookup.UseVariable(v)

	if v.isSecret {
		return v.secret(lookup)
	}

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
func (v Variable) secret(lookup Environmenter) string {
	useDev, _ := strconv.ParseBool(envDevelopment.Value(lookup))

	if useDev {
		defaultVal := v.Default
		if defaultVal == "" {
			defaultVal = "openslides"
		}

		return defaultVal
	}

	path := path.Join(envSecretsPath.Value(lookup), v.Key)
	secret, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Can not read secret in %s: %v", path, err))
	}

	return string(secret)
}

// Environmenter is an type, that can return the value for environment
// variables.
//
// It also saves the used environment variables.
type Environmenter interface {
	Getenv(key string) string
	UseVariable(v Variable)
}

// ForProduction is an environment used for production.
//
// It fetches the environment variables from os.Getenv()
type ForProduction struct {
	usedVariables []Variable
}

// Getenv calls os.Getenv.
func (e *ForProduction) Getenv(key string) string {
	return os.Getenv(key)
}

// UseVariable saves the used Variables
func (e *ForProduction) UseVariable(v Variable) {
	e.usedVariables = append(e.usedVariables, v)
}

// BuildDoc create the environment documentation with all used variables.
func (e *ForProduction) BuildDoc() (string, error) {
	var variables struct {
		Env    []Variable
		Secret []Variable
	}

	for _, v := range e.usedVariables {
		if v.isSecret {
			variables.Secret = append(variables.Secret, v)
			continue
		}

		variables.Env = append(variables.Env, v)
	}

	tmpl, err := template.New("Doc").Parse(tmplDoc)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, variables); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	return strings.ReplaceAll(buf.String(), "$", "`"), nil
}

// ForTests is a map that simulates environment variables.
type ForTests map[string]string

// Getenv returns the fake enironment variable for a key.
func (e ForTests) Getenv(key string) string {
	v := e[key]

	if key == envDevelopment.Key && v == "" {
		return "true"
	}

	return v
}

// UseVariable does nothing for tests.
func (e ForTests) UseVariable(v Variable) {}

// ParseDuration is like time.ParseDuration but uses second as default unit.
func ParseDuration(s string) (time.Duration, error) {
	sec, err := strconv.Atoi(s)
	if err == nil {
		return time.Duration(sec) * time.Second, nil
	}

	return time.ParseDuration(s)
}

// InterruptContext works like signal.NotifyContext. It returns a context that
// is canceled, when a signal is received.
//
// It listens on os.Interrupt and unix.SIGTERM. If the signal is received two
// times, os.Exit(2) is called.
func InterruptContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, unix.SIGTERM)
		<-sig
		cancel()
		<-sig
		os.Exit(2)
	}()
	return ctx, cancel
}

const tmplDoc = `<!--- Code generated with go generate ./... DO NOT EDIT. --->
# Configuration

## Environment Variables

The Service uses the following environment variables:
{{range .Env}}
* ${{.Key}}$: {{.Description}} The default is ${{.Default}}$.
{{- end}}

{{if .Secret}}
## Secrets

Secrets are filenames in the directory $SECRETS_PATH$ (default: $/run/secrets/$). 
The service only starts if it can find each secret file and read its content. 
The default values are only used, if the environment variable $OPENSLIDES_DEVELOPMENT$ is set.
{{range .Secret}}
* ${{.Key}}$: {{.Description}} The default is ${{.Default}}$.
{{- end}}
{{- end}}`
