package main

import (
	"fmt"

	"github.com/poxate/gotasia"
)

func main() {
	p := gotasia.NewProject(1920, 1080)
	fmt.Println("project width:", p.Width)
	fmt.Println("project height:", p.Height)
}
