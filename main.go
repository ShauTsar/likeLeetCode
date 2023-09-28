package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var k int
	var n int
	cards := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString(' ')
	k, _ = strconv.Atoi(strings.TrimSpace(s))
	s, _ = reader.ReadString(' ')
	n, _ = strconv.Atoi(strings.TrimSpace(s))
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString(' ')
		x, _ := strconv.Atoi(strings.TrimSpace(s))
		cards = append(cards, x)
	}
	writer := bufio.NewWriter(os.Stdout)
	for _, i := range cards {
		writer.WriteString(strconv.Itoa(i))
	}
	writer.WriteString(strconv.Itoa(k))
	writer.Flush()

}
