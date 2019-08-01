package packs

// Package is comprised of fields from all package types.
type Package struct {
	Name     string `json:"name,omitempty"`     // Package Name
	Version  string `json:"version,omitempty"`  // Package Version
	Source   string `json:"source,omitempty"`   // Source (repo) of a Package
	Path     string `json:"path,omitempty"`     // Path to local installation
	From     string `json:"from,omitempty"`     // NPM
	Resolved string `json:"resolved,omitempty"` // NPM
}
