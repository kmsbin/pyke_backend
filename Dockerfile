FROM golang:1.14.4
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN cd app && go build -o main .
CMD ["/app/app/main"]