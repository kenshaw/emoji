package emoji

import (
	"reflect"
	"testing"
)

func TestDumpInfo(t *testing.T) {
	i := 0
	for _, v := range emoticonMap {
		i += len(v)
	}

	t.Logf("codes: %d", len(codeMap))
	t.Logf("aliases: %d", len(aliasMap))
	t.Logf("emoticons: %d", i)
}

func TestLookup(t *testing.T) {
	a := FromCode("\U0001f37a")
	b := FromCode("🍺")
	c := FromAlias(":beer:")
	d := FromAlias("beer")

	if !reflect.DeepEqual(a, b) {
		t.Errorf("a and b should equal")
	}
	if !reflect.DeepEqual(b, c) {
		t.Errorf("b and c should equal")
	}
	if !reflect.DeepEqual(c, d) {
		t.Errorf("c and d should equal")
	}
	if !reflect.DeepEqual(a, d) {
		t.Errorf("a and d should equal")
	}

	m := FromCode("\U0001f44d")
	n := FromAlias(":thumbsup:")
	o := FromAlias("+1")

	if !reflect.DeepEqual(m, n) {
		t.Errorf("m and n should equal")
	}
	if !reflect.DeepEqual(n, o) {
		t.Errorf("n and o should equal")
	}
	if !reflect.DeepEqual(m, o) {
		t.Errorf("m and o should equal")
	}
}

func TestReplacers(t *testing.T) {
	tests := []struct {
		f      func(string) string
		v, exp string
	}{
		{ReplaceCodes, ":thumbsup: +1 for \U0001f37a! 🍺 \U0001f44d", ":thumbsup: +1 for :beer:! :beer: :+1:"},
		{ReplaceAliases, ":thumbsup: +1 :+1: :beer:", "\U0001f44d +1 \U0001f44d \U0001f37a"},

		{ReplaceEmoticonsWithCodes, "foobar", "foobar"},
		{ReplaceEmoticonsWithCodes, " foobar", " foobar"},
		{ReplaceEmoticonsWithCodes, " foobar ", " foobar "},
		{ReplaceEmoticonsWithCodes, ":o)", "\U0001f435"},
		{ReplaceEmoticonsWithCodes, " :o)", " \U0001f435"},
		{ReplaceEmoticonsWithCodes, ":o) ", "\U0001f435 "},
		{ReplaceEmoticonsWithCodes, " :o) ", " \U0001f435 "},
		{ReplaceEmoticonsWithCodes, "a:o)", "a:o)"},
		{ReplaceEmoticonsWithCodes, "a:o) ", "a:o) "},
		{ReplaceEmoticonsWithCodes, "a:o)b", "a:o)b"},
		{ReplaceEmoticonsWithCodes, ":o) :b", "\U0001f435 \U0001f61b"},
		{ReplaceEmoticonsWithCodes, " :o) :b", " \U0001f435 \U0001f61b"},
		{ReplaceEmoticonsWithCodes, ":o) :b ", "\U0001f435 \U0001f61b "},
		{ReplaceEmoticonsWithCodes, " :o) :b ", " \U0001f435 \U0001f61b "},
		{ReplaceEmoticonsWithCodes, ":o) :b :b", "\U0001f435 \U0001f61b \U0001f61b"},
		{ReplaceEmoticonsWithCodes, " :o) :b :b", " \U0001f435 \U0001f61b \U0001f61b"},
		{ReplaceEmoticonsWithCodes, " :o) :b :b ", " \U0001f435 \U0001f61b \U0001f61b "},
		{ReplaceEmoticonsWithCodes, ":o) :b :stuck_out_tongue:", "\U0001f435 \U0001f61b :stuck_out_tongue:"},
		{ReplaceEmoticonsWithCodes, " :o) :b :stuck_out_tongue:", " \U0001f435 \U0001f61b :stuck_out_tongue:"},
		{ReplaceEmoticonsWithCodes, " :o) :b :stuck_out_tongue: ", " \U0001f435 \U0001f61b :stuck_out_tongue: "},
		{ReplaceEmoticonsWithCodes, ":b :o) :beer: \U0001f37a", "\U0001f61b \U0001f435 :beer: \U0001f37a"},

		{ReplaceEmoticonsWithAliases, "foobar", "foobar"},
		{ReplaceEmoticonsWithAliases, " foobar", " foobar"},
		{ReplaceEmoticonsWithAliases, " foobar ", " foobar "},
		{ReplaceEmoticonsWithAliases, ":o)", ":monkey_face:"},
		{ReplaceEmoticonsWithAliases, " :o)", " :monkey_face:"},
		{ReplaceEmoticonsWithAliases, ":o) ", ":monkey_face: "},
		{ReplaceEmoticonsWithAliases, " :o) ", " :monkey_face: "},
		{ReplaceEmoticonsWithAliases, "a:o)", "a:o)"},
		{ReplaceEmoticonsWithAliases, "a:o) ", "a:o) "},
		{ReplaceEmoticonsWithAliases, "a:o)b", "a:o)b"},
		{ReplaceEmoticonsWithAliases, ":o) :b", ":monkey_face: :stuck_out_tongue:"},
		{ReplaceEmoticonsWithAliases, " :o) :b", " :monkey_face: :stuck_out_tongue:"},
		{ReplaceEmoticonsWithAliases, ":o) :b ", ":monkey_face: :stuck_out_tongue: "},
		{ReplaceEmoticonsWithAliases, " :o) :b ", " :monkey_face: :stuck_out_tongue: "},
		{ReplaceEmoticonsWithAliases, ":o) :b :b", ":monkey_face: :stuck_out_tongue: :stuck_out_tongue:"},
		{ReplaceEmoticonsWithAliases, " :o) :b :b", " :monkey_face: :stuck_out_tongue: :stuck_out_tongue:"},
		{ReplaceEmoticonsWithAliases, " :o) :b :b ", " :monkey_face: :stuck_out_tongue: :stuck_out_tongue: "},
		{ReplaceEmoticonsWithAliases, ":o) :b :stuck_out_tongue:", ":monkey_face: :stuck_out_tongue: :stuck_out_tongue:"},
		{ReplaceEmoticonsWithAliases, " :o) :b :stuck_out_tongue:", " :monkey_face: :stuck_out_tongue: :stuck_out_tongue:"},
		{ReplaceEmoticonsWithAliases, " :o) :b :stuck_out_tongue: ", " :monkey_face: :stuck_out_tongue: :stuck_out_tongue: "},
		{ReplaceEmoticonsWithAliases, ":b :o) :beer: \U0001f37a", ":stuck_out_tongue: :monkey_face: :beer: \U0001f37a"},

		{ReplaceAliasesWithEmoticons, ":monkey_face: monkey_face :beer:", ":o) monkey_face :beer:"},

		// FIXME: these will do replacements for :b, but shouldn't ...
		//{ReplaceEmoticonsWithCodes, ":o):b", ":o):b"},
		//{ReplaceEmoticonsWithAliases, ":o):b", ":o):b"},
	}

	for i, x := range tests {
		s := x.f(x.v)
		if s != x.exp {
			t.Errorf("test %d `%s` expected `%s`, got: `%s`", i, x.v, x.exp, s)
		}
	}
}
