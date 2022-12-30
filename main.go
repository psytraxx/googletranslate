package main

import (
	"flag"
	"fmt"
	"github.com/psytraxx/googletranslate/cli"
	"strings"
	"sync"
)

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "auto", "source language[en]")
	flag.StringVar(&targetLang, "t", "de", "target language[de]")
	flag.StringVar(&sourceText, "st", "", "text to translate")
}
func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		SourceText: sourceText,
		TargetLang: targetLang,
	}

	strChan := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go cli.RequestTranslate(reqBody, strChan, wg)

	processedStr := strings.ReplaceAll(<-strChan, " + ", " ")

	fmt.Printf("%s\n", processedStr)
	close(strChan)
	wg.Wait()
}
