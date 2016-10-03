docker exec fileServe go build -o /go/fileServe && docker exec fileServe cp /go/fileServe /go/src/app/app && sudo chmod +x ./app && echo 'run' && ./app
