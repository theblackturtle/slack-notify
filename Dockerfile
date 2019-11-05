FROM alpine:latest as downloader

RUN apk --no-cache add wget && rm -rf /var/cache/apk/*

RUN wget https://github.com/theblackturtle/github-releases/releases/download/v0.1.1/github-releases_linux_amd64
RUN chmod +x github-releases_linux_amd64
RUN ./github-releases_linux_amd64 --user theblackturtle --project slack-notify --kernel linux --architecture amd64 --download
RUN rm github-releases_linux_amd64
RUN mv slack-notify_linux_amd64 /slack-notify && chmod +x /slack-notify

FROM alpine:latest
COPY --from=downloader /slack-notify /usr/local/bin/slack-notify
WORKDIR /
ENTRYPOINT ["slack-notify"]
