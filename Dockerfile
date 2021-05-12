FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o math

CMD [ "./math" ]