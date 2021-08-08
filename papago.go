package papago

import (
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

// Detect tries to guess the input language from the given text
func Detect(text string) (Language, error) {
	var lang Language
	text = strings.Replace(text, "\n", "\\n", -1)
	body := strings.NewReader(fmt.Sprintf(detectParams, text))
	req, err := http.NewRequest("POST", detectURL, body)
	req.Header.Set("Timestamp", "1628347197554")
	req.Header.Set("Authorization", "PPG 89ec95ad-1ebf-43eb-acea-653661f0dbe6:mjYzsMUb3v5p36qr/qBC/g==")
	req.Header.Set("Cookie", "NID_SES=AAABi7Old7TjuSxa2xqnbupBiUQTtDvrCoBkyDEKuq0x75IDAKOBo6b0Sc4/diQsgNL6QWSpMkxzAWHnYLP9rbn2deIUNz0fo+OVnqD4ryuQbWaPlI8J+H2Pf85maCPkmBIkC3zr55oNinUIA9q2PmNjoRJ0NolktreZzeFkJKGkanArO/siZHG0L+WzwByfb3wCJd+ECAsQz8e55Y7jZrXmupkQViIRD4XSAFKTlMmW/bLhwV4ENNLjGaJaP/25Jc8O5Sh4QXDG1Jbd3Rdlqu5MTclWwnf/bBnbs0OwaqQF5J6a4rv5/o1J3pPBXaAF/zUL9gNUVOE+2CsUFeYSBfvWBXersQTIlp0sSdZc/YMtklTeWilBzfxPW1LuPi0Np1LK+qkLK4RdXOJoHu2UvrL+K/mtfVoYTHxpiOdsruSsChjZGXnOGfJX4gsys92bc39WtnKL2KNwibPPemLeZaa0dkY3s23XIjhX6UFUZqgGBovLobo7w91NngmHcbBvg6bqoO6AsJ9vHCQDRI6dXzAokec=")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
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

// Translate translates the text from source Language to target Language
func Translate(text string, source Language, target Language, opt TranslateOptions) (string, error) {
	text = strings.Replace(text, "\n", "\\n", -1)
	opt.Source = source.Code()
	opt.Target = target.Code()
	opt.Text = text
	body := strings.NewReader(fmt.Sprintf("%s", opt))
	req, err := http.NewRequest("POST", translateURL, body)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:90.0) Gecko/20100101 Firefox/90.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Device-Type", "pc")
	req.Header.Set("X-Apigw-Partnerid", "papago")
	req.Header.Set("Timestamp", "1628410431800")
	req.Header.Set("Origin", "https://papago.naver.com")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Gpc", "1")
	req.Header.Set("Authorization", "PPG 89ec95ad-1ebf-43eb-acea-653661f0dbe6:HkbHZAAr0gkJBgX1+HxD1A==")
	req.Header.Set("Referer", "https://papago.naver.com/")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "nid_inf=389294109; NID_SES=AAABg9jc70bXZZL81Ybfe9qkezxKLchTl7b/RSGnYAr0B1ZmQrrkyHRy//OduSV2icNh040K6GfI8yG/a1xcLXKs3Q2IezozPiF/7cifXSLdrwG8cCQicPIUmG3l3hcz6HN0mYAhkcsTsetKjVld+gBrw56PBs+lFnlRYGMzBVxaYHcXfqW4Vv3e4xsBd5Z6GditvOzf8ih64MBBpOYmxnHp2/x1UK2bK9fAyDwwpjCRhFcaFjG0WSJu6NF4/wADuepnGUzPkXMXfrtfekxJIqdGPcf2HRSGwJkWbdzZ2mDkecoEtVWRBuSbzvP4SoTvJ+h8F3VE27+wSpTM8Fu0xZ92T201boFAVUGuHiJIfDEqG5B0mqI5XQzMSfZP0CbdH49j65fHr6d9u8pQWtd/UAsAnmYuGFglG5EYR6bXItdAaUKNUWuN5DNl4X0vOkCGtP6Mj8aPmonxh7WMxZjoti2WqHsTF21L/x892hyAWDl1pxCWKEiRBJ8rJtZa45qArDvHq+3beqolHvr25WeIoNgoRRg=; NID_AUT=qwnqOu7GMH7iHfu6JUiGwvBpSg/H0DxdR5Qz1WfEBuhQSOzE98duqnvtYujJq8gQ; NID_JKL=EqJzzBfo7DZd/cxyTExhOTgtVfKsE3RxX1kA45Q7iE0=; JSESSIONID=62EE77D7AFFD634F2237AE506BB04F56; papago_skin_locale=en")
	req.Header.Set("Te", "trailers")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
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
	body := strings.NewReader(fmt.Sprintf("%s", params))
	req, err := http.NewRequest("POST", ttsURL, body)
	req.Header.Set("Timestamp", "1628351433505")
	req.Header.Set("Authorization", "PPG 89ec95ad-1ebf-43eb-acea-653661f0dbe6:YyvcSv8dGXXcREGqEcGeQw==")
	req.Header.Set("Cookie", "NID_SES=AAABj0/3uKsCO1oBeydwj7kskROHQZvK2ZwRJVqcadUMwzdIL6m+z3GEkH7cvlOeu563DtvJ067kpXMt4vOg7D/JH9tOdItne5h8MRsdO3zHvtqz+M89XDXp9V8gfWRAoYX82c2ZY5Qb7MM+RLie8f2mnyLqYT/qhjMZ8X8HhXpENYifbEww91zwXeXNGxHiDXi9yNcPi7oc3dMFOKVl6Wflr92HqoXjbAW1C3SSZ7kGIgGsPaV0lxSE1GFnDg/4JJgjWiDc6PlmTboKPzPscFEYgo0gzt65LXSzrc0VtuBdvh0x0POfuDyta9uFnOON4CJpzW3cjbsgkGwNovBuCfPYisp8vV1MDJAZV6AKXxs5H9m+j5z5NyLFxoVid+iuf6FqdwMFkFPZnaCBJqvYHSdum8SRpsEnjld8WcK4mafCEIneJFcpPS8Wz1UNGTKLGwaeuPmZMY4AGBAAIq6h2bGGlqk9x8FV1etxAofhH3pnNYhKQhD5NM4BobR5Wf2X3rt9Qjpe7K/u/3oH3bPEm5K8HT4=;")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
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
