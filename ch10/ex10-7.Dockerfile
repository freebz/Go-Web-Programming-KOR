# 예제 10-7. 간단한 웹 서비스에 대한 도커 파일

FROM golang

ADD . /go/src/github.com/sausheong/ws-d
WORKDIR /go/src/github.com/sausheong/ws-d

RUN go get github.com/lib/pq
RUN go get install github.com/sausheong/ws-d

ENTRYPOINT /go/bin/ws-d

EXPORSE 8080
