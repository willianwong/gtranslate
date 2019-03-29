package main

import (
	"fmt"

	"github.com/willianwong/gtranslate"
)

func main() {
	text := "Hello World"
	translated, err := gtranslate.TranslateWithFromTo(
		text,
		gtranslate.FromTo{
			From: "en",
			To:   "zh-CN",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | zh-CN: %s \n", text, translated)
	// en: Hello World | ja: こんにちは世界
}
