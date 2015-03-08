FROM golang

ADD ./docker-with-golang /

ENTRYPOINT /docker-with-golang

EXPOSE 8080
