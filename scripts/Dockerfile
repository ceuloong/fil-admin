FROM alpine

COPY ./fil-admin /
EXPOSE 8000

CMD ["/go-admin","server","-c", "/config/settings.yml"]
