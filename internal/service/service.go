package service

import "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"

func ConvMorseString(morseOrString string) string {
	var formatMorse bool
	var resConv string
	for _, v := range morseOrString {
		if v == 45 || v == 46 || v == 32 {
			formatMorse = true
		} else {
			formatMorse = false
			break
		}
	}
	if formatMorse {

		// ToText converts a morse string to his textual representation, it is an alias to DefaultConverter.ToText.
		resConv = morse.ToText(morseOrString)
	}
	if !formatMorse {
		// ToMorse converts a text to his morse rrpresentation, it is an alias to DefaultConverter.ToMorse.
		resConv = morse.ToMorse(morseOrString)
	}
	return resConv
}
