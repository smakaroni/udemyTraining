package main

import (
	"fmt"

	"github.com/smakaroni/udemyTraining/04_scope/01_package-scope/02_visability/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
	//vis.PrintVar is exported from the vis package
}
