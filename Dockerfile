FROM golang
WORKDIR src/github.com/Fengxq2014
RUN go get -u -v github.com/golang/dep/cmd/dep
RUN git clone https://github.com/Fengxq2014/coupon.git
WORKDIR coupon
RUN dep ensure
ADD .env .env
EXPOSE 9000
RUN go run *.go