package rbxbin

//go:generate go run dirgen/main.go

// PackageDirectories is a map of where binary deployment packages should go.
type PackageDirectories map[string]string
