package helper_test
import (
	"testing"


   // "github.com/stretchr/testify/assert"
	"mutants/helper"
    "mutants/models"
)



func TestVerifyDNA(t *testing.T) {

	//mutant1 := models.DNA{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}
   // mutant1.DNA = {"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}
	
   	dna := models.DNA {
	Letters : []string{"ATGbGA","CABTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"},				
}

	actualResult := helper.VerifyDNA(&dna)
	if !actualResult {
		t.Error("")
	}

    if actualResult != true {
		 t.Log("Es un Humano")
    }else{
		t.Log("Es mutante")
	}
    //assert.Equal(t, expectedResult, actualResult)
}

