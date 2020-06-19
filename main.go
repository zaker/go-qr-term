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
	RESERVE
)

func NewQR(version int) (*QR, error) {
	if version < 1 || version > 40 {
		return nil, fmt.Errorf("Can only produce qr versions 1 - 40, not %d", version)
	}
	size := (((version - 1) * 4) + 21)

	bitmap := make([][]int, size)
	for i := range bitmap {
		bitmap[i] = make([]int, size)
	}

	qr := &QR{bitmap, version, size, 0}

	qr.addFinders()
	qr.addSeparators()
	qr.addAlignments()
	qr.addTiming()
	qr.addReserved()
	qr.addBlackModule()
	return qr, nil
}

func (qr *QR) addFinders() {

	finderPattern := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 2, 2, 2, 2, 2, 1},
		{1, 2, 1, 1, 1, 2, 1},
		{1, 2, 1, 1, 1, 2, 1},
		{1, 2, 1, 1, 1, 2, 1},
		{1, 2, 2, 2, 2, 2, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	// Add top-left finder

	for y, l := range finderPattern {
		copy(qr.bitmap[y], l)
	}

	offsett := (((qr.version - 1) * 4) + 21) - 7
	// Add bottom-left finder
	for y, l := range finderPattern {
		copy(qr.bitmap[y+offsett], l)
	}

	// Add top-right finder
	for y, l := range finderPattern {
		copy(qr.bitmap[y][offsett:], l)
	}

}

func (qr *QR) addSeparators() {

	offsett := 7

	offsett2 := (((qr.version - 1) * 4) + 21) - (offsett + 1)

	y := 0
	for i := 0; i < 7; i++ {
		qr.bitmap[y][offsett] = WHITE
		qr.bitmap[y][offsett2] = WHITE
		y++
	}

	for i := 0; i <= offsett; i++ {
		qr.bitmap[y][i] = WHITE
	}

	for i := offsett2; i < qr.size; i++ {
		qr.bitmap[y][i] = WHITE

	}

	for ; y < offsett2; y++ {
	}

	for i := 0; i <= offsett; i++ {
		qr.bitmap[y][i] = WHITE
	}
	y++
	for i := 0; i < 7; i++ {
		qr.bitmap[y][offsett] = WHITE
		y++
	}
}

var alignmentMap map[int][]int = map[int][]int{
	1:  {},
	2:  {6, 18},
	3:  {6, 22},
	4:  {6, 26},
	5:  {6, 30},
	6:  {6, 34},
	7:  {6, 22, 38},
	8:  {6, 24, 42},
	9:  {6, 26, 46},
	10: {6, 28, 50},
	11: {6, 30, 54},
	12: {6, 32, 58},
	13: {6, 34, 62},
	14: {6, 26, 46, 66},
	15: {6, 26, 48, 70},
	16: {6, 26, 50, 74},
	17: {6, 30, 54, 78},
	18: {6, 30, 56, 82},
	19: {6, 30, 58, 86},
	20: {6, 34, 62, 90},
	21: {6, 28, 50, 72, 94},
	22: {6, 26, 50, 74, 98},
	23: {6, 30, 54, 78, 102},
	24: {6, 28, 54, 80, 106},
	25: {6, 32, 58, 84, 110},
	26: {6, 30, 58, 86, 114},
	27: {6, 34, 62, 90, 118},
	28: {6, 26, 50, 74, 98, 122},
	29: {6, 30, 54, 78, 102, 126},
	30: {6, 26, 52, 78, 104, 130},
	31: {6, 30, 56, 82, 108, 134},
	32: {6, 34, 60, 86, 112, 138},
	33: {6, 30, 58, 86, 114, 142},
	34: {6, 34, 62, 90, 118, 146},
	35: {6, 30, 54, 78, 102, 126, 150},
	36: {6, 24, 50, 76, 102, 128, 154},
	37: {6, 28, 54, 80, 106, 132, 158},
	38: {6, 32, 58, 84, 110, 136, 162},
	39: {6, 26, 54, 82, 110, 138, 166},
	40: {6, 30, 58, 86, 114, 142, 170},
}

