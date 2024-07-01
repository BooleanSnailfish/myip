FROM golang:latest

WORKDIR /app

ADD . /app/
ADD ./GeoIP2-City.mmdb /app/

RUN go build -o server myip/*.go

CMD /app/server