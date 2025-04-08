package gungo

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Screen *ebiten.Image
}

func (G *Game) Draw(screen *ebiten.Image) {
	G.Screen = screen
}
func (G *Game) Update() error {
	return nil
}
func (G *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

var (
	GG *Game
)

func Setscreen(width int, height int, name string) {
	GG = &Game{}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(name)
	ebiten.RunGame(GG)

}

func DrawRectBorder(x float32, y float32, width float32, height float32, color color.Color, border_width float32, antialias bool) {
	vector.StrokeRect(GG.Screen, x, y, width, height, border_width, color, antialias)
}

func DrawRectFilled(x float32, y float32, width float32, height float32, color color.Color, antialias bool) {
	vector.DrawFilledRect(GG.Screen, x, y, width, height, color, antialias)
}
