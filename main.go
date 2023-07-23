package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1024
	screenHeight = 800
	squareDir    = "square.png"
)

var (
	ebitenImage *ebiten.Image
	op          *ebiten.DrawImageOptions
)

type Game struct {
}

func init() {
	drawRectangle()

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	filePath := fmt.Sprintf("%s/%s", path, squareDir)

	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Println(err)
	}

	origEbitenImage := ebiten.NewImageFromImage(img)

	s := origEbitenImage.Bounds().Size()
	ebitenImage = ebiten.NewImage(s.X, s.Y)

	op = &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(0.5)
	ebitenImage.DrawImage(origEbitenImage, op)
}

func drawRectangle() (*os.File, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	new_png_file := fmt.Sprintf("%s/%s", path, squareDir)

	myimage := image.NewRGBA(image.Rect(0, 0, 10, 10))
	black := color.RGBA{255, 255, 255, 255}

	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	myfile, err := os.Create(new_png_file)
	if err != nil {
		panic(err)
	}
	png.Encode(myfile, myimage) // output file /tmp/two_rectangles.png
	return myfile, err
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(ebitenImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWitdth, screenHeight int) {
	return 500, 240
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
