FROM alpine:3.12

RUN apk --no-cache add ca-certificates

COPY gasper /bin/

ENTRYPOINT ["/bin/gasper"]