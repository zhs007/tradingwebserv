FROM golang:1.10 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR $GOPATH/src/github.com/zhs007/tradingwebserv

COPY ./Gopkg.* $GOPATH/src/github.com/zhs007/tradingwebserv/

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -vendor-only -v

COPY . $GOPATH/src/github.com/zhs007/tradingwebserv

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tradingwebserv . \
    && mkdir /home/tradingwebserv \
    && mkdir /home/tradingwebserv/cfg \
    && cp ./tradingwebserv /home/tradingwebserv/ \
    && cp ./cfg/config.yaml.default /home/tradingwebserv/cfg/config.yaml \
    && cp -r ./www /home/tradingwebserv/www

FROM scratch
WORKDIR /home/tradingwebserv
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /home/tradingwebserv /home/tradingwebserv
CMD ["./tradingwebserv"]