func (qr *QR) addAlignments() {
	if qr.version == 1 {
		return
	}

	alignment := alignmentMap[qr.version]
	alignerPattern := [][]int{
		{1, 1, 1, 1, 1},
		{1, 2, 2, 2, 1},
		{1, 2, 1, 2, 1},
		{1, 2, 2, 2, 1},
		{1, 1, 1, 1, 1},
	}
	for _, yOff := range alignment {
		for _, xOff := range alignment {
			if qr.bitmap[yOff][xOff] != EMPTY {
				continue
			}
			for y, l := range alignerPattern {
				copy(qr.bitmap[yOff+y-2][xOff-2:], l)
			}
		}
	}

}

func (qr *QR) addTiming() {

	offsett := 7

	offsett2 := (((qr.version - 1) * 4) + 21) - (offsett + 1)

	y := 6
	marker := func(i int) int {
		if i%2 == 1 {
			return WHITE
		}
		return BLACK
	}
	for i := 8; i < offsett2; i++ {
		qr.bitmap[y][i] = marker(i)

	}
	y++
	for ; y < offsett2; y++ {
		qr.bitmap[y][6] = marker(y)

	}

}

func (qr *QR) addBlackModule() {

	y := (4 * qr.version) + 9
	x := 8

	qr.bitmap[y][x] = BLACK
}

func (qr *QR) addReserved() {
	offsett := (((qr.version - 1) * 4) + 21) - 8

	y := 0

	for ; y < 8; y++ {
		if qr.bitmap[y][8] == EMPTY {
			qr.bitmap[y][8] = RESERVE

			if y < 6 {
				qr.bitmap[y][offsett-3] = RESERVE
				qr.bitmap[y][offsett-2] = RESERVE
				qr.bitmap[y][offsett-1] = RESERVE
			}
		}

	}
	for i := 0; i < 9; i++ {
		if qr.bitmap[y][i] == EMPTY {
			qr.bitmap[y][i] = RESERVE
		}
	}

	for i := offsett; i < len(qr.bitmap[y]); i++ {
		if qr.bitmap[y][i] == EMPTY {
			qr.bitmap[y][i] = RESERVE
		}
	}

	for y = offsett - 3; y < offsett; y++ {
		for x := 0; x < 6; x++ {
			qr.bitmap[y][x] = RESERVE
		}
	}
	for y := offsett; y < len(qr.bitmap); y++ {
		if qr.bitmap[y][8] == EMPTY {
			qr.bitmap[y][8] = RESERVE
		}
	}

}

func (qr *QR) Write(b []byte) (int, error) {

	byteToBool := func(b byte) []bool {
		bb := make([]bool, 8)

		bb[0] = b&1 == 1
		bb[1] = b&2 == 2
		bb[2] = b&4 == 4
		bb[3] = b&8 == 8
		bb[4] = b&16 == 16
		bb[5] = b&32 == 32
		bb[6] = b&64 == 64
		bb[7] = b&128 == 128

		return bb
	}

	bb := make([]bool, 0)
	for _, v := range b {
		bb = append(bb, byteToBool(v)...)
	}
	return 0, nil
}

func (qr *QR) String() string {
	s := ""
	for _, l := range qr.bitmap {
		for _, v := range l {
			switch v {
			case EMPTY:
				s += "."
			case BLACK:
				s += " "
			case WHITE:
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

	qr, err := NewQR(3)
	if err != nil {
		log.Fatalf("Making a qr code failed: %v", err)
	}
	_, err = qr.Write([]byte("HELLO"))
	if err != nil {
		log.Fatalf("Writing to qr code : %v", err)
	}
	fmt.Println(qr)

}
