package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *MyEvent){
	// if event == nil {
	// 	return nil, fmt.Errorf("received nil event")
	// }
	// message := fmt.Sprintf("Hello %s!", event.Name)
	// return &message, nil

	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.GET("/hello-world", HelloWorld)
	// message := "Hello!"
	// fmt.Println(message)
	// return &message, nil
}

func main() {
	fmt.Println("main start")
	lambda.Start(HandleRequest)
}

func HelloWorld(c echo.Context) error {
	c.Logger().Info("Hello World!")
	return c.JSON(http.StatusOK, CreateResponse("Hello World!"))
}

func CreateResponse(msg string) any {
	return struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
}
