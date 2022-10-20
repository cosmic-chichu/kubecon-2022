package effects

import (
	"flag"
	"fmt"
	"os"
)

func RunGaussian(img *Image, kernelSize int, sigma float64) *Image {
	gaussian := NewGaussian(kernelSize, sigma)
	outImg, err := gaussian.Apply(img, 0)
	if err != nil {
		fmt.Println("Failed to apply effect:", err)
		os.Exit(1)
	}
	return outImg
}

func RunSobel(img *Image, threshold int, invert bool) *Image {
	sobel := NewSobel(threshold, invert)
	outImg, err := sobel.Apply(img, 0)
	if err != nil {
		fmt.Println("Failed to apply effect:", err)
		os.Exit(1)
	}
	return outImg
}

func RunPencil(img *Image, blurFactor int) *Image {
	pencil := NewPencil(blurFactor)
	outImg, err := pencil.Apply(img, 3)
	if err != nil {
		fmt.Println("Failed to apply effect:", err)
		os.Exit(1)
	}
	return outImg
}

func RunBrightness(img *Image, offset int) *Image {
	brightness := NewBrightness(offset)
	outImg, err := brightness.Apply(img, 0)
	if err != nil {
		fmt.Println("Failed to apply effect:", err)
		os.Exit(1)
	}
	return outImg
}

func RunOil(img *Image, filterSize, levels int) *Image {

	if filterSize <= 3 {
		fmt.Println("FilterSize must be at least 3")
		os.Exit(1)
	}

	if levels < 1 {
		fmt.Println("Levels must be at least 1")
		os.Exit(1)
	}

	oil := NewOilPainting(filterSize, levels)
	outImg, err := oil.Apply(img, 0)
	if err != nil {
		fmt.Println("Failed to apply effect:", err)
		os.Exit(1)
	}
	return outImg
}

func RunCartoon(img *Image, blurStrength, edgeThreshold, oilFilterSize, oilLevels int) *Image {
	opts := CTOpts{
		BlurKernelSize: blurStrength,
		EdgeThreshold:  edgeThreshold,
		OilFilterSize:  oilFilterSize,
		OilLevels:      oilLevels,
		DebugPath:      "",
	}
	cartoon := NewCartoon(opts)
	outImg, err := cartoon.Apply(img, 0)
	if err != nil {
		fmt.Println("Failed to apply effect:", err)
		os.Exit(1)
	}
	return outImg
}

func RunPixelate(img *Image, blockSize int) *Image {
	pixelate := NewPixelate(blockSize)
	outImg, err := pixelate.Apply(img, 0)
	if err != nil {
		fmt.Println("Failed to apply effect:", err)
		os.Exit(1)
	}
	return outImg
}

//func runEffect(img *Image, effect string) *Image {
//	switch effect {
//	case "brightness":
//		return runBrightness(img)
//	case "cartoon":
//		return runCartoon(img)
//	case "gaussian":
//		return runGaussian(img)
//	case "oil":
//		return runOil(img)
//	case "pencil":
//		return runPencil(img)
//	case "pixelate":
//		return runPixelate(img)
//	case "sobel":
//		return runSobel(img)
//	}
//	return nil
//}

func validateFlags(effect string) {
	switch effect {
	case "brightness":
		if len(flag.Args()) != 3 {
			fmt.Println("The brightness effect requires 3 args, input path, output path offset")
			fmt.Println("Sample usage: goeffects -effect=brightness mypic.jpg mypic-lighten.jpg 200")
			flag.PrintDefaults()
			os.Exit(1)
		}
	case "cartoon":
		if len(flag.Args()) != 6 {
			fmt.Println("The cartoon effect requires 6 args, input path, output path, blurStrength, edgeThreshold, oilBoldness, oilLevels")
			fmt.Println("Sample usage: goeffects -effect=cartoon mypic.jpg mypic-cartoon.jpg 21 40 15 15")
			flag.PrintDefaults()
			os.Exit(1)
		}
	case "gaussian":
		if len(flag.Args()) != 4 {
			fmt.Println("The gaussian effect requires 4 args, input path, output path, kernelSize, sigma")
			fmt.Println("Sample usage: goeffects -effect=gaussian mypic.jpg mypic-gaussian.jpg 9 1")
			flag.PrintDefaults()
			os.Exit(1)
		}
	case "oil":
		if len(flag.Args()) != 4 {
			fmt.Println("The oil effect requires 4 args, input path, output path, filterSize, levels")
			fmt.Println("Sample usage: goeffects -effect=oil mypic.jpg mypic-oil.jpg 5 30")
			flag.PrintDefaults()
			os.Exit(1)
		}
	case "pencil":
		if len(flag.Args()) != 3 {
			fmt.Println("The pencil effect requires 3 args, input path, output path blurFactor")
			fmt.Println("Sample usage: goeffects -effect=pencil mypic.jpg mypic-pencil.jpg")
			flag.PrintDefaults()
			os.Exit(1)
		}
	case "pixelate":
		if len(flag.Args()) != 3 {
			fmt.Println("The pixelate effect requires 3 args, input path, output path, block size")
			fmt.Println("Sample usage: goeffects -effect=pixelate mypic.jpg mypic-pixelate.jpg 12")
			flag.PrintDefaults()
			os.Exit(1)
		}
	case "sobel":
		if len(flag.Args()) != 4 {
			fmt.Println("The sobel effect requires 4 args, input path, output path, threshold invert")
			fmt.Println("Sample usage: goeffects -effect=sobel mypic.jpg mypic-sobel.jpg 100 false")
			flag.PrintDefaults()
			os.Exit(1)
		}
	case "":
		fmt.Println("The effect option is required")
		flag.PrintDefaults()
		os.Exit(1)

	default:
		fmt.Println("Unknown effect option value")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
