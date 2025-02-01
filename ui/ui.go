package ui

import (
	"fmt"
	"image/color"

	"github.com/Maduki-tech/mandelbrot-go/mandelbrot"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Screen dimensions and maximum iterations for Mandelbrot.
const (
	screenWidth  = 800
	screenHeight = 600
	maxIter      = 100
)

type Game struct {
	centerX, centerY float64
	scale            float64
}

func NewGame() *Game {
	return &Game{
		centerX: -0.5,
		centerY: 0.0,
		scale:   3.5,
	}
}

func (g *Game) Update() error {
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		mouseX, mouseY := ebiten.CursorPosition()

		aspectRatio := float64(screenWidth) / float64(screenHeight)
		halfWidth := g.scale / 2.0
		halfHeight := halfWidth / aspectRatio
		realMin := g.centerX - halfWidth
		imagMin := g.centerY - halfHeight

		cx := realMin + (float64(mouseX)/float64(screenWidth))*g.scale
		cy := imagMin + (float64(mouseY)/float64(screenHeight))*(2*halfHeight)

		g.centerX = cx
		g.centerY = cy

		zoomFactor := 0.9
		if wheelY > 0 {
			g.scale *= zoomFactor
		} else {
			g.scale /= zoomFactor
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.centerX -= g.scale * 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.centerX += g.scale * 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.centerY -= g.scale * 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.centerY += g.scale * 0.02
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(screenWidth, screenHeight)
	aspectRatio := float64(screenWidth) / float64(screenHeight)
	halfWidth := g.scale / 2.0
	halfHeight := halfWidth / aspectRatio
	realMin := g.centerX - halfWidth
	imagMin := g.centerY - halfHeight

	for py := 0; py < screenHeight; py++ {
		imag := imagMin + (float64(py)/float64(screenHeight))*(2*halfHeight)
		for px := 0; px < screenWidth; px++ {
			real := realMin + (float64(px)/float64(screenWidth))*g.scale
			c := complex(real, imag)
			iter := mandelbrot.Mandelbrot(c, maxIter)

			var col color.RGBA
			if iter == maxIter {
				col = color.RGBA{0, 0, 0, 255}
			} else {
				gray := uint8(255 - iter*255/maxIter)
				col = color.RGBA{gray, gray, gray, 255}
			}
			img.Set(px, py, col)
		}
	}

	screen.DrawImage(img, nil)

	msg := fmt.Sprintf("Center: (%.5f, %.5f)  Scale: %.5f\nZoom with the mouse wheel, pan with arrow keys.", g.centerX, g.centerY, g.scale)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
