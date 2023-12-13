FROM golang
RUN mkdir -p /go/src/PJ-02
RUN mkdir -p /go/src/PJ-02/buffer
RUN mkdir -p /go/src/PJ-02/pipeline
WORKDIR /go/src/PJ-02
COPY . ./
RUN go install .

FROM alpine:latest
LABEL version="1.0.0"
LABEL maintainer="Rishat Tarisov<r.tarisov@gmaol.com>"
WORKDIR /root/
COPY --from=0 /go/bin/PJ-02 .
ENTRYPOINT ./PJ-02
