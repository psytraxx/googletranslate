package cli

import (
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"sync"
)

type RequestBody struct {
	SourceLang string
	SourceText string
	TargetLang string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", translateUrl, nil)
	query := req.URL.Query()
	query.Add("client", "dict-chrome-ex")

	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("dt", "sp")
	query.Add("dt", "ls")
	query.Add("dj", "1")
	query.Add("q", body.SourceText)

	req.URL.RawQuery = query.Encode()

	if err != nil {
		log.Fatalf("1 There was a problem: %s", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("There was a problem: %s", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("There was a problem: %s", err)
		}
	}(res.Body)

	if res.StatusCode == http.StatusTooManyRequests {
		str <- "You have been rate limited, Try again later."
		wg.Done()
		return
	}

	json, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	translated := gjson.GetBytes(json, "sentences.#.trans")
	textArray := translated.Array()
	str <- textArray[0].Str
	wg.Done()
}
