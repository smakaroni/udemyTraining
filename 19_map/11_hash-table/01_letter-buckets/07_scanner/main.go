package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "Sjung för gamla djurgårn nu sjung i ur och skur, visa folk att våra ränder dom går aldrig ur, vem som helst kan sticka upp men vi är nummer ett. Så länge gamla djurgårn finns känns livet ändå lätt"
	scanner := bufio.NewScanner(strings.NewReader(input))
	//Set the split function
	scanner.Split(bufio.ScanWords)
	//Count the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
