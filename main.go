package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type SampleData struct {
	Avg_QPA_Given   float64  `json:"avg_qpa_given"`
	Salary          *float64 `json:"salary"`
	Children        int      `json:"children"`
	Rating          *float64 `json:"rating"`
	Avg_Grade_Given *float64 `json:"avg_grade_Given"`
}

func main() {
	deletedRows := DeleteRow()
	log.Println("Deleted Rows")
	PrintResult(deletedRows)
	nullRows := SetRowNull()
	log.Println("Nulled Rows")
	PrintResult(nullRows)
}

func DeleteRow() []SampleData {
	file, err := os.Open("sampleData.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	// Read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	finalData := []SampleData{}
	for dec.More() {
		var m SampleData

		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		if m.Salary == nil {
			continue
		}
		if m.Rating == nil {
			continue
		}
		if m.Avg_Grade_Given == nil {
			continue
		}
		finalData = append(finalData, m)
	}

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	return finalData
}

func SetRowNull() []SampleData {
	file, err := os.Open("sampleData.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	// Read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	finalData := []SampleData{}
	for dec.More() {
		var m SampleData

		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		if m.Salary == nil {
			m.Salary = new(float64)
			*m.Salary = 0.0
		}
		if m.Rating == nil {
			m.Rating = new(float64)
			*m.Rating = 0.0
		}
		if m.Avg_Grade_Given == nil {
			m.Avg_Grade_Given = new(float64)
			*m.Avg_Grade_Given = 0.0
		}

		finalData = append(finalData, m)
	}

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	return finalData
}

func PrintResult(sample_data []SampleData) {
	for i, data := range sample_data {

		fmt.Printf("DATA [%v]: | %.2f | %.2f | %d | %.2f | %.2f |\n",
			i,
			data.Avg_QPA_Given,
			*data.Salary,
			data.Children,
			*data.Rating,
			*data.Avg_Grade_Given)
	}
}

// func DeleteRow() []SampleData {
// 	file, err := os.Open("sampleData.json")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	defer file.Close()

// 	dec := json.NewDecoder(file)

// 	// Read open bracket
// 	t, err := dec.Token()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%T: %v\n", t, t)

// 	for dec.More() {
// 		var m SampleData

// 		err := dec.Decode(&m)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if m.Salary == nil {
// 			m.Salary = new(float64)
// 			*m.Salary = 0.0
// 		}
// 		if m.Rating == nil {
// 			m.Rating = new(float64)
// 			*m.Rating = 0.0
// 		}
// 		if m.Avg_Grade_Given == nil {
// 			m.Avg_Grade_Given = new(float64)
// 			*m.Avg_Grade_Given = 0.0
// 		}
// 	}

// 	t, err = dec.Token()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%T: %v\n", t, t)
// }
