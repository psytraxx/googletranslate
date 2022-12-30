package main

import (
	"flag"
	"fmt"
	"github.com/psytraxx/googletranslate/cli"
	"strings"
	"sync"
)

var sourceLang string
var targetLanguage string
var sourceText string

var wg sync.WaitGroup

func init() {
	flag.StringVar(&sourceLang, "s", "auto", "source language[en]")
	flag.StringVar(&targetLanguage, "t", "de", "target language[de] (or comma separated)")
	flag.StringVar(&sourceText, "st", "", "text to translate")
}
func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	targetLanguages := strings.Split(targetLanguage, ",")

	strChan := make(chan string, len(targetLanguages))
	defer close(strChan)
	for _, targetLang := range targetLanguages {

		wg.Add(1)
		reqBody := &cli.RequestBody{
			SourceLang: sourceLang,
			SourceText: sourceText,
			TargetLang: targetLang,
		}
		go cli.RequestTranslate(reqBody, strChan, &wg)
		processedStr := strings.ReplaceAll(<-strChan, " + ", " ")
		fmt.Printf("%s\n", processedStr)
	}
	wg.Wait()
}
