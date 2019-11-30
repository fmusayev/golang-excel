package types

import (
	"strings"
)

// Extending string type for additional methods
type String string

// Extended string method to uppercase text & transliterate azerbaijan symbols
func (s String) Transliterate() String {
	str := strings.ToUpper(string(s))
	replaceables := map[string]string{"Ü": "U", "Ö": "O", "Ğ": "G", "I": "I", "İ": "I", "Ə": "A", "Ç": "CH", "Ş": "SH"}
	for k, v := range replaceables {
		str = strings.ReplaceAll(str, k, v)
	}

	return String(str)
}
