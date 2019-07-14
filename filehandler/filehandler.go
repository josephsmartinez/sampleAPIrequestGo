package filehandler

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

// BUFFLIMIT - can be increased if needed
var (
	BUFFLIMIT = 1024
)

// PrintToFile -
func PrintToFile(data *string, filename string) (bool, int) {
	var isDataWritten bool
	f, err := os.Create(filename)
	if err != nil {
		isDataWritten = false
		log.Fatal(err)
	}
	buffSize := (BUFFLIMIT)
	w := bufio.NewWriterSize(f, buffSize)
	numBytesWritten, err := w.WriteString(*data)
	if err != nil {
		isDataWritten = false
		log.Fatal(err)
	}
	isDataWritten = true
	w.Flush()
	return isDataWritten, numBytesWritten
}

// ReadJSONfile -
func ReadJSONfile(fileName string) []byte {

	pFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer pFile.Close()
	byteValue, _ := ioutil.ReadAll(pFile)

	return byteValue
}
