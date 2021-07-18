FROM golang:1.16

WORKDIR /code
COPY . .

RUN go get .

# TODO: use go build
CMD [ "go", "run", "main.go" ]
