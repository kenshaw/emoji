package emoji

// emoticonMap is a map of emoji aliases to emoticon counterparts.
var emoticonMap = map[string][]string{
	"angry":                        []string{">:(", ">:-("},
	"anguished":                    []string{"D:"},
	"broken_heart":                 []string{"</3"},
	"confused":                     []string{":/", ":-/", `:\`, `:-\`},
	"disappointed":                 []string{":(", "):", ":-("},
	"heart":                        []string{"<3"},
	"kiss":                         []string{":*", ":-*"},
	"laughing":                     []string{":>", ":->"},
	"monkey_face":                  []string{":o)"},
	"neutral_face":                 []string{":|"},
	"open_mouth":                   []string{":o", ":O", ":-o", ":-O"},
	"slightly_smiling_face":        []string{":)", "(:", ":-)"},
	"smile":                        []string{":D", ":-D"},
	"smiley":                       []string{"=)", "=-)"},
	"stuck_out_tongue":             []string{":p", ":P", ":-p", ":-P", ":b", ":-b"},
	"stuck_out_tongue_winking_eye": []string{";p", ";P", ";-p", ";-P", ";b", ";-b"},
	"sunglasses":                   []string{"8)"},
	"wink":                         []string{";)", ";-)"},
}
