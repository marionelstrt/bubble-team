FROM golang:1.25-alpine

RUN apk add entr \
	build-base # for sqlite3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

CMD ["sh", "-c", \
	"find -name '*.go' | entr -nr go run ."]
