package main

import "flag"

func argsManager() (*int64, *int64, *string, *string, *string, *string) {
	startIndex := flag.Int64("startIndex", 0, "Index de départ, 0 par défaut")
	resultLength := flag.Int64("resultLength", 10, "Nombre de résultat, 10 par défaut")
	token := flag.String("token", "", "Token d'accès à l'API, obligatoire")
	fileName := flag.String("fileName", "out.csv", "file name for the output csv file to create")
	cookie := flag.String("cookie", "", "Cookie d'accès à l'API, obligatoire")
	keywords := flag.String("keywords", "", "enum of keyword to look for separated by |")

	flag.Parse()

	return startIndex, resultLength, token, fileName, cookie, keywords
}
