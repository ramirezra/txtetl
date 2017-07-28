package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func file2Bytes(file *os.File) {

}
func main() {
	// var value float64
	// Open file
	// path := "TERMSA9.bak"
	// file, err := os.Open(path)
	// if err != nil {
	// 	log.Fatal("Error while opening file:", err)
	// }
	// defer file.Close()

	// var pi float64
	// err = binary.Read(file, binary.LittleEndian, &pi)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	data, err := ioutil.ReadFile("TERMSA9.bak")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", data)

	var y map[string]interface{}
	json.Unmarshal(data, &y)
	fmt.Printf("%+s\n", y)
	// contents, err := ioutil.ReadFile(path)
	// buf := bytes.NewReader(contents)
	// binary.Read(buf, binary.BigEndian, &value)
	// fmt.Println(pi)
	// fmt.Printf("%s opened\n", path)

	// // Check to see if it is a BBX file.
	// formatName := readNextBytes(file, 7)
	// // fmt.Printf("Parsed format: %s\n", formatName)
	// if string(formatName) != "<<bbx>>" {
	// 	log.Fatal("Provided file is not in correct format.")
	// }

	// data := readNextBytes(file, 10000)
	// fmt.Printf(string(data))
	// buffer := bytes.NewBuffer(data)
	// err = binary.Read(buffer, binary.LittleEndian, &values)

	// r := strings.NewReader(file)
	// if err != nil {
	// 	log.Fatal("binary.Read failed:", err)
	// }
	// fmt.Println()
	// data := make([]byte, 5000)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var pi float64
	// buf := bytes.NewReader(data)
	// binary.Read(buf, binary.LittleEndian, &pi)
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
