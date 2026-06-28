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

	emoji.AddAlias(":rolling_on_the_floor_laughing:", ":rofl:")
	e3 := emoji.FromAlias("rolling_on_the_floor_laughing")
	fmt.Println(e3.Emoji)

	s1 := emoji.ReplaceEmoticonsWithAliases(":-) :D >:(")
	fmt.Println(":-) :D >:(", "--", s1)

	s2 := emoji.ReplaceEmoticonsWithCodes(":-) :D >:(")
	fmt.Println(":-) :D >:(", "--", s2)

	s3 := emoji.ReplaceAliases("new alias :rolling_on_the_floor_laughing: works")
	fmt.Println(s3)

	// Output:
	// :-) -- 🙂
	// :-) -- 🙂
	// 🤣
	// :-) :D >:( -- :slightly_smiling_face: :smile: :angry:
	// :-) :D >:( -- 🙂 😄 😠
	// new alias 🤣 works
}

func ExampleSkinTone() {
	for _, alias := range []string{"thumbsup", "man_technologist", "couplekiss_woman_woman", "female_detective"} {
		e := emoji.FromAlias(alias)
		s := []string{e.Emoji}
		for skinTone := emoji.Light; skinTone <= emoji.Dark; skinTone++ {
			s = append(s, e.Tone(skinTone))
		}
		fmt.Println(strings.Join(s, " "))
	}
	for _, alias := range []string{":+1_light_skin_tone:", ":+1_medium_light_skin_tone:", "+1_dark_skin_tone"} {
		e := emoji.FromAlias(alias)
		fmt.Println(e)
	}
	// Output:
	// 👍 👍🏻 👍🏼 👍🏽 👍🏾 👍🏿
	// 👨‍💻 👨🏻‍💻 👨🏼‍💻 👨🏽‍💻 👨🏾‍💻 👨🏿‍💻
	// 👩‍❤️‍💋‍👩 👩🏻‍❤️‍💋‍👩🏻 👩🏼‍❤️‍💋‍👩🏼 👩🏽‍❤️‍💋‍👩🏽 👩🏾‍❤️‍💋‍👩🏾 👩🏿‍❤️‍💋‍👩🏿
	// 🕵️‍♀️ 🕵🏻️‍♀️ 🕵🏼️‍♀️ 🕵🏽️‍♀️ 🕵🏾️‍♀️ 🕵🏿️‍♀️
	// 👍🏻
	// 👍🏼
	// 👍🏿
}
