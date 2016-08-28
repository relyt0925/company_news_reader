FROM golang

ENV GOPATH /go

RUN go get github.com/relyt0925/company_news_reader/...

CMD go run /go/src/github.com/relyt0925/company_news_reader/main/main.go