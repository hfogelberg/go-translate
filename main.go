package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/translate"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

var text = flag.String("x", " ", "Text to translate")
var tl = flag.String("t", "en", "Language to translate too")

func main() {
	flag.Parse()
	ctx := context.Background()

	fmt.Println(*text)
	fmt.Println(*tl)

	apiKey := os.Getenv("GOOGLE_API_KEY")
	log.Printf("Key: %s", apiKey)
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	target, err := language.Parse(*tl)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	translations, err := client.Translate(ctx, []string{*text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v\n", translations[0].Text)
}
