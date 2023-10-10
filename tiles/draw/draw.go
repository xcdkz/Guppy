package draw

import (
	"github.com/fatih/color"
	"projects/twpsx/guppy/tiles"
	"projects/twpsx/guppy/tiles/cursor"
)

func DrawBorder(leaf *tiles.Tile) {
	var c color.Color
	if leaf.IsFocused {
		c = *color.New(color.FgHiRed)
	} else {
		c = *color.New(color.FgBlack)
	}
	px, py := leaf.PosX, leaf.PosY
	sx, sy := leaf.SizeX, leaf.SizeY
	cursor.MoveTo(px, py+1)
	for i := 0; i < sx; i++ {
		c.Print("▄")
	}
	cursor.MoveTo(px, py+sy+1)
	for i := 0; i < sx; i++ {
		c.Print("▀")
	}
	cursor.MoveTo(px, py+2)
	for i := 1; i < sy+1; i++ {
		c.Print("█")
		cursor.MoveTo(px, py+1+i)
	}
	c.Print("▀")
	cursor.MoveTo(px+sx, py+2)
	for i := 1; i < sy+1; i++ {
		c.Print("█")
		cursor.MoveTo(px+sx, py+1+i)
	}
	c.Print("▀")
}