package main

import (
	"encoding/csv"
	"os"
)

func mapToCsvFileCreator(recordsMap [][]string, fileName string) error {

	_, err := os.Stat(fileName)
	if !os.IsNotExist(err) {
		err = os.Remove(fileName)

		if err != nil {
			return err
		}
	}

	f, err := os.Create(fileName)

	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range recordsMap {
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}
