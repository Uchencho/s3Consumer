package main

import (
	"github.com/Uchencho/s3Consumer/internal/app"

	"github.com/Uchencho/commons/aws/storage"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	a := app.New()
	s3ConsumerAdapter := storage.NewConsumerAdapter(a.Consumer())
	lambda.Start(s3ConsumerAdapter.Consume)
}
