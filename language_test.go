package papago

import (
	"testing"
)

var langs = SupportedLanguages()

var names = [...]string{
	"Korean",
	"English",
	"Japanese",
	"Simplified Chinese",
	"Traditional Chinese",
	"Spanish",
	"French",
	"German",
	"Russian",
	"Portuguese",
	"Italian",
	"Vietnamese",
	"Thai",
	"Indonesian",
	"Hindi",
}

var codes = [...]string{
	"ko",    // Korean
	"en",    // English
	"ja",    // Japanese
	"zh-CN", // Simplified Chinese
	"zh-TW", // Traditional Chinese
	"es",    // Spanish
	"fr",    // French
	"ge",    // German
	"ru",    // Russian
	"pt",    // Portuguese
	"it",    // Italian
	"vi",    // Vietnamese
	"th",    // Thai
	"id",    // Indonesian
	"hi",    // Hindi
}

var femaleVoices = [...]string{
	"kyuri",  // Korean
	"clara",  // English
	"yuri",   // Japanese
	"meimei", // Simplified Chinese
	"",       // Traditional Chinese
	"carmen", // Spanish
	"roxane", // French
	"",       // German
	"",       // Russian
	"",       // Portuguese
	"",       // Italian
	"",       // Vietnamese
	"",       // Italian
	"",       // Thai
	"",       // Indonesian
	"",       // Hindi
}

var maleVoices = [...]string{
	"jinho",      // Korean
	"matt",       // English
	"shinji",     // Japanese
	"liangliang", // Simplified Chinese
	"",           // Traditional Chinese
	"jose",       // Spanish
	"louis",      // French
	"",           // German
	"",           // Russian
	"",           // Portuguese
	"",           // Italian
	"",           // Vietnamese
	"",           // Italian
	"",           // Thai
	"",           // Indonesian
	"",           // Hindi
}

func TestLanguageLengths(t *testing.T) {
	if len(langs) != len(names) {
		t.Errorf("Different lengths for languages (%d) and names (%d)", len(langs), len(names))
	}
}

func TestLanguageNames(t *testing.T) {
	for i := 0; i < len(langs); i++ {
		if langs[i].String() != names[i] {
			t.Errorf("Error in Canonical name for Language %s (expected %s)", langs[i], names[i])
		}
	}
}

func TestLanguageCodes(t *testing.T) {
	for i := 0; i < len(langs); i++ {
		if langs[i].Code() != codes[i] {
			t.Errorf("Error in Code for Language %s (expected %s)", langs[i], names[i])
		}
	}
}

func TestFemaleVoices(t *testing.T) {
	for i := 0; i < len(langs); i++ {
		if langs[i].FemaleVoice() != femaleVoices[i] {
			t.Errorf("Error in Female Voice name for Language %s (expected %s)", langs[i], femaleVoices[i])
		}
	}
}

func TestMaleVoices(t *testing.T) {
	for i := 0; i < len(langs); i++ {
		if langs[i].MaleVoice() != maleVoices[i] {
			t.Errorf("Error in Male Voice name for Language %s (expected %s)", langs[i], maleVoices[i])
		}
	}
}
