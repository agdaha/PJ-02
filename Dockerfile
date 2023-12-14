FROM golang
WORKDIR /go/src/PJ-02
COPY . ./
RUN go install .

FROM alpine:latest
LABEL version="1.0.0"
LABEL maintainer="Rishat Tarisov<r.tarisov@gmail.com>"
WORKDIR /root/
COPY --from=0 /go/bin/PJ-02 .
ENTRYPOINT ./PJ-02
