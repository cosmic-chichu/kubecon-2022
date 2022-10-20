package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	functionsdk "github.com/numaproj/numaflow-go/pkg/function"
	"github.com/numaproj/numaflow-go/pkg/function/server"
	"go-numa-effects/effects"
	"image/jpeg"
)

type Input struct {
	Email string `json:"email"`
	Value []byte `json:"value"`
}

func handle(ctx context.Context, key string, data functionsdk.Datum) functionsdk.Messages {
	_ = data.EventTime() // Event time is available
	_ = data.Watermark() // Watermark is available

	return functionsdk.MessagesBuilder().Append(functionsdk.MessageToAll(getMsgBytes(data.Value())))
}

func getMsgBytes(input []byte) []byte {
	var incomingData Input
	err := json.Unmarshal(input, &incomingData)
	if err != nil {
		e := fmt.Errorf("unable to unmarshal incoming data: %v", err)
		fmt.Println(e.Error())
		return nil
	}

	opImgBytes := udfTransformImgBytes(incomingData.Value)

	return opImgBytes
}

func udfTransformImgBytes(value []byte) []byte {
	if value == nil {
		return nil
	}

	// load image bytes for transformation
	img, err := effects.LoadImageBytes(value)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// apply effect
	outImg := applyRandomEffect(img)
	if outImg == nil {
		fmt.Println("failed to cartoonize")
		return nil
	}

	var outBytes bytes.Buffer
	err = jpeg.Encode(&outBytes, outImg.Img, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return outBytes.Bytes()
}


func applyRandomEffect(img *effects.Image) *effects.Image {
	// rand.Seed(time.Now().UnixNano())
	// max := 3
	// index := rand.Intn(max)
	// switch index {
	// case 0:
	// 	return effects.RunPencil(img, 1)
	// case 1:
	// 	return effects.RunCartoon(img, 3, 150, 20, 6000)
	// default:
	// 	return effects.RunPixelate(img, 1)
	// }
	return effects.RunPencil(img, 1)
}

func main() {
	server.New().RegisterMapper(functionsdk.MapFunc(handle)).Start(context.Background())
}
