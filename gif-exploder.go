package main

import (
	"image"
	"image/gif"
	"image/jpeg"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeImage(img image.Image, filename string) {
	workingDir, err := os.Getwd()
	check(err)

	output, err := os.Create(workingDir + "/" + filename + ".jpg")
	check(err)

	jpeg.Encode(output, img, nil)
}

func main() {
	fileName := os.Args[1]

	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	decodedGIF, err := gif.DecodeAll(f)
	check(err)

	for index, image := range decodedGIF.Image {
		writeImage(image, "output"+strconv.Itoa(index))
	}
}
