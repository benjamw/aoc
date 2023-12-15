package structs

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

const (
	On      byte = '#'
	Off     byte = '.'
	Unknown byte = '?'
)

type Nonogram struct {
	Line        []byte
	Pattern     []int
	NumPossible int
}

// Filter out any impossible locations for bits
func (n *Nonogram) Filter() {
	// look at the start of the line and see if any ? can be inferred
	for i, v := range n.Line {
		_ = i
		_ = v
	}
}

// CountPossible counts the number of possible solutions to the known
// line and pattern
func (n *Nonogram) CountPossible() int {

	return n.NumPossible
}

// Unfold makes c copies of the Line (with '?' between each copy)
// and c copies of the Pattern
func (n *Nonogram) Unfold(c int) {
	ol := n.Line
	op := n.Pattern
	for i := 1; i < c; i++ {
		n.Line = append(n.Line, Unknown)
		n.Line = append(n.Line, ol...)
		n.Pattern = append(n.Pattern, op...)
	}
}

// MakeAllPossible is the brute force method where
// all possible lines that can be created from the known line
// are created and then checked against the pattern to see what matches
func (n *Nonogram) MakeAllPossible() [][]byte {
	p := make([][]byte, 0)

	// find the indexes of all the Unknown ('?') values
	re := regexp.MustCompile(fmt.Sprintf("\\%s", string(Unknown)))
	idxs := re.FindAllIndex(n.Line, -1)
	// create a binary number with the same number of bits as num of '?'
	b, _ := strconv.ParseInt(fmt.Sprintf("0b%s", strings.Repeat("1", len(idxs))), 0, 64)
	for i := b; i >= 0; i-- {
		l := n.Line
		p = append(p, makeBinaryPattern(l, idxs, i))
	}

	f := n.FilterPatterns(p)
	n.NumPossible = len(f)

	return f
}

func (n *Nonogram) FilterPatterns(p [][]byte) [][]byte {
	ret := make([][]byte, 0)
	for _, b := range p {
		if n.PatternMatches(b) {
			ret = append(ret, b)
		}
	}

	return ret
}

func (n *Nonogram) PatternMatches(p []byte) bool {
	// count the number of contiguous '#' values
	cs := make([]int, 0)
	c := 0
	for _, b := range p {
		switch b {
		case Off:
			if c > 0 {
				cs = append(cs, c)
				c = 0
			}
		case On:
			c++
		}
	}

	if c > 0 {
		cs = append(cs, c)
		c = 0
	}

	return reflect.DeepEqual(n.Pattern, cs)

}

func makeBinaryPattern(l []byte, idxs [][]int, v int64) []byte {
	cl := append(make([]byte, 0), l...) // clone l
	// convert the v value into a []byte of chars
	b := []byte(strconv.FormatInt(v, 2))
	// pad to the left with 0s if not enough
	for len(b) < len(idxs) {
		// prepend
		b = append([]byte{'0'}, b...)
	}

	for i, bv := range b {
		switch bv {
		case '0':
			cl = replaceChar(cl, idxs[i][0], Off)
		case '1':
			cl = replaceChar(cl, idxs[i][0], On)
		}
	}

	return cl
}

func replaceChar(l []byte, idx int, ch byte) []byte {
	return slices.Replace(l, idx, idx+1, ch)
}

func printPatterns(p [][]byte) {
	for _, bytes := range p {
		printPattern(bytes)
	}
}

func printPattern(p []byte) {
	spew.Dump(string(p))
}
