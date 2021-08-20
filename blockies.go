package blockies

import (
	"bytes"
	b64 "encoding/base64"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"os"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/transform"
)

// RGB godoc
type RGB struct {
	R, G, B float64
}

var randseed = [4]int32{}

func seedrand(seed string) {
	for i := 0; i < len(seed); i++ {
		randseed[i%4] = ((randseed[i%4] << 5) - randseed[i%4]) + int32(seed[i])
	}
}

func rand() float64 {
	var t = randseed[0] ^ (randseed[0] << 11)
	randseed[0] = randseed[1]
	randseed[1] = randseed[2]
	randseed[2] = randseed[3]
	randseed[3] = (randseed[3] ^ (randseed[3] >> 19) ^ t ^ (t >> 8))
	return float64(uint32(randseed[3])>>0) / float64(uint32((1<<31))>>0)
}

func hslToRgb(h float64, s float64, l float64) RGB {
	// Must be fractions of 1
	s /= 100
	l /= 100

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod((h/60), 2)-1))
	m := l - c/2
	r := 0.0
	g := 0.0
	b := 0.0

	if 0 <= h && h < 60 {
		r = c
		g = x
		b = 0
	} else if 60 <= h && h < 120 {
		r = x
		g = c
		b = 0
	} else if 120 <= h && h < 180 {
		r = 0
		g = c
		b = x
	} else if 180 <= h && h < 240 {
		r = 0
		g = x
		b = c
	} else if 240 <= h && h < 300 {
		r = x
		g = 0
		b = c
	} else if 300 <= h && h < 360 {
		r = c
		g = 0
		b = x
	}
	r = math.Round((r + m) * 255)
	g = math.Round((g + m) * 255)
	b = math.Round((b + m) * 255)
	return RGB{r, g, b}
}

func createColor() RGB {
	var h = math.Floor(rand() * 360)
	var s = ((rand() * 60) + 40)
	var l = ((rand() + rand() + rand() + rand()) * 25)
	return hslToRgb(h, s, l)
}

func createImageData(size int) []int {
	var width = size
	var height = size
	var dataWidth = int(math.Ceil(float64(width / 2)))
	var mirrorWidth = width - dataWidth
	var data = []int{}

	for y := 0; y < height; y++ {
		var row = []int{}
		for x := 0; x < dataWidth; x++ {
			row = append(row, int(math.Floor(float64(float64(rand())*float64(2.3)))))
		}
		var r = row[0:mirrorWidth]
		var reverse = []int{}
		j := 0
		for i := len(r) - 1; i >= 0; i-- {
			reverse = append(reverse, r[i])
			j++
		}
		row = append(row, reverse...)
		for i := 0; i < len(row); i++ {
			data = append(data, row[i])
		}
	}
	return data
}

// Option godoc
type Option struct {
	Seed      string
	Size      int
	Scale     int
	Color     RGB
	BgColor   RGB
	SpotColor RGB
}

func buildOpts(opts Option) Option {
	// || math.Floor((math.random() * math.pow(10, 16))).toString(16),
	seedrand(opts.Seed)
	var newOption = Option{
		Seed:      opts.Seed,
		Size:      checkIfZero(opts.Size, 8),
		Scale:     checkIfZero(opts.Scale, 4),
		Color:     createColor(),
		BgColor:   createColor(),
		SpotColor: createColor(),
	}
	return newOption
}

func checkIfZero(value int, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}

// RenderIcon godoc
func RenderIcon(opts Option) (string, error) {
	randseed = [4]int32{}
	if opts.Seed == "" {
		return opts.Seed, errors.New("Seed is mandatory parameter")
	}

	opts = buildOpts(opts)
	var imageData = createImageData(opts.Size)
	var width = int(math.Sqrt(float64(len(imageData))))
	RGB := opts.BgColor
	background := color.NRGBA{uint8(RGB.R), uint8(RGB.G), uint8(RGB.B), 255}
	img := createImage(opts.Size, opts.Size, background)
	for i := 0; i < len(imageData); i++ {
		if imageData[i] != 0 {
			var y = int(math.Floor(float64(i / width)))
			var x = int(i % width)
			c := color.NRGBA{}
			if imageData[i] == 1 {
				RGB := opts.Color
				c = color.NRGBA{
					uint8(RGB.R), uint8(RGB.G), uint8(RGB.B), 255}
			} else {
				RGB := opts.SpotColor
				c = color.NRGBA{
					uint8(RGB.R), uint8(RGB.G), uint8(RGB.B), 255}
			}
			img.Set(x, y, c)
		}
	}

	inverted := effect.UnsharpMask(img, 9, 0)
	resized := transform.Resize(inverted, opts.Size*opts.Scale, opts.Size*opts.Scale, transform.Box)

	var buf bytes.Buffer
	png.Encode(&buf, resized)
	sEnc := b64.StdEncoding.EncodeToString([]byte(buf.Bytes()))
	f, err := os.Create(opts.Seed + "_.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, resized, nil)
	return sEnc, nil
}

func createImage(width int, height int, background color.NRGBA) *image.RGBA {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)
	return img
}
