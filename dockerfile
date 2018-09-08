FROM golang:1.8.3 as gollow

ENV MYSQL_HOST $MYSQL_HOST
ENV MYSQL_PORT $MYSQL_PORT
ENV MYSQL_PASSWORD $MYSQL_PASSWORD
ENV MYSQL_ROOT_PASSWORD $MYSQL_ROOT_PASSWORD
ENV AWS_ACCESS_KEY_ID $AWS_ACCESS_KEY_ID
ENV AWS_SECRET_ACCESS_KEY $AWS_SECRET_ACCESS_KEY

# load config json
CMD if [ ${APP_ENV} = production ]; \
        then \
        ENV GOLLOW_CF ../config.development.json; \
        else \
        ENV GOLLOW_CF ../config.json; \
        fi

#RUN mkdir /gollow
#ADD . /gollow/
#WORKDIR /gollow

#RUN go build -o main .

#CMD ["/gollow/main"]

WORKDIR /go/src/app
COPY . .

RUN go get ./
RUN go build -o main .

CMD ["main"]