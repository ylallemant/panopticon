Panopticon
===

Parental control for Linux

## Contribute

### Generate API

```
protoc -I ./proto/ --go_out=./pkg/api --go_opt=module=github.com/ylallemant/panopticon --go-grpc_out=./pkg/api --go-grpc_opt=module=github.com/ylallemant/panopticon proto/v1/*.proto

```

### Push Image

```
export IMAGE=ghcr.io/ylallemant/panopticon
export YES=1

echo "<GitHub-personal-access-token>" |  docker login ghcr.io -u <username>  --password-stdin

make push-container

```