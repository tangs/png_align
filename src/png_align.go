package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func max2(x, y int) (int, int) {
	if x < y {
		return y, y
	}
	return x, x
}

func alignNum(num int) int {
	dest := 1
	for ;dest < num; dest <<= 1{
	}
	return dest
}

func getAlignSize(w int, h int) (int, int) {
	w, h = alignNum(w), alignNum(h)
	w, h = max2(w, h)
	return w, h
}

func alignPng(path string, destPath string) {
	fmt.Println("align png(src dest):", path, destPath)
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	//log.Println(f)
	defer f.Close()
	if img, _, err := image.Decode(bufio.NewReader(f)); err == nil {
		fmt.Println("src w h:", img.Bounds().Dx(), img.Bounds().Dy())
		width, height := getAlignSize(img.Bounds().Dx(), img.Bounds().Dy())
		newImg := image.NewNRGBA(image.Rect(0, 0, width, height))
		fmt.Println("dest w h:", width, height)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				newImg.Set(x, y, img.At(x, y))
			}
		}
		outFile, err := os.Create(destPath)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()

		f.Close()
		if png.Encode(outFile, newImg); err != nil {
			log.Fatal(err)
		}
		return;
	} else {
		log.Fatal(err)
	}
}

func main() {
	for idx := 1; idx < len(os.Args); idx += 2 {
		alignPng(os.Args[idx], os.Args[idx + 1])
	}
	//alignPng("5.png", "5.png")
}
