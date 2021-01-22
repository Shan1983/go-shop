# specify the base image to use for the application
FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/Shan1983/go-shop/cmd/app/
RUN cd /build  && git clone https://github.com/Shan1983/go-shop.git

RUN cd /build/go-shop/cmd/app && go build

EXPOSE 8080

ENTRYPOINT [ "/build/go-shop/cmd/app/main" ]