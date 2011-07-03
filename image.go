package main

import (
    "http"
    "image"
    "image/png"
)

type Image interface {
    // ColorModel returns the Image's ColorModel.
    ColorModel() ColorModel
    // Bounds returns the domain for which At can return non-zero color.
    Bounds() Rectangle
    // At returns the color of the pixel at (x, y).
    At(x, y int) Color
}

var m image.Image

func img(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "image/png")
    png.Encode(w, m)
}

func main() {
    rgba := image.NewRGBA(300, 300)
    for x := 0; x < 300; x++ {
        for y := 0; y < 300; y++ {
            rgba.Set(x, y, image.RGBAColor{0, 0, 0, 255})
        }
    }
    m = rgba
    http.HandleFunc("/", img)
    http.ListenAndServe(":8080", nil)
}

/* - Final code
package main

import (
    "http"
    "image"
    "image/png"
)

var m image.Image

func img(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "image/png")
    png.Encode(w, m)
}

type MyImage struct {
    width, height int
}

func (m MyImage) ColorModel() image.ColorModel {
    return image.RGBAColorModel
}

func (m MyImage) Bounds() image.Rectangle {
    return image.Rect(0, 0, m.width, m.height)
}

func (m MyImage) At(x, y int) image.Color {
    c := uint8(x*x + y*y)
    return image.RGBAColor{c, c, c, 255}
}

func main() {
    m = MyImage{300, 300}
    http.HandleFunc("/", img)
    http.ListenAndServe(":8080", nil)
}

*/