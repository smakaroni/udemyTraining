package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	p1 := Person{"Jokke", "Persson", 20, 200}
	json.NewEncoder(os.Stdout).Encode(p1)
	//bs := json.NewEncoder(os.Stdout)
	//bs.Encode(p1)
}
