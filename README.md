# go + lambda + LocalStack + Terraform のサンプル

```sh
# ビルド
GOOS=linux GOARCH=amd64 go build -o main main.go
# zip
zip main.zip main

# docker
make localstack

# terraform
make terraform

# Lambda ができているか確認
aws --endpoint-url=http://localhost:4566 lambda get-function --function-name go-lambda-demo

# Lambda 実行
aws lambda --endpoint-url=http://localhost:4566 invoke --function-name go-lambda-demo result.log

# test
curl --location --request GET "http://localhost:4566/" --raw-data '{"name": "John"}' | jq

aws --endpoint-url=http://localhost:4566 lambda invoke --function-name go-lambda-demo response.json

curl --location --request GET "http://localhost:4566/restapis/$(tflocal output -raw api_gw_id)/test/_user_request_/hello-world"
```
