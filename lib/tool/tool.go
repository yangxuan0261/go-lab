package tool

import (
	"bufio"
	"io"
	"os"
)

func Readline(path string) ([]string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	var arr []string
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if len(a) > 0 {
			arr = append(arr, string(a))
		}
	}
	return arr, nil
}
