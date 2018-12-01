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
	trans, err := papago.Translate(text, papago.English, papago.Korean)
	if err != nil {
		panic(err)
	}
	fmt.Println(trans)

	tts, err := papago.TTS(trans, papago.English, papago.Male)
	if err != nil {
		panic(err)
	}
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
