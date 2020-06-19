package main

import (
	"fmt"
	"log"
)

type QR struct {
	bitmap     [][]int
	version    int
	size       int
	writeIndex int
}

const (
	EMPTY = iota
	BLACK
	WHITE
	FUNCTION_BLACK
	FUNCTION_WHITE
	RESERVE
)

func (qr *QR) Write(b []byte) (int, error) {
	marker := func(t bool) int {
		if t {
			return WHITE
		}
		return BLACK
	}
	byteToMark := func(b byte) []int {
		bb := make([]int, 8)

		bb[0] = marker(b&1 == 1)
		bb[1] = marker(b&2 == 2)
		bb[2] = marker(b&4 == 4)
		bb[3] = marker(b&8 == 8)
		bb[4] = marker(b&16 == 16)
		bb[5] = marker(b&32 == 32)
		bb[6] = marker(b&64 == 64)
		bb[7] = marker(b&128 == 128)

		return bb
	}

	bb := make([]int, 0)
	for _, v := range b {
		bb = append(bb, byteToMark(v)...)
	}

	p := struct{ X, Y int }{qr.size - 1, qr.size - 1}
	dirList := []struct{ X, Y int }{{-1, 0}, {1, -1}, {1, 1}, {-1, 0}}
	nextDir := 0
	up := true
	count := 0

	for _, m := range bb {
		for f := true; f; {
			if p.X == -1 {

				return -1, fmt.Errorf("Binary to large for this QR Version  b > %d", count>>3)
			}
			mark := qr.bitmap[p.Y][p.X]
			if mark == EMPTY {
				qr.bitmap[p.Y][p.X] = m
				f = false
				count++
			}
			dir := dirList[nextDir]
			p.X += dir.X
			p.Y += dir.Y

			if nextDir == 0 && up {
				nextDir = 1
			} else if nextDir == 0 {
				nextDir = 2
			} else {
				nextDir = 0
			}

			if p.Y == 0 && up && nextDir != 0 {
				up = false
				nextDir = 3

			}
			if p.Y == qr.size-1 && !up && nextDir != 0 {
				up = true
				nextDir = 3

			}
			if p.X == 6 {

				p.X--
			}
		}
	}
	return len(b), nil
}

func (qr *QR) String() string {
	s := ""
	for _, l := range qr.bitmap {
		for _, v := range l {
			switch v {
			case EMPTY:
				s += "."
			case BLACK, FUNCTION_BLACK:
				s += " "
			case WHITE, FUNCTION_WHITE:
				s += "#"
			case RESERVE:
				s += "R"
			default:
				s += "-"
			}
		}
		s += "\n"
	}
	return s
}

func main() {

	qr, err := NewQR([]byte("HELLO WORLD"), M)
	if err != nil {
		log.Fatalf("Making a qr code failed: %v", err)
	}
	// _, err = qr.Write([]byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"))
	// _, err = qr.Write([]byte{0, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 0xff})
	// b := make([]byte, 0)

	// for i := 0; i < 26; i++ {
	// 	b = append(b, byte(i))
	// }
	// _, err = qr.Write(b)
	// if err != nil {
	// 	log.Fatalf("Writing to qr code : %v", err)
	// }
	fmt.Println(qr)

}
