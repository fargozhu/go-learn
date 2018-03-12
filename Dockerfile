FROM iron/go
WORKDIR /app

ENV SRC_DIR=/go/src/github.com/fargozhu/go-learn
ADD . $SRC_DIR
RUN cd $SRC_DIR; go build -o test; cp test /app/
ENTRYPOINT [ "./test" ]