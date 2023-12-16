package main

import (
	"Awesome/component"
)

func main() {
	supervisers := component.Start(1)
	for i := range supervisers {
		go supervisers[i].PrintInfo()
	}
	print("start")
}
