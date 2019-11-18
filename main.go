package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/zxbit2011/protobuf-go-demo/demo"
	"github.com/zxbit2011/protobuf-go-demo/innerMap"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			println("异常：%v", err)
			time.Sleep(1 * time.Minute)
		}
	}()
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	//Gzip压缩
	e.Use(middleware.Gzip())
	//文件清单
	e.Static("/page", "page")
	e.File("/", "page/index.html")
	e.GET("/innerMap.pbf", func(c echo.Context) error {
		data, err := ReadAll("demo.json")
		if err != nil {
			return c.String(http.StatusOK, "读取文件失败")
		}
		mapData := &innerMap.Map{}
		err = proto.Unmarshal([]byte(data), mapData)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("innerMap Unmarshal err: %s", err.Error()))
		}

		b, _ := proto.Marshal(mapData)
		return c.Blob(http.StatusOK, "", b)
	})
	e.GET("/demo.pbf", func(c echo.Context) error {
		mapData := demo.Demo{}
		err := json.Unmarshal([]byte(`{"floor":"楼层","fill":"填充","label":"标签"}`), &mapData)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("Demo Unmarshal err: %s", err.Error()))
		}

		b, _ := proto.Marshal(&mapData)
		return c.Blob(http.StatusOK, "application/octet-stream", b)
	})
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", 8080)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
