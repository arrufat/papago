package papago

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// TranslateURL contains Papago's translation URL.
	TranslateURL string = "https://papago.naver.com/apis/n2mt/translate"
	// TranslateHeader contains Papago's translation header for a request.
	TranslateHeader string = "\xaeU\xb10\xa3\x1c/b\x160\xf5z\"7b0e4eca-c538-417f-8bf5-43a9e6ef160b\","
	// TranslateParams contains the formating string for a translation request on Papago.
	TranslateParams string = "\"dict\":true,\"dictDisplay\":30,\"source\":\"%s\",\"target\":\"%s\",\"text\":\"%s\"}"
	// TtsURL contains Papago's TTS URL.
	TtsURL string = "https://papago.naver.com/apis/tts/makeID"
	// TtsHeader contains Papago's TTS header for a request.
	TtsHeader string = "\xaeU\xae\xa1C\x9b,Uzd\xf8\xef"
	// TtsParams contains the formating string for a TTS request on Papago.
	TtsParams string = "pitch\":0,\"speaker\":\"%s\",\"speed\":0,\"text\":\"%s\"}"
	// DetectURL contains Papago's Language Detection URL.
	DetectURL string = "https://papago.naver.com/apis/langs/dect"
	// DetectHeader contains Papago's Language Detection header for a request.
	DetectHeader string = "\xaeU\xa4\xa8\x92%\xacUzV\xfd"
	// DetectParams contains the formating string for a Language Detection request on Papago.
	DetectParams string = "-%s\"}"
)

// Translate translates the text from source Language to target Language
func Translate(text string, source Language, target Language) (string, error) {
	params := fmt.Sprintf(TranslateParams, source.Code(), target.Code(), text)
	data := fmt.Sprintf("%s%s", TranslateHeader, params)
	encData := base64.StdEncoding.EncodeToString([]byte(data))
	body := fmt.Sprintf("data=%s", encData)
	resp, err := http.Post(TranslateURL, "text/plain", bytes.NewBuffer([]byte(body)))
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
	ans, ok := res["translatedText"].(string)
	if !ok {
		return "", errors.New("error decoding translate type")
	}
	return ans, nil
}

// TTS generates a URL to the MP3 file containing the sound
func TTS(text string, lang Language, gender Gender) (string, error) {
	var voice string
	if gender == Male {
		voice = lang.MaleVoice()
	} else {
		voice = lang.FemaleVoice()
	}
	if voice == "" {
		desc := fmt.Sprintf("%s voice not supported for %s", gender, lang)
		return "", errors.New(desc)
	}
	params := fmt.Sprintf(TtsParams, voice, text)
	data := fmt.Sprintf("%s%s", TtsHeader, params)
	encData := base64.StdEncoding.EncodeToString([]byte(data))
	body := fmt.Sprintf("data=%s", encData)
	resp, err := http.Post(TtsURL, "text/plain", bytes.NewBuffer([]byte(body)))
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
	fileURL := strings.Replace(TtsURL, "makeID", ans, 1)
	return fileURL, nil
}

// Detect tries to guess the input language from the given text
func Detect(text string) (Language, error) {
	var lang Language
	params := fmt.Sprintf(DetectParams, text)
	data := fmt.Sprintf("%s%s", DetectHeader, params)
	encData := base64.StdEncoding.EncodeToString([]byte(data))
	body := fmt.Sprintf("data=%s", encData)
	resp, err := http.Post(DetectURL, "text/plain", bytes.NewBuffer([]byte(body)))
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
