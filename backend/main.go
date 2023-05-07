package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/api/getbeer", getBeer)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello,World!!")
}

func getBeer(c echo.Context) error {
	var ctx = context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	client := dynamodb.NewFromConfig(cfg)

	param := &dynamodb.ScanInput{TableName: aws.String("BEER")}

	output, err := client.Scan(ctx, param)
	if err != nil {
		fmt.Println(err)
		return err
	}

	items, err := json.Marshal(output.Items)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(items))
	return c.String(http.StatusOK, string(items))
}
