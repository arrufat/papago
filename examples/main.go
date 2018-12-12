package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/arrufat/papago"
)

func main() {
	source := flag.String("source", "", "code for the source language")
	target := flag.String("target", "", "code for the target language")
	list := flag.Bool("list", false, "list all possible language codes")
	output := flag.String("output", "", "path to the output file")
	flag.Parse()

	if *list {
		fmt.Println("Supported language codes:")
		for i, lang := range papago.SupportedLanguages() {
			fmt.Printf("%2d: %s => %s\n", i+1, lang, lang.Code())
		}
		return
	}

	if *target == "" {
		fmt.Println("Specify at least a target language")
		return
	}

	// use the unparsed arguments as the input text
	text := strings.Join(flag.Args(), " ")
	if text == "" {
		fmt.Println("Specify some text to translate")
		return
	}
	// perform language detection if not source language not specified
	var sourceLang papago.Language
	if *source == "" {
		detectedLang, err := papago.Detect(text)
		if err != nil {
			fmt.Println(err)
			return
		}
		sourceLang = detectedLang
	} else {
		parsedLang, err := papago.ParseLanguageCode(*source)
		if err != nil {
			fmt.Println(err)
			return
		}
		sourceLang = parsedLang
	}
	// get the target language
	targetLang, err := papago.ParseLanguageCode(*target)
	if err != nil {
		fmt.Println(err)
		return
	}

	// translation
	fmt.Printf("Translating \"%s\" from %s to %s\n", text, sourceLang, targetLang)
	trans, err := papago.Translate(text, sourceLang, targetLang)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Translation: %s\n", trans)

	// text to speech
	voice := papago.Voice{
		Language: targetLang,
		Gender:   papago.Female,
		Speed:    papago.Normal,
		Pitch:    papago.Medium,
	}
	tts, err := papago.TTS(trans, voice)
	if err != nil {
		fmt.Println(err)
		return
	}
	if *output == "" {
		fmt.Printf("Audio file available here:\n\t%s\n", tts)
		return
	}

	// file download
	fmt.Printf("Downloading file from:\n\t%s\n", tts)
	if err := downloadFile(*output, tts); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Audio file downloaded to %s\n", *output)
}

func downloadFile(filepath string, url string) error {
	// create file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
