/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/LorenzHohermuth/cli-shapes/cmd"
	"github.com/LorenzHohermuth/cli-shapes/internal/color"
	"github.com/LorenzHohermuth/cli-shapes/internal/renderer"
	"github.com/LorenzHohermuth/cli-shapes/pkg/image"
)

func main() {
	//pix := image.ImageToPixles("/home/lorenz/Downloads/mushroom-mini.png")
	//pix := image.ImageToPixles("/home/lorenz/Downloads/rgb.png")
	pix := image.ImageToPixles("/home/lorenz/Downloads/yoshi.png")
	s := ""
	for _, y := range pix {
		for _, x := range y {
			color, non := color.GetColor([3]int{x.R, x.G, x.B})
			char := renderer.GetChar(x.Brightness())
			s += color + char + non
		}
		s += "\n"
	}
	fmt.Println(s)
	color.PrintColorPallet()
	cmd.Execute()
}
