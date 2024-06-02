package readmaze

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// This package will provide functions used for taking the
// maze file and reading it into a data object that we can
// use. It offers no other functionality than that.

func Ingest(f string) []string {
	fp, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	var retval []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		retval = append(retval, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return retval
}

func PrintMaze(maze []string) {
	// just prints the maze after we ingest it
	for _, v := range maze {
		fmt.Println(v)
	}
}
