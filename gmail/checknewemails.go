package gmail

import (
	"context"
	"log"

	gmail "google.golang.org/api/gmail/v1"
)

func CheckNew() {

	ctx := context.Background()

	client, err := gmail.New(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

}
