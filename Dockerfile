FROM alpine
ADD micro-example-srv /micro-example-srv
ENTRYPOINT [ "/micro-example-srv" ]
