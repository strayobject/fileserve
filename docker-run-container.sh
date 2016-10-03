docker rm -f fileServe && docker run -itd -p8081:8080 -v "$PWD":/go/src/app --name fileServe fileserve
