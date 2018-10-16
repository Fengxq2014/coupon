FROM golang
WORKDIR src/github.com/Fengxq2014
RUN go get -u -v github.com/golang/dep/cmd/dep
RUN git clone https://github.com/Fengxq2014/coupon.git
WORKDIR coupon
RUN dep ensure
ADD .env .env
EXPOSE 9000
ENTRYPOINT ["go","run","/go/src/github.com/Fengxq2014/coupon/main.go","/go/src/github.com/Fengxq2014/coupon/router.go","/go/src/github.com/Fengxq2014/coupon/logMiddleware.go"]