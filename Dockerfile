FROM resin/raspberry-pi-golang:latest

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p $GOPATH/src/github.com/stmu/co2meter
COPY . $GOPATH/src/github.com/stmu/co2meter

RUN cd $GOPATH/src/github.com/stmu/co2meter
WORKDIR  $GOPATH/src/github.com/stmu/co2meter

RUN go get
RUN make build_pi

CMD ./co2meter /dev/hidraw0
