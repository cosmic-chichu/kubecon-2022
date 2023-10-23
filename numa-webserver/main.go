package main

import (
	"bytes"
	"crypto/tls"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
	"io"
	"net/http"
	"time"
)

var router *gin.Engine

var (
	//go:embed assets/* templates/*
	f      embed.FS
	webcam *gocv.VideoCapture
	stream *mjpeg.Stream
	err    error
)

type Datum struct {
	Email string `json:"email"`
	Value []byte `json:"value"`
}

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.Static("/assets",
		"./assets")
	router.LoadHTMLGlob("templates/*")

	// open webcam
	webcam, err = gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Printf("Error opening capture device: %v\n",
			0)
		return
	}
	defer webcam.Close()

	// create the mjpeg stream
	stream = mjpeg.NewStream()
	// start capturing
	go mjpegCapture()

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/",
		func(c *gin.Context) {

			// Call the HTML method of the Context to render a template
			c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"index.html",
				// Pass the data that the page uses (in this case, 'title')
				gin.H{
					"title": "Numaproj",
				},
			)

		})

	router.GET("/favicon.ico",
		func(context *gin.Context) {
			file, _ := f.ReadFile("assets/numaproj.svg")
			context.Data(
				http.StatusOK,
				"image/x-icon",
				file,
			)
		})

	router.GET("/stream",
		func(context *gin.Context) {
			stream.ServeHTTP(context.Writer, context.Request)
		})

	// Start serving the application
	router.Run()
}

func postToPipeline(client *http.Client, imgBytes []byte) {
	res := new(Datum)
	res.Email = "foo"
	res.Value = imgBytes

	opBytes, err := json.Marshal(res)
	if err != nil {
		e := fmt.Errorf("failed to marshal Datum: %v",
			err)
		fmt.Println(e)
	}

	resp, err := client.Post("https://localhost:8444/vertices/input",
		"application/json; charset=UTF-8",
		bytes.NewBuffer(opBytes))
	if err != nil {
		fmt.Println(err)
	}
	if resp != nil {
		io.Copy(io.Discard,
			resp.Body)
		resp.Body.Close()
	}
}

func mjpegCapture() {
	img := gocv.NewMat()
	defer img.Close()

	// create http client
	tr := &http.Transport{
		MaxIdleConns:    8,
		MaxConnsPerHost: 15,
		IdleConnTimeout: 2 * time.Minute,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n",
				0)
		}
		time.Sleep(time.Millisecond * 80)
		if img.Empty() {
			continue
		}

		buf, _ := gocv.IMEncode(".jpg", img)

		opImgBytes := buf.GetBytes()
		postToPipeline(client, opImgBytes)
		stream.UpdateJPEG(opImgBytes)
		buf.Close()
	}
}
