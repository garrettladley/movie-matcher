FROM golang:1.22-alpine as builder

WORKDIR /app
RUN apk add --no-cache make nodejs npm git

COPY . ./
RUN make install
RUN make build-prod

FROM scratch
COPY --from=builder /app/config/ /config/
COPY --from=builder /app/bin/movie_matcher /movie_matcher
COPY --from=builder /app/config/ /config/
ENV APP_ENVIRONMENT production

ENTRYPOINT [ "./movie_matcher" ]