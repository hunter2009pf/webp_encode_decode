package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"

	"github.com/chai2010/tiff"
	"github.com/chai2010/webp"
)

func main() {

	var buf bytes.Buffer
	var jpg_img image.Image

	jpg, err := ioutil.ReadFile("jpg_img.jpg")
	if err != nil {
		return
	}

	jpg_img, _ = jpeg.Decode(bytes.NewReader(jpg))

	if jpg_img == nil {
		return
	}

	if err = webp.Encode(&buf, jpg_img, &webp.Options{Lossless: false, Quality: 100}); err != nil {
		return
	}
	if err = ioutil.WriteFile("jpg_img.webp", buf.Bytes(), 0644); err != nil {
		return
	}

	buf.Reset()
	var jpeg_img image.Image
	jpeeg, err1 := ioutil.ReadFile("jpeg_img.jpeg")
	if err1 != nil {
		return
	}

	jpeg_img, _ = jpeg.Decode(bytes.NewReader(jpeeg))

	if jpeg_img == nil {
		return
	}

	if err = webp.Encode(&buf, jpeg_img, &webp.Options{Lossless: false, Quality: 100}); err != nil {
		return
	}
	if err = ioutil.WriteFile("jpeg_img.webp", buf.Bytes(), 0644); err != nil {
		return
	}

	buf.Reset()
	var JPEG_img image.Image
	JPEG, err2 := ioutil.ReadFile("JPEG_UPPERCASE.JPEG")
	if err2 != nil {
		return
	}

	JPEG_img, _ = jpeg.Decode(bytes.NewReader(JPEG))

	if JPEG_img == nil {
		return
	}

	if err = webp.Encode(&buf, JPEG_img, &webp.Options{Lossless: false, Quality: 100}); err != nil {
		return
	}
	if err = ioutil.WriteFile("JPEG_UPPERCASE.webp", buf.Bytes(), 0644); err != nil {
		return
	}

	buf.Reset()
	var png_img image.Image
	PNG, err3 := ioutil.ReadFile("def.PNG")
	if err3 != nil {
		return
	}

	png_img, _ = png.Decode(bytes.NewReader(PNG))

	if png_img == nil {
		return
	}

	if err = webp.Encode(&buf, png_img, &webp.Options{Lossless: false, Quality: 100}); err != nil {
		return
	}
	if err = ioutil.WriteFile("def.webp", buf.Bytes(), 0644); err != nil {
		return
	}

	buf.Reset()
	var tiff_img image.Image
	TIFF, err3 := ioutil.ReadFile("tiff2.tiff")
	if err3 != nil {
		fmt.Println("line102 error: ", err3)
		return
	}

	tiff_img, errbig := tiff.Decode(bytes.NewReader(TIFF))
	if errbig != nil {
		fmt.Println("errbig is ", errbig)
	}

	if tiff_img == nil {
		fmt.Println("line109 nil")
		return
	}

	if err = webp.Encode(&buf, tiff_img, &webp.Options{Lossless: false, Quality: 100}); err != nil {
		fmt.Println("line114 error: ", err)
		return
	}
	if err = ioutil.WriteFile("tiff_img.webp", buf.Bytes(), 0644); err != nil {
		fmt.Println("line118 error: ", err)
		return
	}

}
