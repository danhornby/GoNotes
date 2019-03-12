FROM golang AS builder
ADD . /go/src/notes
WORKDIR /go/src/notes
RUN go get -d -v github.com/gorilla/mux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/src/notes1

FROM scratch
COPY --from=builder /go/src/notes1 /go/bin/notes
ENTRYPOINT ["/go/bin/notes"]
