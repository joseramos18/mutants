package helper
import (
	"fmt"
	"mutants/models"
	"strings"

	"errors"

	"github.com/go-playground/validator/v10"
	"strconv"
)

func VerifyDNA(dna *models.DNA) (bool, error){
	var mapLetters [6][6]string
	dnaMutantCounter := 0
	if err := constructMapLetter(dna, &mapLetters); err != nil {
		return false, err
	}
	return (verifyRow(dna, &mapLetters, &dnaMutantCounter) ||
	verifyByColumn(&mapLetters, &dnaMutantCounter) ||
	verifyObliqueString(&mapLetters, &dnaMutantCounter)),nil
		
}

func constructMapLetter(dna *models.DNA, mapLetters *[6][6]string,) error{
	validate := validator.New()
	for i, rowletters := range dna.Letters {
		var rowString models.RowLetters
		rowletters = strings.Replace(rowletters," ", "",50)
		rowString.Letters = strings.Split(rowletters, "")
		if err:= validate.Struct(rowString); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				errorTag := err.Tag()
				fmt.Println(i)
				if errorTag == "eq=A|eq=T|eq=G|eq=C"{
					return errors.New("Solo se permiten las letras A. T, G, C. Error encontrado en el grupo: " + strconv.Itoa(i+1))
				} else {
					return errors.New("La longitud de Dna es incorrecto en el grupo: "+ strconv.Itoa(i+1))
				}	
			}
		} 
		for j, letter := range rowString.Letters {
			mapLetters[i][j] = strings.ToUpper(letter)
		}
	}
	return nil
}

func verifyRow(dna *models.DNA, mapLetters *[6][6]string, dnaMutantCounter *int) bool {
	for _,rowletters := range dna.Letters {
		rowString := strings.Split(rowletters, "")
		if verifyElement(rowString,dnaMutantCounter){
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