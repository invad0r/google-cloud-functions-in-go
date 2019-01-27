# Google Cloud Functions

snippets while following allong two examples using google cloud functions:
- [YouTube](https://www.youtube.com/watch?v=RitskkjSih0)
source at [packagemain/15-google-cloud-functions-in-go/api at master · plutov/packagemain · GitHub](https://github.com/plutov/packagemain/tree/master/15-google-cloud-functions-in-go/api)
- [Image Resizing with Go and Cloud Functions – Adil H – Medium](https://medium.com/@didil/image-resizing-with-go-and-cloud-functions-792a47e6473d)

```sh
# setup gcloud
gcloud services enable cloudfunctions.googleapis.com
# Operation "operations/acf.f452964f-7fe4-42e6-84c0-982bcf373cf6" finished successfully.

# create topic
gcloud alpha pubsub topics create randomNumbers
# Created topic [projects/principal-fact-205806/topics/randomNumbers].

# deploy funciton to gcloud
gcloud alpha functions deploy api \
    --entry-point Send \
    --runtime go111 \
    --trigger-http \
    --set-env-vars PROJECT_ID=principal-fact-205806

gcloud functions deploy consumer \
    --entry-point Consume \
    --runtime go111 \
    --trigger-topic=randomNumbers


gcloud functions logs read consumer

# teardown
gcloud functions delete api
gcloud functions delete consumer
gcloud pubsub topics delete randomNumbers
```

```sh
gcloud functions deploy ResizeImage --runtime go111 --trigger-http

curl 'https://us-central1-principal-fact-205806.cloudfunctions.net/ResizeImage?url=https://github.com/didil/gcf-go-image-resizer/raw/master/example/gopherizeme_orig.jpg&width=10&height=10'
```

## go mod
```sh
cd api/
export GO111MODULE=on
go mod init

go mod tidy     # checks which dependencies are used and downloads them#

go mod vendor   # create vendor dir +all deps


go mod graph
go list -m all

go mod why <module> # explain why is needed
```