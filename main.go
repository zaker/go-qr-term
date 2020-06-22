package main

import (
	"fmt"
	"log"
)

const (
	EMPTY = iota
	BLACK
	WHITE
	FUNCTION_BLACK
	FUNCTION_WHITE
	RESERVE
)

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

	qr, err := NewQR([]byte("HELLO WORLD"), M, Byte)
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
