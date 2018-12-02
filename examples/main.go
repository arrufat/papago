package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/arrufat/papago"
)

func main() {
	fmt.Println("Papago Examples")
	text := "Hello, how are you?"

	// language detection
	sourceLang, err := papago.Detect(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s detected\n", sourceLang)

	// translation
	targetLang := papago.Korean
	fmt.Printf("Translating \"%s\" from %s to %s\n", text, sourceLang, targetLang)
	trans, err := papago.Translate(text, sourceLang, targetLang)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Translation: %s\n", trans)

	// text to speech
	voice := papago.Voice{Language: targetLang, Gender: papago.Female, Speed: papago.Normal}
	tts, err := papago.TTS(trans, voice)
	if err != nil {
		panic(err)
	}
	// file download
	fmt.Printf("Downloading file from:\n\t%s\n", tts)
	fileDest := "/tmp/papago.mp3"
	if err := downloadFile(fileDest, tts); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("Audio file downloaded to %s\n", fileDest)
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
