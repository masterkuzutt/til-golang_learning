package main

import (
	"fmt"
	"image"
	"local/utils/stringutil"
)

func main() {

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.ColorModel())
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0))
	fmt.Println(m.At(0, 0).RGBA())
	fmt.Println(stringutil.Reverse("Helllo"))
}
