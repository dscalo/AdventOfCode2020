package fs

import (
	"bufio"
	"os"
)

func LineCount(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, err
	}

	lines := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines++
	}

	err = file.Close()

	return lines, err

}
