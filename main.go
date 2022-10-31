package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

var invokeCount = 0
func init() {
	fmt.Println("init() function has been called")
}

type MyEvent struct {
	Name string `json:"What is your name?"`
    Age int     `json:"How old are you?"`
}

type MyResponse struct {
    Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	invokeCount += 1
    return MyResponse{Message: fmt.Sprintf("%s is %d years old! count: %d", event.Name, event.Age, invokeCount)}, nil
}

func main() {
	fmt.Println("Go, good to see you again !")
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(HandleLambdaEvent)
}