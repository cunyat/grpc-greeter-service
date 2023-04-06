# GRPC Greeter service

This is a really simple grpc service created for testing some integrations.
Provides a Greeter service, which is just a "Hello world" example, and also reflection and health services.

Health and reflection services are enabled by default and can be disabled by command arguments:

```shell
docker run ghcr.io/cunyat/grpc-greeter-service -enable-health=false -enable-reflection=false
```

Also, you can configure host and port where the grpc is listening:

```shell
docker run -p 6565:6565 ghcr.io/cunyat/grpc-greeter-service -port 6565 -host 0.0.0.0
```

### Testing

For simple grpcurl testing:

```shell
grpcurl -plaintext -d='{"name": "cunyat"}' localhost:6565 main.Greeter/SayHello
```

Playing with reflection:

```shell
grpcurl -plaintext localhost:6565 list
grpcurl -plaintext localhost:6565 describe main.Greeter
grpcurl -plaintext localhost:6565 describe main.HelloRequest
```

## Ideas

This is just a silly project for testing grpc services in kubernetes. 
I'll add support for authentication with ssl/tls, but not in my priorities for the moment...
