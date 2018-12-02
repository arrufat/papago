package papago

import "errors"

// Language defines supported languages by Papago
type Language int32

// Languages supported by Papago
const (
	Korean Language = iota
	English
	Japanese
	Chinese
	TraditionalChinese
	Spanish
	French
	German
	Russian
	Portuguese
	Italian
	Vietnamese
	Thai
	Indonesian
	Hindi
)

// SupportedLanguages returns the languages supported by Papago
func SupportedLanguages() []Language {
	return []Language{
		Korean,
		English,
		Japanese,
		Chinese,
		TraditionalChinese,
		Spanish,
		French,
		German,
		Russian,
		Portuguese,
		Italian,
		Vietnamese,
		Thai,
		Indonesian,
		Hindi,
	}
}

// String prints the canonical Language name
func (lang Language) String() string {
	names := [...]string{
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
	return names[lang]
}

// Code returns the language code
func (lang Language) Code() string {
	codes := [...]string{
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
	return codes[lang]
}

// ParseLanguageCode returns the corresponding Language for a given code
func ParseLanguageCode(code string) (Language, error) {
	var lang Language
	for _, lang = range SupportedLanguages() {
		if lang.Code() == code {
			return lang, nil
		}
	}
	return lang, errors.New("unable to parse the language code")
}

// FemaleVoice returns the name of the female voice for that Language
func (lang Language) FemaleVoice() string {
	names := [...]string{
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
	return names[lang]
}

// MaleVoice returns the name of the female voice for that Language
func (lang Language) MaleVoice() string {
	names := [...]string{
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
	return names[lang]
}

// Speed controls the TTS speed
type Speed int32

// Possible values for TTS speed
const (
	VerySlow Speed = iota
	Slow
	Normal
	Fast
)

func (speed Speed) String() string {
	values := [...]string{
		"5",
		"3",
		"0",
		"-1",
	}
	return values[speed]
}

// Gender used for the TTS
type Gender int32

// Possible values for Gender
const (
	Male = iota
	Female
)

func (gender Gender) String() string {
	names := [...]string{
		"Male",
		"Female",
	}
	return names[gender]
}

// Voice contains the parameters for text to speech generation
type Voice struct {
	Language Language
	Gender   Gender
	Speed    Speed
}
