FROM golang:24.3=alpine as stage1

WORDIR /app

COPY go.mod go.sum ./app

RUN go mod dowloand

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

##########################################################

FROM golang scratch

COPY --from=stage1 /app/main /

ENTRYPOINT ["/main"]