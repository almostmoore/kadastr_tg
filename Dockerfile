FROM alpine:3.5

RUN apk add --no-cache ca-certificates && update-ca-certificates

ADD kadastr_tg /
RUN chmod +x /kadastr_tg

CMD ./kadastr_tg tg -tgtoken $TG_TOKEN -addr $ADDR