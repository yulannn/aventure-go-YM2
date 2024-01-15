package backend

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

func PersoLoad() ([]Personnage, error) {
    fileData, err := os.ReadFile("perso.json")
    if err != nil {

        return nil, err
    }

    var forms []Personnage
    if len(fileData) != 0 {
        err = json.Unmarshal(fileData, &forms)
        if err != nil {
			fmt.Println(err)
            return nil, err
        }
    }
    return forms, nil
}

func RandomID() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(9000) + 1000
}

func Contains(list []int, value int) bool {
    for _, v := range list {
        if v == value {
            return true
        }
    }
    return false
}

func PersoIDs(filePath string) ([]int, error) {
	aventurierIDs := []int{}

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}


	var aventuriers []Personnage
	err = json.Unmarshal(fileContent, &aventuriers)
	if err != nil {
		return nil, err
	}

	for _, aventurier := range aventuriers {
		idInt, err := strconv.Atoi(aventurier.ID)
		if err != nil {
			return nil, err
		}
		aventurierIDs = append(aventurierIDs, idInt)
	}

	return aventurierIDs, nil
}