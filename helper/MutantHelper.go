package helper
import (
	"mutants/models"
	"strings"
)

func VerifyDNA(dna *models.DNA) bool{
	var mapLetters [6][6]string
	dnaMutantCounter := 0
	return (verifyByRow(dna, &mapLetters, &dnaMutantCounter) ||
	verifyByColumn(&mapLetters, &dnaMutantCounter) ||
	verifyObliqueString(&mapLetters, &dnaMutantCounter))
		
}

func verifyByRow(dna *models.DNA, mapLetters *[6][6]string, dnaMutantCounter *int) bool {
	for i, rowletters := range dna.Letters {
		rowString := strings.Split(rowletters, "")
		for j, letter := range rowString {
			mapLetters[i][j] = strings.ToUpper(letter)
		}
		if verifyElement(rowString, dnaMutantCounter) {
			return true
		}
	}
	return false
}

func verifyByColumn(mapLetters *[6][6]string, dnaMutantCounter *int) bool {
	var columnString []string
	for column := 0; column < 6; column++ {
		for line := 0; line < 6; line++ {
			columnString = append(columnString, mapLetters[line][column])
		}
		if verifyElement(columnString, dnaMutantCounter) {
			return true
		}
		columnString = columnString[:0]
	}
	return false
}

func verifyObliqueString(mapLetters *[6][6]string, dnaMutantCounter *int) bool {
	var arrayObliqueString [10][]string
	for j := 0; j < 6; j++ {
		x := len(mapLetters) - (j + 1)
		arrayObliqueString[0] = append(arrayObliqueString[0], mapLetters[j][j])
		arrayObliqueString[1] = append(arrayObliqueString[1], mapLetters[j][x])
		if j > 0 {
			y := j - 1
			createObliqueArray(2, x, y, j, mapLetters, &arrayObliqueString)
		}
		if j > 1 {
			y := j - 2
			createObliqueArray(6, x, y, j, mapLetters, &arrayObliqueString)
		}
	}
	for _, array := range arrayObliqueString {
		if verifyElement(array, dnaMutantCounter) {
			return true
		}
	}
	return false
}

func createObliqueArray(i int, x int, y int, j int, mapLetters *[6][6]string, arrayString *[10][]string) {
	arrayString[i] = append(arrayString[i], mapLetters[y][j])
	arrayString[i+1] = append(arrayString[i+1], mapLetters[j][y])
	arrayString[i+2] = append(arrayString[i+2], mapLetters[y][x])
	arrayString[i+3] = append(arrayString[i+3], mapLetters[x][y])
}

func verifyElement(letters []string, dnaMutantCounter *int) bool {
	var coincidence int
	for i := 0; i < len(letters); i++ {
		j := i + 1
		if j != len(letters) {
			if letters[i] == letters[j] {
				coincidence = coincidence + 1
			} else {
				coincidence = 0
			}
			if coincidence == 3 {
				*dnaMutantCounter++
			}
		}
	}
	if *dnaMutantCounter > 1 {
		return true
	}
	return false
}