/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/LorenzHohermuth/cli-shapes/internal/color"
	"github.com/LorenzHohermuth/cli-shapes/internal/renderer"
	"github.com/LorenzHohermuth/cli-shapes/pkg/image"
	"github.com/spf13/cobra"
)

// drawCmd represents the draw command
var drawCmd = &cobra.Command{
	Use:   "draw",
	Short: "Draws a defined Image on Screen in ASCII",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileFlag, _ := cmd.Flags().GetString("png")
		if fileFlag != "" {
			pix := image.ImageToPixles(fileFlag)
			s := ""
			for _, y := range pix {
				for _, x := range y {
					color, non := color.GetTrueColor([3]int{x.R, x.G, x.B})
					char := renderer.GetChar(x.Brightness())
					s += color + char + non
				}
				s += "\n"
			}
			fmt.Println(s)

		}
	},
}

func init() {
	rootCmd.AddCommand(drawCmd)
	drawCmd.Flags().String("png", "", "png image")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// drawCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// drawCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
