FROM iron/go:dev
WORKDIR /app

RUN mkdir /app
ADD . /app

RUN go build -o test .
CMD ["/app/test"]