package emoji_test

import (
	"fmt"
	"strings"

	"github.com/kenshaw/emoji"
)

func Example() {
	e1 := emoji.FromEmoticon(":-)")
	fmt.Println(":-)", "--", e1.Emoji)

	e2 := emoji.FromAlias("slightly_smiling_face")
	fmt.Println(":-)", "--", e2.Emoji)

	s1 := emoji.ReplaceEmoticonsWithAliases(":-) :D >:(")
	fmt.Println(":-) :D >:(", "--", s1)

	s2 := emoji.ReplaceEmoticonsWithCodes(":-) :D >:(")
	fmt.Println(":-) :D >:(", "--", s2)

	// Output:
	// :-) -- 🙂
	// :-) -- 🙂
	// :-) :D >:( -- :slightly_smiling_face: :smile: :angry:
	// :-) :D >:( -- 🙂 😄 😠
}

func ExampleSkinTone() {
	e := emoji.FromAlias("thumbsup")
	s := []string{e.Emoji}
	for skinTone := emoji.Light; skinTone <= emoji.Dark; skinTone++ {
		s = append(s, e.Tone(skinTone))
	}
	fmt.Println(strings.Join(s, " "))
	// Output:
	// 👍 👍🏻 👍🏼 👍🏽 👍🏾 👍🏿
}
