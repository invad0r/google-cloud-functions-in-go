package api

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/pubsub"
)

func Send(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	topic := client.Topic("randomNumbers")

	rand.Seed(time.Now().UnixNano())

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(strconv.Itoa(rand.Intn(1000))),
	})

	id, err := res.Get(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}
