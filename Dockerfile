FROM alpine:latest
RUN apk update
RUN apk add --no-cache go
RUN cd /root/ && mkdir adsense_server && cd adsense_server
COPY main.go /root/adsense_server
EXPOSE 8080/tcp
RUN cd /root/adsense_server && go build
CMD /root/adsense_server/adsense_server
