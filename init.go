package main

import "errors"

type ErrorCorrectionLevel int

const (
	L ErrorCorrectionLevel = iota
	M
	Q
	H
)

type QR struct {
	bitmap [][]int
	meta   *VersionMetaData
	size   int
	mode   *ModeIndicator
}

func NewQR(data []byte, el ErrorCorrectionLevel, dm DataMode) (*QR, error) {
	qrMeta, err := SmallestQRVersion(len(data), el)

	if err != nil {
		return nil, err
	}
	size := (((qrMeta.Version - 1) * 4) + 21)

	bitmap := make([][]int, size)
	for i := range bitmap {
		bitmap[i] = make([]int, size)
	}
	mode, err := getModeIndicator(dm)
	if err != nil {
		return nil, err
	}
	qr := &QR{bitmap, qrMeta, size, mode}

	qr.addFinders()
	qr.addSeparators()
	qr.addAlignments()
	qr.addTiming()
	qr.addReserved()
	qr.addModule()
	indicator, err := qr.getIndicators(data)
	if err != nil {
		return nil, err
	}
	encodedData, err := qr.encode(data)
	if err != nil {
		return nil, err
	}
	maxBits := qr.meta.MaxCodewords * 8

	encodedData = append(indicator, encodedData...)
	for i := 0; len(encodedData) < maxBits && i < 4; i++ {
		encodedData = append(encodedData, BLACK)
	}
	for len(encodedData)%8 != 0 {
		encodedData = append(encodedData, BLACK)
	}

	pad236 := []int{WHITE, WHITE, WHITE, BLACK, WHITE, WHITE, BLACK, BLACK}
	pad7 := []int{BLACK, BLACK, BLACK, WHITE, BLACK, BLACK, BLACK, WHITE}

	for i := false; len(encodedData) < maxBits; i = !i {
		if i {
			encodedData = append(encodedData, pad236...)
		} else {
			encodedData = append(encodedData, pad7...)
		}

	}
	err = qr.addEncodedData(encodedData)
	if err != nil {
		return nil, err
	}
	return qr, nil
}

func (qr *QR) encode(data []byte) ([]int, error) {

	if qr.mode.Mode == Byte {
		return byteEncode(data), nil
	}
	return []int{}, errors.New("No encoding mode set")
}

func (qr *QR) addFinders() {

	finderPattern := [][]int{
		{FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK},
	}
	// Add top-left finder

	for y, l := range finderPattern {
		copy(qr.bitmap[y], l)
	}

	offsett := (((qr.meta.Version - 1) * 4) + 21) - 7
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

	offsett2 := (((qr.meta.Version - 1) * 4) + 21) - (offsett + 1)

	y := 0
	for i := 0; i < 7; i++ {
		qr.bitmap[y][offsett] = FUNCTION_WHITE
		qr.bitmap[y][offsett2] = FUNCTION_WHITE
		y++
	}

	for i := 0; i <= offsett; i++ {
		qr.bitmap[y][i] = FUNCTION_WHITE
	}

	for i := offsett2; i < qr.size; i++ {
		qr.bitmap[y][i] = FUNCTION_WHITE

	}

	for ; y < offsett2; y++ {
	}

	for i := 0; i <= offsett; i++ {
		qr.bitmap[y][i] = FUNCTION_WHITE
	}
	y++
	for i := 0; i < 7; i++ {
		qr.bitmap[y][offsett] = FUNCTION_WHITE
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
	if qr.meta.Version == 1 {
		return
	}

	alignment := alignmentMap[qr.meta.Version]
	alignerPattern := [][]int{
		{FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_WHITE, FUNCTION_BLACK},
		{FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK, FUNCTION_BLACK},
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

	offsett2 := (((qr.meta.Version - 1) * 4) + 21) - (offsett + 1)

	y := 6
	marker := func(i int) int {
		if i%2 == 1 {
			return FUNCTION_WHITE
		}
		return FUNCTION_BLACK
	}
	for i := 8; i < offsett2; i++ {
		qr.bitmap[y][i] = marker(i)

	}
	y++
	for ; y < offsett2; y++ {
		qr.bitmap[y][6] = marker(y)

	}

}

func (qr *QR) addModule() {

	y := (4 * qr.meta.Version) + 9
	x := 8

	qr.bitmap[y][x] = FUNCTION_BLACK
}

func (qr *QR) addReserved() {

	offsett := (((qr.meta.Version - 1) * 4) + 21) - 8

	y := 0

	for ; y < 8; y++ {
		if qr.bitmap[y][8] == EMPTY {
			qr.bitmap[y][8] = RESERVE

			if y < 6 && qr.meta.Version != 1 {
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

	for y = offsett - 3; y < offsett && qr.meta.Version != 1; y++ {
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
