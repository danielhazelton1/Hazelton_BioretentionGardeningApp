FROM golang:1.16.3

## Create workdir and add all relevant files to it
RUN mkdir /app
ADD ./server /app
WORKDIR /app

## Download dependencies listed in the go module
RUN go mod download

## Compile the code
RUN go build -o main .

## Run the webserver
CMD ["/app/main"]
