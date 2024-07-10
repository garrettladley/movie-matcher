FROM golang:1.22-alpine as builder

WORKDIR /app
RUN apk add --no-cache make nodejs npm git

COPY . ./
RUN make install
RUN make build

FROM scratch
COPY --from=builder /app/bin/movie_matcher /movie_matcher

EXPOSE 3000
ENTRYPOINT [ "./movie_matcher" ]