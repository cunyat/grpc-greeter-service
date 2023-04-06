FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.20 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

COPY main.go .
COPY greeter.pb.go .
COPY greeter_grpc.pb.go .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o greeter .

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch

WORKDIR /app/

COPY --from=builder /app/greeter /app/greeter

ENTRYPOINT ["/app/greeter"]

