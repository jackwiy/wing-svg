package main

import (
	"github.com/rogpeppe/misc/svg"
	"fmt"
	"image"
	"io/ioutil"
	"bytes"
	"bufio"
	"image/png"
)

func main() {
	b, _ := ioutil.ReadFile("test.svg")
	r := bytes.NewReader(b)
	w, _ := svg.Render(r, image.Point{500, 500})
	var pngb []byte
	buf := bytes.NewBuffer(pngb)

	bw := bufio.NewWriter(buf)
	// Write the image into the buffer
	png.Encode(bw, w)
	bw.Flush()

	fmt.Println(string(pngb))
}
