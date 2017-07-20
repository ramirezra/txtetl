package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LedgerDetail struct {
	AccountNumber  byte
	Annotation     string
	CreditAmount   float64
	DebitAmount    float64
	DocumentNumber string
	EffectiveDate  string
	EntryDate      string
	EntryNumber    int
	FiscalPeriod   string
	LineReference  int
	SequenceNumber int
	SourceJournal  string
}

func main() {
	data := jsonify()

	// var transaction LedgerDetail
	// var transactions []LedgerDetail

	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// jsonFile, err := os.Create("data.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer jsonFile.Close()
	// jsonFile.Write(jsonData)
	// fmt.Printf("%T", data)
	// fmt.Printf("%T", jsonData)
}

func jsonify() {
	// Open file
	file, err := os.Open("testquery.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Scan .txt file into buffer and check for errors.
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	// scanner.Split(splitcolon)

	var newtext []string
	newtext = append(newtext, "[")
	for scanner.Scan() {
		lines := scanner.Text()

		index := strings.Index(lines, ":")

		// newline := strings.Split(lines, ":")
		if index > 0 {
			category := lines[:index]
			values := lines[index+1:]
			// fmt.Printf("%q:%q,", category, values)
			newline := category + ":" + values + ","
			newtext = append(newtext, newline)
		} else {
			newtext = append(newtext, "},\n{")
		}

		// newtext = append(newtext, newline)
	}

	fmt.Println(newtext)

	// for i := 0; i < 12; i++ {
	// 	fmt.Printf(record[i][0])

	// }

}

func csvformat() string {
	// Open file
	file, err := os.Open("testquery.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Scan .txt file into buffer and check for errors.
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	// scanner.Split(splitcolon)

	var newtext []string
	for scanner.Scan() {
		lines := scanner.Text()

		index := strings.Index(lines, ":")
		// newline := strings.SplitAfter(lines, ":")

		// newline := lines[index+1:]
		// fmt.Printf("%v - %q\n", index, newline)
		// newtext = append(newtext, newline)

		if index > 0 {
			values := lines[index+1:]
			// fmt.Printf("%q:%q,", category, values)
			newline := values + ","
			newtext = append(newtext, newline)
		} else {
			length := len(newtext)
			newtext = append(newtext[:length-1], strings.TrimRight(newtext[length-1], ","))
			newtext = append(newtext, "\n")

		}

	}
	bytes := strings.Join(newtext, "")
	return bytes
}

func splitcolon(data []byte, atEOF bool) (int, []byte, error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ':' {
			return i + 1, data[:i], nil
		}
	}
	return 0, data, bufio.ErrFinalToken
}
