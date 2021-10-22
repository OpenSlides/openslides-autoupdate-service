package modelsversion

const (
	// Repo is the url to the repo where the models.yml and simular files are
	// located.
	Repo = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/"

	// Version is the git-tag to use.
	//
	// Can be the name of a branch like "master" or the name of a tag like
	// "v1.0.1" or a commit id.
	Version = "master"
)

const (
	modelsPath      = "/docs/models.yml"
	exampleDataPath = "/docs/example-data.json"
	permissionPath  = "/docs/permission.yml"
)

// ModelsYMLURL return the url to the models.yml.
func ModelsYMLURL() string {
	return Repo + Version + modelsPath
}

// ExampleDataURL returns the url to the example-data file.
func ExampleDataURL() string {
	return Repo + Version + exampleDataPath
}

// PermissionURL returns the url to the example-data file.
func PermissionURL() string {
	return Repo + Version + permissionPath
}
