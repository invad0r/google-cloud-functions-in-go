# Google Cloud Functions

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
```

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