// Package emoji provides standard tools for working with emoji unicode codes and aliases.
package emoji

import (
	"bytes"
	"regexp"
	"strings"
)

//go:generate go run gen.go

// Gemoji is a set of emoji data.
type Gemoji []Emoji

// Emoji represents a single emoji and associated data.
type Emoji struct {
	Emoji          string   `json:"emoji"`
	Description    string   `json:"description"`
	Category       string   `json:"category"`
	Aliases        []string `json:"aliases"`
	Tags           []string `json:"tags"`
	UnicodeVersion string   `json:"unicode_version"`
	IOSVersion     string   `json:"ios_version"`
}

var (
	// codeMap provides a map of the emoji unicode code to its emoji data.
	codeMap map[string]int

	// aliasMap provides a map of the alias to its emoji data.
	aliasMap map[string]int

	// codeReplacer is the string replacer for emoji codes.
	codeReplacer *strings.Replacer

	// aliasReplacer is the string replacer for emoji aliases.
	aliasReplacer *strings.Replacer

	// aliasEmoticonReplacer is the string replacer for emoji aliases with
	// emoticons.
	aliasEmoticonReplacer *strings.Replacer

	// emoticonRE is the regexp to match emoticons on word boundaries.
	emoticonRE *regexp.Regexp

	// emoticonCodeMap is the map of emoticons to their emoji value.
	emoticonCodeMap map[string]string

	// emoticonCodeMap is the map of emoticons to their emoji alias.
	emoticonAliasMap map[string]string
)

func init() {
	// initialize
	codeMap = make(map[string]int, len(GemojiData))
	aliasMap = make(map[string]int, len(GemojiData))
	emoticonCodeMap = make(map[string]string, 0)
	emoticonAliasMap = make(map[string]string, 0)

	// process emoji codes and aliases
	codePairs := make([]string, 0)
	aliasPairs := make([]string, 0)
	for i, e := range GemojiData {
		if e.Emoji == "" || len(e.Aliases) == 0 {
			continue
		}

		// setup codes
		codeMap[e.Emoji] = i
		codePairs = append(codePairs, e.Emoji, ":"+e.Aliases[0]+":")

		// setup aliases
		for _, a := range e.Aliases {
			if a == "" {
				continue
			}

			aliasMap[a] = i
			aliasPairs = append(aliasPairs, ":"+a+":", e.Emoji)
		}
	}

	// process emoticons
	reVals := make([]string, 0)
	aliasEmoticonPairs := make([]string, 0)
	for a, vals := range emoticonMap {
		alias := ":" + a + ":"
		aliasEmoticonPairs = append(aliasEmoticonPairs, alias, vals[0])
		for _, u := range vals {
			reVals = append(reVals, regexp.QuoteMeta(u))
			emoticonCodeMap[u] = GemojiData[aliasMap[a]].Emoji
			emoticonAliasMap[u] = alias
		}
	}

	// create emoticon regexp
	emoticonRE = regexp.MustCompile(`(?m:^|\A|\s|\B)(` + strings.Join(reVals, "|") + `)(?:$|\z|\s)`)

	// create replacers
	codeReplacer = strings.NewReplacer(codePairs...)
	aliasReplacer = strings.NewReplacer(aliasPairs...)
	aliasEmoticonReplacer = strings.NewReplacer(aliasEmoticonPairs...)
}

// FromCode retrieves the emoji data based on the unicode code.
func FromCode(code string) *Emoji {
	i, ok := codeMap[code]
	if !ok {
		return nil
	}

	return &GemojiData[i]
}

// FromAlias retrieves the emoji data based on the provided alias in the form
// "alias" or ":alias:".
func FromAlias(alias string) *Emoji {
	if strings.HasPrefix(alias, ":") && strings.HasSuffix(alias, ":") {
		alias = alias[1 : len(alias)-1]
	}

	i, ok := aliasMap[alias]
	if !ok {
		return nil
	}

	return &GemojiData[i]
}

// ReplaceCodes replaces all emoji codes with its first corresponding emoji
// alias.
func ReplaceCodes(s string) string {
	return codeReplacer.Replace(s)
}

// ReplaceAliases replaces all aliases of the form ":alias:" with its
// corresponding unicode value.
func ReplaceAliases(s string) string {
	return aliasReplacer.Replace(s)
}

// emoticonReplacer replaces all emoticons in s with the corresponding repl
// value.
func emoticonReplacer(s string, repl map[string]string) string {
	matches := emoticonRE.FindAllStringSubmatchIndex(s, -1)

	// bail if no matches
	if len(matches) == 0 {
		return s
	}

	// build replacement string
	var buf bytes.Buffer
	last := 0
	for _, m := range matches {
		buf.WriteString(s[last:m[2]])
		e, ok := repl[s[m[2]:m[3]]]
		if !ok {
			panic("could not find emoticon!!")
		}
		buf.WriteString(e)
		last = m[3]
	}
	buf.WriteString(s[last:])

	return string(buf.Bytes())
}

// ReplaceEmoticonsWithCodes replaces all emoticons (ie, :D, :p, etc) with its
// corresponding emoji code.
func ReplaceEmoticonsWithCodes(s string) string {
	return emoticonReplacer(s, emoticonCodeMap)
}

// ReplaceEmoticonsWithAliases replaces all emoticons (ie, :D, :p, etc) with
// its corresponding emoji alias (in the form of :alias:).
func ReplaceEmoticonsWithAliases(s string) string {
	return emoticonReplacer(s, emoticonAliasMap)
}

// ReplaceAliasesWithEmoticons replaces all emoji aliases (in the form of
// :alias:) with its corresponding emoticon (ie, :D, :p, etc) if the alias has
// a corresponding emoticon.
func ReplaceAliasesWithEmoticons(s string) string {
	return aliasEmoticonReplacer.Replace(s)
}
