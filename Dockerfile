FROM alpine:latest

RUN apk --no-cache add wget tar && rm -rf /var/cache/apk/*

RUN wget https://github.com/theblackturtle/slack-notify/releases/download/v0.1.0/slack-notify_v0.1.0_Linux_x86_64.tar.gz
RUN tar zxf slack-notify_v0.1.0_Linux_x86_64.tar.gz
RUN rm slack-notify_v0.1.0_Linux_x86_64.tar.gz
RUN chmod +x slack-notify && mv slack-notify /usr/bin/slack-notify

ENTRYPOINT ["slack-notify"]
