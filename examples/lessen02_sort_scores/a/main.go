package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type scoreRecord struct {
	Name    string
	Chinese int
	Math    int
	English int
}

func main() {
	fmt.Println("\n7年9班期末成绩单")
	fmt.Println("================================")
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	r := csv.NewReader(csvfile)

	// Iterate through the records
	fmt.Println("姓名\t语文\t数学\t外语")
	records := make([]*scoreRecord, 0)
	i := 0
	line := ""
	for {
		// Read each record from csv
		i = i + 1
		line = fmt.Sprintf("line-%02d", i)
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%s\t%s\t%s\t%s\n", record[0], record[1], record[2], record[3])

		r := &scoreRecord{Name: record[0]}
		var score = 0
		score, err = strconv.Atoi(strings.TrimSpace(record[1]))
		if err != nil {
			log.Fatal(line, err)
		}
		r.Chinese = score

		score, err = strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			log.Fatal(line, err)
		}
		r.Math = score

		score, err = strconv.Atoi(strings.TrimSpace(record[3]))
		if err != nil {
			log.Fatal(line, err)
		}
		r.English = score
		records = append(records, r)
		// fmt.Printf("%#v\n", r)
	}
	fmt.Println()

	sort.Slice(records, func(i, j int) bool {
		return records[i].Chinese > records[j].Chinese
	})

	outputRecords := make([][]string, 0, 20)
	for i, r := range records {
		fmt.Printf("%02d\t%v\t%v\n", i+1, r.Name, r.Chinese)
		record := []string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%v", r.Name),
			fmt.Sprintf("%d", r.Chinese),
		}
		outputRecords = append(outputRecords, record)
	}

	fmt.Println()

	outputfile, err := os.Create("chinese.csv")
	checkError("Cannot create file", err)
	defer outputfile.Close()

	writer := csv.NewWriter(outputfile)
	defer writer.Flush()

	for _, value := range outputRecords {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
