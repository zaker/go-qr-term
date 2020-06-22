package main

import "fmt"

func byteEncode(b []byte) []int {
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
	return bb
}

func (qr *QR) addEncodedData(encodedData []int) error {

	p := struct{ X, Y int }{qr.size - 1, qr.size - 1}
	dirList := []struct{ X, Y int }{{-1, 0}, {1, -1}, {1, 1}, {-1, 0}}
	nextDir := 0
	up := true
	count := 0

	for _, m := range encodedData {
		for f := true; f; {
			if p.X == -1 {

				return fmt.Errorf("Binary to large for this QR Version  b > %d", count>>3)
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
	return nil
}
