/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/LorenzHohermuth/cli-shapes/cmd"
	"github.com/LorenzHohermuth/cli-shapes/internal/color"
	"github.com/LorenzHohermuth/cli-shapes/internal/renderer"
	"github.com/LorenzHohermuth/cli-shapes/pkg/image"
)

func main() {
	pix := image.ImageToPixles("/home/lorenz/Downloads/crab.png")
	s := ""
	for _, y := range pix {
		for _, x := range y {
			s += renderer.GetChar(x.Brightness())
		}
		s += "\n"
	}
	color.PrintColorPallet()
	cmd.Execute()
}
