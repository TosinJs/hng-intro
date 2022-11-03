package main

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

//Struct CHIP_007 Represents the keys and values in the CHIP_007 Specs
type CHIP_007 struct {
	Format           string `json:"format"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	MintingTools     string `json:"minting_tools"`
	SensitiveContent bool   `json:"sensitive_content"`
	SeriesNumber     int    `json:"series_number"`
	SeriesTotal      int    `json:"series_total"`
	Attributes       []struct {
		TraitType string `json:"trait_type"`
		Value     string `json:"value"`
		MinValue  int    `json:"min_value"`
		MaxValue  int    `json:"max_value"`
	} `json:"attributes"`
	Collection struct {
		Name       string `json:"name"`
		Id         string `json:"id"`
		Attributes []struct {
			Type  string
			Value string
		} `json:"attributes"`
	} `json:"collection"`
	Data struct {
		ExampleData string `json:"example_data"`
	} `json:"data"`
}

//parseCSVIntoJSONAndHash parses each row in the csv into a json string and return the SHA256 hash of the json string
func parseCSVIntoJSONAndHash(arr []string) (string, error) {
	var chip_007 CHIP_007
	chip_007.Format = "CHIP_007"
	chip_007.Name = arr[2]
	chip_007.Description = arr[3]
	chip_007.Attributes = []struct {
		TraitType string "json:\"trait_type\""
		Value     string "json:\"value\""
		MinValue  int    "json:\"min_value\""
		MaxValue  int    "json:\"max_value\""
	}{
		{
			TraitType: "Gender",
			Value:     arr[4],
		},
	}
	jsonVal, err := json.Marshal(chip_007)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	_, err = h.Write(jsonVal)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
	filename := flag.String("csv", "nft.csv", "Enter the CSV File Path")
	flag.Parse()
	flag.Usage()

	file, err := os.Open(*filename)
	if err != nil {
		handleError(err.Error())
	}
	newCSVFile, err := os.Create("nftNew.csv")
	if err != nil {
		handleError(err.Error())
	}
	defer file.Close()
	defer newCSVFile.Close()

	csvReader := csv.NewReader(file)
	//A new CSV File Is Created Which COntains a New Column "SHA256"
	csvWriter := csv.NewWriter(newCSVFile)
	for {
		csvFile, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			handleError(err.Error())
		}
		//Check to See if this is the first row (Hardcoded due to lack of data)
		if csvFile[1] == "Filename" {
			csvFile = append(csvFile, "SHA256")
			csvWriter.Write(csvFile)
			continue
		}
		//Check to see if this row has a name value. If theres no name value, that row will be ignored as you cant have a nameless NFT
		//HArdCOded due to lack of data
		if csvFile[1] == "" {
			csvWriter.Write(csvFile)
			continue
		}
		byteRet, err := parseCSVIntoJSONAndHash(csvFile)
		if err != nil {
			handleError(err.Error())
		}
		csvFile = append(csvFile, string(byteRet))
		err = csvWriter.Write(csvFile)
		if err != nil {
			handleError(err.Error())
		}
	}
}

func handleError(err string) {
	fmt.Fprintf(os.Stderr, err)
	os.Exit(1)
}
