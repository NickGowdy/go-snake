package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	North = iota + 1
	East
	South
	West
	screenWidth  = 320
	screenHeight = 240
	snakeWidth   = 10
	snakeHeight  = 10
	squareFile   = "square.png"
)

var (
	newEbitenImg   *ebiten.Image
	drawImgOptions *ebiten.DrawImageOptions
	direction      int
	currentX       int
	currentY       int
)

type Game struct {
	op     ebiten.DrawImageOptions
	inited bool
}

func createSnakePng() (*os.File, error) {
	squareDir := getSquareDir()

	myimage := image.NewRGBA(image.Rect(0, 0, snakeWidth, snakeHeight))
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

// init gets png file used for the game and decodes it to be used by
// var newEbitenImg and drawImgOptions. These vars are modified throughout
// lifetime of the game
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
	newEbitenImg = ebiten.NewImage(position.X, position.Y)

	drawImgOptions = &ebiten.DrawImageOptions{}
	drawImgOptions.ColorScale.ScaleAlpha(0.5)
	newEbitenImg.DrawImage(origEbitenImg, drawImgOptions)
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	min := North
	max := West
	direction = rand.Intn(max-min) + min
}

func (g *Game) Update() error {

	switch direction {
	case 1:
		currentY++
	case 2:
		currentX++
	case 3:
		currentY--
	case 4:
		currentX++
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	w, h := newEbitenImg.Bounds().Dx(), newEbitenImg.Bounds().Dy()
	g.op.GeoM.Reset()

	if !g.inited {

		xMiddle := screenWidth / 2
		yMiddle := screenHeight / 2
		currentX = xMiddle
		currentY = yMiddle

		g.op.GeoM.Translate(float64(xMiddle), float64(yMiddle))
	} else {
		g.op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		g.op.GeoM.Translate(float64(currentX), float64(currentY))
	}
	
	screen.DrawImage(newEbitenImg, &g.op)

	if !g.inited {
		g.init()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("GO SNAKE")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
