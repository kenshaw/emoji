// Package emoji provides standard tools for working with emoji data.
package emoji

//go:generate go run gen.go

// Gemoji is the set of all emoji data.
type Gemoji []Emoji

// Emoji represenst a single emoji and associated information.
type Emoji struct {
	Emoji          string   `json:"emoji"`
	Description    string   `json:"description"`
	Category       string   `json:"category"`
	Aliases        []string `json:"aliases"`
	Tags           []string `json:"tags"`
	UnicodeVersion string   `json:"unicode_version"`
	IOSVersion     string   `json:"ios_version"`
}
