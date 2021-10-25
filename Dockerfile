FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go test ./... -v

RUN go build -o ./ys-project ./cmd/

EXPOSE 8080

CMD [ "./ys-project" ]