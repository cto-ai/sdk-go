FROM registry.cto.ai/official_images/base:latest as dep
FROM golang:1.16 as mock-build

ARG SDK_TAG=v2.2.3

COPY --from=dep /bin/sdk-daemon /bin/.
COPY --from=dep /bin/interaction-server /bin/.

RUN ls /bin 


RUN go get github.com/cto-ai/sdk-go/v2@${SDK_TAG}

WORKDIR /build

ADD . .

RUN go build -o /bin/mock-api -i mock-api.go

RUN go build -o /bin/verify -i sdk-compatibility-testing/verify-exemplar.go

RUN go build -o testsdk -i sdk-compatibility-testing/golang/main.go

CMD /bin/verify /build/testsdk
