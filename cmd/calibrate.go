/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/LorenzHohermuth/cli-shapes/internal/generated"
	"github.com/LorenzHohermuth/cli-shapes/internal/nerdfont"
	"github.com/LorenzHohermuth/cli-shapes/pkg/canvas"
	"github.com/LorenzHohermuth/cli-shapes/pkg/image"
	"github.com/spf13/cobra"
)

// calibrateCmd represents the calibrate command
var calibrateCmd = &cobra.Command{
	Use:   "calibrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		icons := nerdfont.ParseFile("./nerd-font-icons")
		f := generated.CreatFile("nerdFont")
		defer generated.CloseFile(f)
		for i, v := range icons {
			fmt.Printf("%g %% %s\n", (float64(i) + 1) / float64(len(icons)) * 100, v)
			path := canvas.CharacterAnalyse(v)
			img := image.ImageToPixles(path)
			line := fmt.Sprintf("\t\"%s\":%g,", v, image.GetBrightness(img))
			generated.AddToFile(line, f)
		}
	},
}

func init() {
	rootCmd.AddCommand(calibrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calibrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calibrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
