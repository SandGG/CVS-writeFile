package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	writeFile()
	readFile()
}

func writeFile() {
	var names = [][]string{
		{"first_name", "last_name", "user_name"},
		{"Marco", "Diaz", "marco123"},
		{"Miguel", "Trejo", "magicMike"},
		{"Ana", "Gomez", "anitaQ"},
	}

	var file, err = os.OpenFile("./files/names.csv", os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var w = csv.NewWriter(file)
	for _, item := range names {
		var err = w.Write(item)
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush() //Write stored data

	fmt.Println("File Updated Successfully")
}

func readFile() {
	var file, err = os.OpenFile("./files/names.csv", os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var r = csv.NewReader(file)
	for {
		var data, err = r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		for i := range data {
			fmt.Println(data[i])
		}
	}
}
