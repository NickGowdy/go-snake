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
	screenWidth  = 320
	screenHeight = 240
	maxAngle     = 256
)

var (
	ebitenImage *ebiten.Image
)

type Game struct {
	// touchIDs []ebiten.TouchID
	// sprites  Sprites
	// op     ebiten.DrawImageOptions
	// inited bool
}

func init() {
	// output image will live here
	// x1,y1,  x2,y2 of background rectangle
	//  R, G, B, Alpha
	// backfill entire background surface with color mygreen
	//  geometry of 2nd rectangle which we draw atop above rectangle
	// create a red rectangle atop the green surface
	// ... now lets save output image
	drawRectangle()

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	filePath := fmt.Sprintf("%s/two_rectangles.png", path)

	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	if err != nil {
		log.Println(err)
	}
	fmt.Print(image)

}

func drawRectangle() (*os.File, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	new_png_file := fmt.Sprintf("%s/two_rectangles.png", path)

	myimage := image.NewRGBA(image.Rect(0, 0, 220, 220))
	mygreen := color.RGBA{0, 100, 0, 255}

	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)

	red_rect := image.Rect(60, 80, 120, 160)
	myred := color.RGBA{200, 0, 0, 255}

	draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)

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
	screen.DrawImage(screen, &ebiten.DrawImageOptions{})
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
