package papago

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// translateURL contains Papago's translation URL.
	translateURL string = "https://papago.naver.com/apis/n2mt/translate"
	// translateParams contains the formating string for a translation request on Papago.
	translateParams string = "dict=%v&dictDisplay=%d&instant=%v&paging=%v&source=%s&target=%s&honorific=%v&text=%s"
	// ttsURL contains Papago's TTS URL.
	ttsURL string = "https://papago.naver.com/apis/tts/makeID"
	// ttsParams contains the formating string for a TTS request on Papago.
	ttsParams string = "alpha=0&pitch=%d&speaker=%s&speed=%s&text=%s"
	// detectURL contains Papago's Language Detection URL.
	detectURL string = "https://papago.naver.com/apis/langs/dect"
	// detectParams contains the formating string for a Language Detection request on Papago.
	detectParams string = "query=%s"
)

// TranslateOptions defines the options for the Translate function
type TranslateOptions struct {
	// Dict controls wether to request the dictionary
	Dict bool
	// DictDisplay sets the maximum amount of entries in the dictionary
	DictDisplay int
	// Instant requests instant translation
	Instant bool
	// Paging request
	Paging bool
	// Source is the language code for the source language
	Source string
	// Target is the language code for the target language
	Target string
	// Honorific requests honorific translation (from en to ko only)
	Honorific bool
	// Text is the string to be translated from the source to the target language
	Text string
}

func (opt TranslateOptions) String() string {
	return fmt.Sprintf(translateParams,
		opt.Dict, opt.DictDisplay, opt.Instant, opt.Paging,
		opt.Source, opt.Target, opt.Honorific, opt.Text)
}

// TranslateResponse contains the structure of a translate response
type TranslateResponse struct {
	SrcLangType    string
	TarLangType    string
	TranslatedText string
	Dict           Dict
	TarDict        Dict
	Delay          int
	DelaySmt       int
}

// Dict structure in a TranslateResponse
type Dict struct {
	Items []DictItem
}

// DictItem is an entry in a Dict structure
type DictItem struct {
	Entry         string
	SubEntry      string
	MatchType     string
	HanjaEntry    string
	PhoneticSigns []PhoneticSign
	Pos           []ItemPos
	Source        string
	URL           string
	MURL          string
}

// PhoneticSign describes the phonetic sign in a DictItem
type PhoneticSign struct {
	Type string
	Sign string
}

// ItemPos describes the item pos in a DictItem
type ItemPos struct {
	Type string
}

// PosMeaning describes the pos meaning in a DictItem
type PosMeaning struct {
	Meaning string
}

// Translate translates the text from source Language to target Language
func Translate(text string, source Language, target Language, opt TranslateOptions) (string, error) {
	text = strings.Replace(text, "\n", "\\n", -1)
	opt.Source = source.Code()
	opt.Target = target.Code()
	opt.Text = text
	data := fmt.Sprintf("%s", opt)
	resp, err := http.Post(translateURL, "text/plain", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var res TranslateResponse
	err = json.Unmarshal(bodyByte, &res)
	if err != nil {
		return "", err
	}
	return res.TranslatedText, nil
}

// TTS generates a URL to the MP3 file containing the sound
func TTS(text string, voice Voice) (string, error) {
	var name string
	if voice.Gender == Male {
		name = voice.Language.MaleVoice()
	} else {
		name = voice.Language.FemaleVoice()
	}
	if name == "" {
		desc := fmt.Sprintf("%s voice not supported for %s", voice.Gender, voice.Language)
		return "", errors.New(desc)
	}
	text = strings.Replace(text, "\n", "\\n", -1)
	params := fmt.Sprintf(ttsParams, voice.Pitch, name, voice.Speed, text)
	data := fmt.Sprintf("%s", params)
	resp, err := http.Post(ttsURL, "text/plain", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(bodyByte, &res)
	if err != nil {
		return "", err
	}
	ans, ok := res["id"].(string)
	if !ok {
		return "", errors.New("error decoding TTS type")
	}
	fileURL := strings.Replace(ttsURL, "makeID", ans, 1)
	return fileURL, nil
}

// Detect tries to guess the input language from the given text
func Detect(text string) (Language, error) {
	var lang Language
	text = strings.Replace(text, "\n", "\\n", -1)
	params := fmt.Sprintf(detectParams, text)
	data := fmt.Sprintf("%s", params)
	resp, err := http.Post(detectURL, "text/plain", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return lang, err
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return lang, err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(bodyByte, &res)
	if err != nil {
		return lang, err
	}
	ans, ok := res["langCode"].(string)
	if !ok {
		return lang, errors.New("error decoding language code")
	}
	return ParseLanguageCode(ans)
}
