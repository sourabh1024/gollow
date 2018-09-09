FROM golang:1.8.3 as gollow

# load config json
CMD if [ ${APP_ENV} = production ]; \
        then \
        ENV GOLLOW_CF ../config.development.json; \
        else \
        ENV GOLLOW_CF ../config.json; \
        fi

WORKDIR /go/src/github.com/sourabh1024/gollow
COPY . .

EXPOSE 7777
#EXPOSE 7778
#EXPOSE 2222
EXPOSE 2223

ENV MYSQL_HOST db
RUN ls
RUN pwd
#RUN go get ./
RUN go install ./
RUN go build -o main .

CMD ["./main"]