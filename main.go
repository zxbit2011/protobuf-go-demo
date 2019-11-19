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
		var arr = []float64{123.545, 465456465.1155}

		mapData := &innerMap.Map{}
		mapData.Fill = []*innerMap.Fill{
			{
				Geometry: []*innerMap.Geometry{
					{Type: "Polygon", Coordinates: [][]float64{arr}},
					{Type: "Polygon", Coordinates: [][]float64{arr}},
				},
				Properties: []*innerMap.Properties{
					{
						Id:          "07550046B0400001",
						Name:        "",
						Icon:        "",
						X:           12685203.757971337,
						Y:           2574193.582845682,
						Floor:       -5,
						Height:      15,
						Base:        "",
						Color:       "#CCCCCC",
						Opacity:     1,
						BorderColor: "#E0E0E0",
						Layer:       1,
					},
				},
			},
		}
		mapData.Floor = []*innerMap.Floor{
			{
				Geometry: []*innerMap.Geometry{
					{Type: "Polygon", Coordinates: [][]float64{arr}},
					{Type: "Polygon", Coordinates: [][]float64{arr}},
				},
				Properties: []*innerMap.Properties{
					{
						Id:          "07550046B0400001",
						Name:        "",
						Icon:        "",
						X:           12685203.757971337,
						Y:           2574193.582845682,
						Floor:       -5,
						Height:      15,
						Base:        "",
						Color:       "#CCCCCC",
						Opacity:     1,
						BorderColor: "#E0E0E0",
						Layer:       1,
					},
				},
			},
		}
		mapData.Label = []*innerMap.Label{
			{
				Geometry: []*innerMap.Geometry{
					{Type: "Polygon", Coordinates: [][]float64{arr}},
					{Type: "Polygon", Coordinates: [][]float64{arr}},
				},
				Properties: []*innerMap.Properties{
					{
						Id:          "07550046B0400001",
						Name:        "",
						Icon:        "",
						X:           12685203.757971337,
						Y:           2574193.582845682,
						Floor:       -5,
						Height:      15,
						Base:        "",
						Color:       "#CCCCCC",
						Opacity:     1,
						BorderColor: "#E0E0E0",
						Layer:       1,
					},
				},
			},
		}
		// 编码
		data, err := proto.Marshal(mapData)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		// 解码
		newDeviceInfo := &innerMap.Map{}
		err = proto.Unmarshal(data, newDeviceInfo)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		jsonStr, err := json.Marshal(newDeviceInfo)
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}
		return c.String(http.StatusOK, string(jsonStr))
	})
	e.GET("/demo.pbf", func(c echo.Context) error {
		mapData := demo.Demo{}
		data, err := ReadAll("demo.json")
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}
		err = json.Unmarshal(data, &mapData)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("Demo Unmarshal err: %s", err.Error()))
		}

		b, _ := proto.Marshal(&mapData)
		return c.Blob(http.StatusOK, "application/octet-stream", b)
	})
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", 8083)); err != nil {
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
