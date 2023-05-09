FROM alpine:latest
LABEL authors="rainerosion@gmail.com"
COPY upcos /bin/upcos
ENTRYPOINT ["upcos"]