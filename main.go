package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	start := time.Now()
	bounds := screenshot.GetDisplayBounds(0)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	name := os.Args[1]

	fileName := name + ".png"
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	fmt.Println(fileName)

	fmt.Println(time.Since(start))
}
