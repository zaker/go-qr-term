package main

import (
	"errors"
	"fmt"
	"strconv"
)

type DataMode int

const (
	Numeric DataMode = iota
	Alphanumeric
	Byte
	Kanji
	ECI
)

type ModeIndicator struct {
	Mode      DataMode
	Name      string
	Indicator string
}

var ModeIndicators []ModeIndicator = []ModeIndicator{
	{Numeric, "Numeric", "0001"},
	{Alphanumeric, "Alphanumeric", "0010"},
	{Byte, "Byte", "0100"},
	{Kanji, "Kanji", "1000"},
	{ECI, "ECI", "0111"},
}

func getModeIndicator(dm DataMode) (*ModeIndicator, error) {
	for _, v := range ModeIndicators {

		if dm == v.Mode {
			return &v, nil
		}
	}

	return nil, errors.New("No such data mode")

}

func getMaxBitlength(version int, dm DataMode) int {
	if version >= 27 {
		switch dm {
		case Numeric:
			return 14
		case Alphanumeric:
			return 13
		case Byte, ECI:
			return 16
		case Kanji:
			return 12
		}
	} else if version >= 10 {
		switch dm {
		case Numeric:
			return 12
		case Alphanumeric:
			return 11
		case Byte, ECI:
			return 16
		case Kanji:
			return 10
		}
	} else if version >= 1 {
		switch dm {
		case Numeric:
			return 10
		case Alphanumeric:
			return 9
		case Byte, ECI:
			return 8
		case Kanji:
			return 8

		}
	}

	return 0

}
func (qr *QR) getIndicators(data []byte) ([]int, error) {

	mb := getMaxBitlength(qr.meta.Version, qr.mode.Mode)

	format := "%0" + strconv.Itoa(mb) + "b"
	cci := fmt.Sprintf(format, len(data))
	t := qr.mode.Indicator + cci
	bb := make([]int, 0)
	for _, v := range t {
		if v == '0' {
			bb = append(bb, FUNCTION_BLACK)
		}
		if v == '1' {
			bb = append(bb, FUNCTION_WHITE)
		}
	}
	return bb, nil
}
