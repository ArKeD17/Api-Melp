FROM mvochoa/gomon:1.11-alpine3.8 AS build

WORKDIR /go/src/gitlab.com/api
COPY ./ ./
RUN apk --no-cache add ca-certificates
RUN rm $(find . -name testing.go) && go install

FROM alpine:3.8
ENV PORT 80
ENV HOME /var/api

WORKDIR /var/api
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/api /usr/local/bin/
COPY ./assets /var/api/assets
RUN chmod +x /usr/local/bin/api \
    && mkdir -p /var/api/log

EXPOSE 80
CMD ["api", ">", "/var/api/log/api_$(date +%Y_%m_%d).log 2>&1"]
