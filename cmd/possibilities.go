package cmd

import (
	"fmt"

	cut_dimensions "github.com/lubieniebieski/goCutSomething/internal"
	"github.com/spf13/cobra"
)

var width, margin, desiredDiameter int
var outputFormat string
var availableLength int
var desiredSpacing int
var foundCuts []cut_dimensions.CutDimensions
var minLength int
var minPieces int

var possibilities = &cobra.Command{
	Use: "possibilities",

	Args: cobra.ExactArgs(0),

	Run: func(cmd *cobra.Command, args []string) {
		availableLength := width - margin
		desiredSpacing := 100
		foundCuts = cut_dimensions.PossibleCutDimensions(availableLength, desiredDiameter, desiredSpacing)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		availableLength = width - margin
		showInfo()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if outputFormat == "csv" {
			printAsCSV(foundCuts)
			return
		} else if outputFormat == "google-sheets" {
			printToGoogleSheets(foundCuts)
			return
		} else if outputFormat == "text" {
			printOutput(foundCuts)
			return
		}
	},
}
var best_cut = &cobra.Command{
	Use: "best_cut",

	Args: cobra.ExactArgs(0),

	Run: func(cmd *cobra.Command, args []string) {
		possibilities := cut_dimensions.PossibleCutDimensions(availableLength, desiredDiameter, desiredSpacing)
		bestCut := cut_dimensions.FindBestCut(possibilities, minLength, minPieces, availableLength)
		fmt.Printf("Best cut: %dmm, %d pieces\n", bestCut.GetDistanceBetweenCenters(), bestCut.Pieces)
	},
}

func init() {
	possibilities.PersistentFlags().IntVarP(&width, "width", "w", 100, "Width of the material")
	possibilities.PersistentFlags().IntVarP(&margin, "margin", "m", 0, "Margin from the edge of the material")
	possibilities.PersistentFlags().IntVarP(&desiredDiameter, "diameter", "d", 10, "Desired diameter of the holes")
	possibilities.PersistentFlags().IntVarP(&desiredSpacing, "spacing", "s", 100, "Desired spacing between the holes")
	possibilities.PersistentFlags().StringVarP(&outputFormat, "output", "o", "text", "Output format (text, csv, google-sheets)")
	best_cut.Flags().IntVarP(&minLength, "min-length", "l", 15, "Minimum length of the cut")
	best_cut.Flags().IntVarP(&minPieces, "min-pieces", "p", 2, "Minimum number of pieces")
	rootCmd.AddCommand(possibilities)
	possibilities.AddCommand(best_cut)
}

func printOutput(cuts []cut_dimensions.CutDimensions) {
	for _, possibility := range cuts {
		fmt.Printf("You can drill %d holes with the folllowing parameters:\n", possibility.Pieces)
		fmt.Printf("\tDrill ⌀: %dmm\tdistance : %dmm\n", possibility.Diameter, possibility.GetDistanceBetweenCenters())
	}
}

func printAsCSV(cuts []cut_dimensions.CutDimensions) {
	fmt.Println("Drill diameters,Distance betwen points,Pieces")
	for _, cut := range cuts {
		fmt.Printf("%d,%d,%d\n", cut.Diameter, cut.GetDistanceBetweenCenters(), cut.Pieces)
	}
}

func printToGoogleSheets(cuts []cut_dimensions.CutDimensions) {
	fmt.Println("Drill diameters\tDistance betwen points\tPieces")
	for _, cut := range cuts {
		fmt.Printf("%d\t%d\t%d\n", cut.Diameter, cut.GetDistanceBetweenCenters(), cut.Pieces)
	}
}

func showInfo() {
	fmt.Println("Width:", width, "mm")
	fmt.Println("Margin:", margin, "mm")
	fmt.Println("Available length: ", availableLength, "mm")
}
