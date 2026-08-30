# prompix

## purpose
prometheus query -> image

publicly display prometheus data without exposing prometheus

## why good
- headless
- lightweight
- server side rendering

## why bad
- still WIP

## how to use
### build
```go build -o prompix cmd/prompix/main.go```
### run
```prompix render --query 'rate(http_requests_total[5m])'```
### help
```prompix render --help```