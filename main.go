package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type CutDimensions struct {
	Diameter int
	Spacing  int
	Pieces   int
}

func (c CutDimensions) areNice() bool {
	return math.Mod(float64(c.getDistanceBetweenCenters()), 10.0) == 0
}

func (c CutDimensions) getDistanceBetweenCenters() int {
	return c.Spacing + c.Diameter
}

func generateIntegers(desiredInt int, count int) []int {
	result := make([]int, count*2+1)

	for i := 0; i < (count*2 + 1); i++ {
		result[i] = desiredInt - count + i
	}
	return result

}

func cuttableWithoutMargin(materialWidth int, diameter int, spacing int) bool {
	return math.Mod(float64(materialWidth-diameter), float64(diameter+spacing)) == 0
}

func possibleCutDimensions(materialWidth int, desiredDiameter int, desiredSpacing int) (result []CutDimensions) {

	possibleSpacings := generateIntegers(desiredSpacing, 100)
	possibleDiameters := generateIntegers(desiredDiameter, 100)
	for _, spacing := range possibleSpacings {
		for _, diameter := range possibleDiameters {
			if spacing < 1 || diameter < 1 {
				continue
			}
			if cuttableWithoutMargin(materialWidth, diameter, spacing) {
				newDimensions := CutDimensions{Diameter: diameter, Spacing: spacing}
				newDimensions.Pieces = (materialWidth - diameter) / (diameter + spacing)
				if newDimensions.Pieces > 0 && newDimensions.areNice() {
					result = append(result, newDimensions)
				}
			}
		}
	}
	if len(result) > 0 {
		return result
	}

	println("returning empty")
	return []CutDimensions{}

}

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <width> <margin> <diameter>")
		return
	}

	// Parse the arguments
	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid width:", os.Args[1])
		return
	}

	margin, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid margin:", os.Args[2])
		return
	}

	desiredDiameter, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid diameter:", os.Args[3])
		return
	}
	availableLength := width - margin
	// Print the values
	fmt.Println("Width:", width, "mm")
	fmt.Println("Margin:", margin, "mm")
	fmt.Println("Available length: ", availableLength, "mm")
	desiredSpacing := 100

	possibilities := possibleCutDimensions(availableLength, desiredDiameter, desiredSpacing)

	// printAsCSV(possibilities)
	// printToGoogleSheets(possibilities)
	// printOutput(possibilities)
	bestCut := findBestCut(possibilities, 15, 2, availableLength)
	fmt.Printf("Best cut: %dmm, %d pieces\n", bestCut.getDistanceBetweenCenters(), bestCut.Pieces)

}

func printOutput(cuts []CutDimensions) {
	for _, possibility := range cuts {

		fmt.Printf("You can drill %d holes with the folllowing parameters:\n", possibility.Pieces)
		fmt.Printf("\tDrill ⌀: %dmm\tdistance : %dmm\n", possibility.Diameter, possibility.getDistanceBetweenCenters())
	}
}

func printAsCSV(cuts []CutDimensions) {
	fmt.Println("Drill diameters,Distance betwen points,Pieces")
	for _, cut := range cuts {
		fmt.Printf("%d,%d,%d\n", cut.Diameter, cut.getDistanceBetweenCenters(), cut.Pieces)
	}
}

func printToGoogleSheets(cuts []CutDimensions) {
	fmt.Println("Drill diameters\tDistance betwen points\tPieces")
	for _, cut := range cuts {
		fmt.Printf("%d\t%d\t%d\n", cut.Diameter, cut.getDistanceBetweenCenters(), cut.Pieces)
	}
}

func findBestCut(cuts []CutDimensions, minLength int, minPieces int, length int) CutDimensions {
	var bestCut CutDimensions
	percentage := 0.0
	for _, cut := range cuts {
		if cut.Pieces < minPieces {
			continue
		}
		if cut.getDistanceBetweenCenters() < minLength {
			continue
		}
		newPercentage := float64(cut.Pieces*cut.getDistanceBetweenCenters()) / float64(length)
		fmt.Printf("%d cuts by %dmm (%.1f%%)\n", cut.Pieces, cut.getDistanceBetweenCenters(), newPercentage*100)
		if newPercentage > percentage {
			percentage = newPercentage
			bestCut = cut
		}

	}
	return bestCut

}
