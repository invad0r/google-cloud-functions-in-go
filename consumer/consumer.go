package consumer

import (
	"context"
	"log"
)

type event struct {
	Data []byte
}

func Consume(ctx context.Context, e event) error {
	log.Printf("data: %v", e.Data)
	return nil
}
