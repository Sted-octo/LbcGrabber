package main

import (
	"fmt"
	"log"
)

func main() {

	startIndex, resultLength, token, fileName, cookie, keywords := argsManager()

	body, err := littleBigConnectionApiListGetter(*startIndex, *resultLength, *token, *cookie, *keywords)
	if err != nil {
		log.Fatalln(err)
	}

	lbcResult, err := jsonToLittleBigConnectionResultObjectConverter(body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Nombre de résultats disponibles : %d\n", lbcResult.Count)

	missionsDetails := littlebigconnectionMissionsInformationsGetter(lbcResult.Results, *token, *cookie)

	recordsMap, err := littleBigConnectionResultObjectToStringMapConverter(lbcResult, missionsDetails)
	if err != nil {
		log.Fatalln(err)
	}

	err = mapToCsvFileCreator(recordsMap, *fileName)

	if err != nil {
		log.Fatalln(err)
	}
}
