package console

import "fmt"

type Buffer struct {
	h      int
	w      int
	canvas [][]rune
}

// set height and width of a buffer
func (b *Buffer) Init(h int, w int) *Buffer {
	b.h = h
	b.w = w

	return b
}

// displays content of a buffer
func (b *Buffer) Print() {
	for i := 0; i < len(b.canvas); i++ {
		fmt.Println(b.canvas[i])
	}
}

// set pixel of a buffer
func (b *Buffer) SetPixel(x int, y int, symbol rune) {
	b.canvas[x][y] = symbol
}