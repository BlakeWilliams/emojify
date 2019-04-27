package emojify

import "testing"

func TestNesting(t *testing.T) {
	originalString := "heart"

	emojiString := Emojify(originalString)

	if emojiString != ":heart:" {
		t.Errorf("Wanted :heart:; got %s", emojiString)
	}
}

func TestNormal(t *testing.T) {
	originalString := "heart hater"

	emojiString := Emojify(originalString)

	if emojiString != ":heart: h:a:ter" {
		t.Errorf("\n- :heart: h:a:ter\n+ %s", emojiString)
	}
}
