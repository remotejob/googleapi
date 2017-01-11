package google_translate

import (
	"context"

	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func Service(lang string, text string) string {
	ctx := context.Background()

	// Creates a client
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// The target language
	target, err := language.Parse(lang)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates some text into Russian
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	log.Println(translations[0].Text)

	return translations[0].Text
}
