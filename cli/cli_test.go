package cli

import (
	"strings"
	"sync"
	"testing"
)

func TestRequestTranslate(t *testing.T) {

	t.Run("Translate", func(t *testing.T) {
		wg := &sync.WaitGroup{}
		strChan := make(chan string)

		body := &RequestBody{
			SourceLang: "en",
			SourceText: "Banana",
			TargetLang: "de",
		}
		wg.Add(1)

		const expected = "Banane"

		go RequestTranslate(body, strChan, wg)

		processedStr := strings.ReplaceAll(<-strChan, " + ", " ")

		if processedStr != expected {
			t.Errorf("RequestTranslate() = %v, want %v", processedStr, expected)
		}

		close(strChan)

		wg.Wait()
	})

}
