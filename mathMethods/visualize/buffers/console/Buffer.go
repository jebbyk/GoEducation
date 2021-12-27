package console

import (
	"fmt"
	"mathMethods/geometry/types"
)

type Buffer struct {
	h        int
	w        int
	canvas   [][]rune
	bufScale *types.Vector2
}

// set height and width of a buffer
func (b *Buffer) Init(h int, w int) *Buffer {
	b.h = h
	b.w = w

	b.bufScale = new(types.Vector2).Set(float64(w), float64(h))

	for i := 0; i < h; i++ {
		row := []rune{}

		for j := 0; j < w; j++ {
			row = append(row, ' ')
		}

		b.canvas = append(b.canvas, row)
	}

	return b
}

// displays content of a buffer
func (b *Buffer) Print() {
	for i := 0; i < len(b.canvas); i++ {
		fmt.Println(string(b.canvas[i]))
	}
}

// set pixel of a buffer using coordinates int ranges of buffer rows and columns amount
func (b *Buffer) SetPixelDirect(x int, y int, symbol rune) {
	if x > 0 && x < b.w && y > 0 && y < b.h {
		b.canvas[y][x] = symbol
	}

}

// set pixel of a buffer using coordinates in range between 0.....1
func (b *Buffer) SetPixelRelative(pos *types.Vector2, symbol rune) {
	pos = pos.Mul(b.bufScale)

	x := int(pos.GetX())
	y := int(pos.GetY())

	b.SetPixelDirect(x, y, symbol)
}

func (b *Buffer) DrawRect(topLeft *types.Vector2, bottomRight *types.Vector2, symbol rune) {
	topLeftScaled := topLeft.Mul(b.bufScale)
	bottomRightScaled := bottomRight.Mul(b.bufScale)

	tLx := int(topLeftScaled.GetX())
	tLy := int(topLeftScaled.GetY())

	bRx := int(bottomRightScaled.GetX())
	bRy := int(bottomRightScaled.GetY())

	for i := tLy; i < bRy; i++ {
		for j := tLx; j < bRx; j++ {
			b.SetPixelDirect(j, i, symbol)
		}
	}

}

// fill entire buffer with space symbols
func (b *Buffer) Clear() *Buffer {
	for i := 0; i < b.h; i++ {
		for j := 0; j < b.w; j++ {
			b.SetPixelDirect(j, i, ' ')
		}
	}

	return b
}
