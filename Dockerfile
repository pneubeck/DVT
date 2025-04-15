FROM golang:1.24.2-alpine3.21 AS build
# Build binary from go source
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN GOOS=linux go build -ldflags="-s" -o /go/bin/app -v .
# Create minimal /etc/passwd wiht appuser
RUN echo "appuser:x:10001:10001:App User:/:/sbin/nologin" > /etc/minimal-passwd
FROM scratch
# Copy binary from build step
COPY --from=build /go/bin/app /go/bin/app
# Create and set nonroot user
COPY --from=build /etc/minimal-passwd /etc/passwd
USER appuser
# Set startup options
EXPOSE 8080
ENTRYPOINT [ "/go/bin/app" ]