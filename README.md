Panopticon
===

Parental control for Linux

## Contribute

### Generate API

```
protoc -I ./proto/ --go_out=./pkg/api --go_opt=module=github.com/ylallemant/panopticon --go-grpc_out=./pkg/api --go-grpc_opt=module=github.com/ylallemant/panopticon proto/v1/*.proto

```