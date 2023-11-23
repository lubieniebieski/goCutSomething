package cut_dimensions

import (
	"fmt"
	"math"
)

type CutDimensions struct {
	Diameter int
	Spacing  int
	Pieces   int
}

func (c CutDimensions) AreNice() bool {
	return math.Mod(float64(c.GetDistanceBetweenCenters()), 10.0) == 0
}

func (c CutDimensions) GetDistanceBetweenCenters() int {
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

func PossibleCutDimensions(materialWidth int, desiredDiameter int, desiredSpacing int) (result []CutDimensions) {

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
				if newDimensions.Pieces > 0 && newDimensions.AreNice() {
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

func FindBestCut(cuts []CutDimensions, minLength int, minPieces int, length int) CutDimensions {
	var bestCut CutDimensions
	percentage := 0.0
	for _, cut := range cuts {
		if cut.Pieces < minPieces {
			continue
		}
		if cut.GetDistanceBetweenCenters() < minLength {
			continue
		}
		newPercentage := float64(cut.Pieces*cut.GetDistanceBetweenCenters()) / float64(length)
		fmt.Printf("%d cuts by %dmm (%.1f%%)\n", cut.Pieces, cut.GetDistanceBetweenCenters(), newPercentage*100)
		if newPercentage > percentage {
			percentage = newPercentage
			bestCut = cut
		}

	}
	return bestCut

}
