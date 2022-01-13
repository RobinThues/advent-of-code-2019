package twentyone

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BingoField struct {
	number   int
	isMarked bool
}

type BingoBoard struct {
	fields   [][]BingoField
	isSolved bool
}

func (b *BingoBoard) String() string {
	var s string
	for k, v := range b.fields {
		for kk, vv := range v {
			s += fmt.Sprintf("F[%d,%d]: %d %t ", k, kk, vv.number, vv.isMarked)
		}
		s += "\n"
	}
	return s
}

func (b *BingoBoard) Points() int {
	p := 0
	for k, _ := range b.fields {
		for kk, _ := range b.fields[k] {
			if b.fields[k][kk].isMarked == false {
				p += b.fields[k][kk].number
			}
		}
	}
	return p
}

func (b *BingoBoard) markNumber(n int) bool {
	for row, lines := range b.fields {
		for col, _ := range lines {
			if b.fields[row][col].number == n {
				b.fields[row][col].isMarked = true
				r := b.isRowSolved(row)
				c := b.isColSolved(col)

				if r || c {
					b.isSolved = true
					return b.isSolved
				}
			}
		}
	}
	return false
}

func (b *BingoBoard) isRowSolved(rowNumber int) bool {
	for k, _ := range b.fields[rowNumber] {
		if b.fields[rowNumber][k].isMarked == false {
			return false
		}
	}
	return true
}

func (b *BingoBoard) isColSolved(colNumber int) bool {
	for _, v := range b.fields {
		if v[colNumber].isMarked == false {
			return false
		}
	}
	return true
}

func readBingoLineToInt(line string) []int {
	var intLine []int

	line = strings.TrimSpace(line)

	space := regexp.MustCompile(`\s+`)
	line = space.ReplaceAllString(line, " ")

	l := strings.Split(line, " ")

	for _, v := range l {
		i, _ := strconv.Atoi(v)
		intLine = append(intLine, i)
	}
	return intLine
}

func D4one(lines []string) {
	// collect drawn numbers
	s := strings.Split(lines[0], ",")
	var drawnNumbers []int
	for _, nbr := range s {
		i, _ := strconv.Atoi(nbr)
		drawnNumbers = append(drawnNumbers, i)
	}

	// setup bingo fields with no numbers marked
	var currentlyReadBoard BingoBoard
	var boards []BingoBoard
	boardFillLevel := 0
	for k, l := range lines {
		if len(l) == 0 {
			continue
		}
		if k < 2 {
			continue
		}

		numbersInLine := readBingoLineToInt(l)
		var bingoLine []BingoField
		for _, vv := range numbersInLine {
			bingoLine = append(bingoLine, BingoField{
				vv,
				false,
			})
		}
		currentlyReadBoard.fields = append(currentlyReadBoard.fields, bingoLine)
		boardFillLevel += 1

		if boardFillLevel == 5 {
			boardFillLevel = 0
			boards = append(boards, currentlyReadBoard)
			currentlyReadBoard.fields = [][]BingoField{}
		}
	}

	// P1
	/*
		for _, v := range drawnNumbers {
			for k, _ := range boards {
				if boards[k].markNumber(v) {
					p := boards[k].Points()
					println("Board: ", k, " Points: ", p, " Nbr: ", v, " Solution: ", p*v)
					println(boards[k].String())
					return
				}
			}
		}

	*/

	// P2
	unsolvedBingos := len(boards)
	for _, v := range drawnNumbers {
		for k, _ := range boards {
			if boards[k].isSolved {
				continue
			}

			if boards[k].markNumber(v) {
				if unsolvedBingos == 1 {
					p := boards[k].Points()
					println("Board: ", k, " Points: ", p, " Nbr: ", v, " Solution: ", p*v)
					return
				}
				unsolvedBingos -= 1
			}
		}
	}

}
