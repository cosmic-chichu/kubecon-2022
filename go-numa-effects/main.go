package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/numaproj/numaflow-go/pkg/mapper"
	"go-numa-effects/effects"
	"image/jpeg"
	"log"
)

type Input struct {
	Email string `json:"email"`
	Value []byte `json:"value"`
}

func (f *Input) Map(ctx context.Context, keys []string, d mapper.Datum) mapper.Messages {
	// directly forward the input to the output
	val := d.Value()
	eventTime := d.EventTime()
	_ = eventTime
	watermark := d.Watermark()
	_ = watermark

	var resultKeys = keys
	var resultVal = getMsgBytes(val)

	return mapper.MessagesBuilder().Append(mapper.NewMessage(resultVal).WithKeys(resultKeys))
}

func getMsgBytes(input []byte) []byte {
	var incomingData Input
	err := json.Unmarshal(input,
		&incomingData)
	if err != nil {
		e := fmt.Errorf("unable to unmarshal incoming data: %v",
			err)
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
	err = jpeg.Encode(&outBytes,
		outImg.Img,
		nil)
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
	return effects.RunPencil(img,
		1)
}

func main() {
	err := mapper.NewServer(&Input{}).Start(context.Background())
	if err != nil {
		log.Panic("Failed to start map function server: ",
			err)
	}
}
