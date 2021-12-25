package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"
)

const (
	Width    = 2000
	Height   = Width
	sinRange = math.Pi / 2
)

func main() {
	c := gg.NewContext(int(Width), int(Height))
	plotFunc(c, math.Asin, -0.99, 0.99, -math.Pi/2, math.Pi/2, "arcSin.png")
	plotFunc(c, math.Acos, -0.99, 0.99, 0, math.Pi, "arcCos.png")
	plotFunc(c, math.Atan, -10, 10, -math.Pi/2, math.Pi/2, "arcTan.png")
	plotFunc(c, dAsin, -0.99, 0.99, 0, 2, "DarcSin.png")
	plotFunc(c, dAcos, -0.99, 0.99, 2, 0, "DarcCos.png")
	plotFunc(c, dAtan, -10, 10, 0, 2, "DarcTan.png")
}

func dAsin(x float64) float64 {
	return x*math.Asin(x) + math.Sqrt(1-x*x)
}

func dAcos(x float64) float64 {
	return x*x*math.Acos(x) + math.Asin(x)/4 - x*math.Sqrt(1-x*x)/4
}

func dAtan(x float64) float64 {
	return x*math.Atan(x) - math.Log(1+x*x)
}
func plotFunc(c *gg.Context, fn func(x float64) float64, start, end, fnStart, fnEnd float64, saveAt string) {

	c.SetRGBA(1, 1, 1, 1)
	c.Clear()
	c.SetRGBA(0.3, 0.2, 0.2, 1)
	fmt.Println("Starting fn")
	c.SetLineWidth(4)
	fmt.Println("Started Horizontal Lines")
	for i := 0.0; i < Width; i += Width / 20.0 {
		c.DrawLine(i, 0, i, Height)
	}
	fmt.Println("Started Vertical Lines")
	for i := 0.0; i < Height; i += Width / 20.0 {
		c.DrawLine(0, i, Width, i)
	}
	fmt.Println("Finished making grid and started witing to image")
	c.Stroke()
	fmt.Println("Finished writing to image")
	c.SetRGBA(0.3, 0.7, 0.8, 1)
	fmt.Println("Started calculating function")
	for i := 0.0; i < Width; i++ {
		x := (end-start)*i/Width + start
		y := Width * (fn(x) - fnStart) / (fnEnd - fnStart)
		fmt.Println(y, x)
		c.DrawPoint(y, i, 5)
	}
	fmt.Println("Finished calculating, started drawing")
	c.Stroke()
	fmt.Println("finished drawing, started writing")
	c.SavePNG(saveAt)
}
