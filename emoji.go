// Package emoji provides standard tools for working with emoji data.
package emoji

import "strings"

//go:generate go run gen.go

// Gemoji is the set of emoji data.
type Gemoji []Emoji

// Emoji represents a single emoji and its associated data.
type Emoji struct {
	Emoji          string   `json:"emoji"`
	Description    string   `json:"description"`
	Category       string   `json:"category"`
	Aliases        []string `json:"aliases"`
	Tags           []string `json:"tags"`
	UnicodeVersion string   `json:"unicode_version"`
	IOSVersion     string   `json:"ios_version"`
}

// codeMap provides a map of the unicode code to emoji values.
var codeMap map[string]*Emoji

// aliasMap provides a map of the alias to emoji values.
var aliasMap map[string]*Emoji

// codeReplacer is the string replacer for emoji values.
var codeReplacer *strings.Replacer

// aliasReplacer is the string replacer for aliases.
var aliasReplacer *strings.Replacer

func init() {
	// initialize
	codeMap = make(map[string]*Emoji, len(GemojiData))
	aliasMap = make(map[string]*Emoji, len(GemojiData))

	codePairs := make([]string, 0)
	aliasPairs := make([]string, 0)
	for _, e := range GemojiData {
		if e.Emoji == "" || len(e.Aliases) == 0 {
			continue
		}

		// setup codes
		codeMap[e.Emoji] = &e
		codePairs = append(codePairs, e.Emoji, ":"+e.Aliases[0]+":")

		// setup aliases
		for _, a := range e.Aliases {
			if a == "" {
				continue
			}

			aliasMap[a] = &e
			aliasPairs = append(aliasPairs, ":"+a+":", e.Emoji)
		}
	}

	// create replacers
	codeReplacer = strings.NewReplacer(codePairs...)
	aliasReplacer = strings.NewReplacer(aliasPairs...)
}

// FromCode retrieves the emoji entry based on the unicode code.
func FromCode(code string) *Emoji {
	return codeMap[code]
}

// FromAlias retrieves the emoji entry based on the provided alias in the form
// "alias" or ":alias:".
func FromAlias(alias string) *Emoji {
	if strings.HasPrefix(alias, ":") && strings.HasSuffix(alias, ":") {
		alias = alias[1 : len(alias)-1]
	}

	return aliasMap[alias]
}

// ReplaceCodes is the inverse of ReplaceAliases, replacing all emoji codes
// with its corresponding alias.
func ReplaceCodes(s string) string {
	return codeReplacer.Replace(s)
}

// ReplaceAliases replaces all aliases of the form ":alias:" with its
// corresponding unicode value.
func ReplaceAliases(s string) string {
	return aliasReplacer.Replace(s)
}
