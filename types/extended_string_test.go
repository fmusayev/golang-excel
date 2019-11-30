package types

import (
	"testing"
)

func TestTransliterate(t *testing.T) {

	testData := []struct {
		text   String
		result string
	}{
		{"TestTEXT", "TESTTEXT"},
		{"aüüöüçsaaşə", "AUUOUCHSAASHA"},
		{"ÜüƏəÇçŞşÖöĞğİiIı", "UUAACHCHSHSHOOGGIIII"},
	}

	for _, v := range testData {
		r := v.text.Transliterate()
		if string(r) != v.result {
			t.Errorf("TransliterateText was incorrect, got: %s, want: %s.", r, v.result)
		}
	}
}
