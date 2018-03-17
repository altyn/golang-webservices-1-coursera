package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		txt := in.Text()

		if txt == prev {
			continue
		}

		if txt < prev {
			return fmt.Errorf("file not sorted")
		}

		prev = txt
		fmt.Fprintln(output, txt)
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error)
	}
}

// func main() {
// 	in := bufio.NewScanner(os.Stdin)
// 	var prev string
// 	for in.Scan() {
// 		txt := in.Text()

// 		if txt == prev {
// 			continue
// 		}

// 		if txt < prev {
// 			panic("file not sorted")
// 		}

// 		prev = txt
// 		fmt.Println(txt)
// 	}
// }

// func main() {
// 	in := bufio.NewScanner(os.Stdin)
// 	alreadySeen := make(map[string]bool)
// 	for in.Scan() {
// 		txt := in.Text()

// 		if _, found := alreadySeen[txt]; found {
// 			continue
// 		}
// 		alreadySeen[txt] = true
// 		fmt.Println(txt)
// 	}
// }
