package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
)

// Stringers
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Errors
type NegativeSqrtError float64

func (e NegativeSqrtError) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, NegativeSqrtError(x)
	}

	z := 1.0
	prev := z

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) < 1e-18 {
			break
		}

		prev = z
	}

	return z, nil
}

// Reader
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for {
		reader := strings.NewReader("A")
		return reader.Read(b)
	}
}

// Rot13Reader
type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (int, error) {
	n, err := rot.r.Read(p)
	if err != nil {
		return 0, err
	}

	for i := range p {
		if p[i] >= 'A' && p[i] <= 'Z' {
			p[i] = 65 + (p[i]+13-65)%26
		} else if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = 97 + (p[i]+13-97)%26
		}
	}

	return n, nil
}

// Images
type Image struct {
	l int
	b int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.l, i.b)
}

func (i Image) At(x, y int) color.Color {
	v := uint16(x ^ y)
	return color.RGBA64{v, v, 255, 255}
}

func main() {
	// stringer exercise
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// errors exercise
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	// reader exercise
	reader.Validate(MyReader{})

	// rot13reader exercise
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, &r)

	if err != nil {
		fmt.Println(err)
	}

	// images exercise
	m := Image{l: 12, b: 10}
	pic.ShowImage(m)
}
