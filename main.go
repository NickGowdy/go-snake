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
	squareFile   = "square.png"
)

var (
	ebitenImg      *ebiten.Image
	drawImgOptions *ebiten.DrawImageOptions
)

type Game struct {
}

func init() {
	createSnakePng()

	squareDir := getSquareDir()

	f, err := os.Open(squareDir)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImg := ebiten.NewImageFromImage(img)

	position := origEbitenImg.Bounds().Size()
	ebitenImg = ebiten.NewImage(position.X, position.Y)

	drawImgOptions = &ebiten.DrawImageOptions{}
	drawImgOptions.ColorScale.ScaleAlpha(0.5)
	ebitenImg.DrawImage(origEbitenImg, drawImgOptions)
}

func createSnakePng() (*os.File, error) {
	squareDir := getSquareDir()

	myimage := image.NewRGBA(image.Rect(0, 0, 10, 10))
	black := color.RGBA{255, 255, 255, 255}

	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	myfile, err := os.Create(squareDir)
	if err != nil {
		panic(err)
	}
	png.Encode(myfile, myimage)
	return myfile, err
}

func getSquareDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	squareDir := fmt.Sprintf("%s/%s", path, squareFile)
	return squareDir
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(ebitenImg, drawImgOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWitdth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
