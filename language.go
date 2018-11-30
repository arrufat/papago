package papago

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

// FemaleVoice returns the name of the female voice for that Language
func (lang Language) FemaleVoice() string {
	names := []string{
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
	names := []string{
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
