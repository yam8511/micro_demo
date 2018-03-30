FROM yam8511/go-micro

ARG SERVICE_NAME
COPY ./proto /go/src/proto
COPY ./services/${SERVICE_NAME} /go/src/app
WORKDIR /go/src/app

RUN govendor init
RUN govendor add +external
RUN govendor fetch +missing
RUN go build -o app .

ENTRYPOINT [ "./app" ]