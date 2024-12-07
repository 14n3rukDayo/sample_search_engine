# only for dev
FROM golang:1.23

WORKDIR /src/app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/google/wire/cmd/wire@latest

COPY . .
RUN go mod download
EXPOSE 8000
CMD ["air", "-c", ".air.toml"]