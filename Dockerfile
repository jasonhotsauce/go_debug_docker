FROM base:latest AS builder
RUN go build -o /demo_release ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates \
    libc6-compat
COPY --from=builder /demo_release /
CMD [ "/demo_release" ]
LABEL Name=docker_go_demo Version=0.0.1