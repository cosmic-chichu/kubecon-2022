package main

import (
	"context"
	"fmt"
	"github.com/numaproj/numaflow-go/pkg/mapper"
	"log"
	"net/http"
	"time"

	"github.com/hybridgroup/mjpeg"
)

var (
	stream *mjpeg.Stream
)

type F struct {
}

func (f *F) Map(ctx context.Context, keys []string, d mapper.Datum) mapper.Messages {
	// directly forward the input to the output
	val := d.Value()
	eventTime := d.EventTime()
	_ = eventTime
	watermark := d.Watermark()
	_ = watermark

	updateStreamImg(val)

	var resultKeys = keys
	logMsg := fmt.Sprintf("Updated stream image at %s",
		d.EventTime().Format(time.RFC3339))
	var resultVal = []byte(logMsg)

	return mapper.MessagesBuilder().Append(mapper.NewMessage(resultVal).WithKeys(resultKeys))
}

func startStreamServer(host string) {
	// create the mjpeg stream
	stream = mjpeg.NewStream()

	fmt.Println("Capturing. Point your browser to " + host)
	// start http server
	http.Handle("/",
		stream)
	log.Fatal(http.ListenAndServe(host,
		nil))
}

func updateStreamImg(value []byte) {
	stream.UpdateJPEG(value)
}

func main() {
	host := "0.0.0.0:9898"
	go startStreamServer(host)

	err := mapper.NewServer(&F{}).Start(context.Background())
	if err != nil {
		log.Panic("Failed to start map function server: ",
			err)
	}
}
