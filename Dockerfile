FROM golang:1.22.2

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

CMD ["go", "run", "app/main.go"]

# docker build -t my-golang-app .
# docker run -p 8080:8080 -it --rm --name my-running-app my-golang-app
