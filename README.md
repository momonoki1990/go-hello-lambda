## Getting Started

```
$ go build
$ GOOS=linux GOARCH=amd64 go build -o main main.go
$ zip main.zip main

# After integration of lambda to api gateway(path is /hellogolang, method is POST)
$ curl -X POST -H "Content-Type: application/json" -d '{"What is your name?": "Nick", "How old are you?": 30 }' https://{api id}.execute-api.{aws region}.amazonaws.com/hellogolang

```

## Reference

[Go の AWS Lambda 関数ハンドラー - AWS Lambda](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-handler.html)
