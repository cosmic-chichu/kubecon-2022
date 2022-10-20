package main

import (
	"context"
	"fmt"
	functionsdk "github.com/numaproj/numaflow-go/pkg/function"
	"github.com/numaproj/numaflow-go/pkg/function/server"
	"log"
	"net/http"
	"time"

	"github.com/hybridgroup/mjpeg"
)

var (
	stream   *mjpeg.Stream
)

// Simply return the same msg
func handle(ctx context.Context, key string, data functionsdk.Datum) functionsdk.Messages {
	_ = data.EventTime() // Event time is available
	_ = data.Watermark() // Watermark is available
	updateStreamImg(data.Value())

	logMsg := fmt.Sprintf("Updated stream image at %s", data.EventTime().Format(time.RFC3339))
	return functionsdk.MessagesBuilder().Append(functionsdk.MessageToAll([]byte(logMsg)))
}

func startStreamServer(host string) {
	// create the mjpeg stream
	stream = mjpeg.NewStream()


	fmt.Println("Capturing. Point your browser to " + host)
	// start http server
	http.Handle("/", stream)
	log.Fatal(http.ListenAndServe(host, nil))
}

func updateStreamImg(value []byte) {
	stream.UpdateJPEG(value)
}

func main() {
	host := "0.0.0.0:9898"
	go startStreamServer(host)

	server.New().RegisterMapper(functionsdk.MapFunc(handle)).Start(context.Background())
}
