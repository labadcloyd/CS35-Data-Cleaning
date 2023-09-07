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
	meanRows := RowMean()
	log.Println("Mean Rows")
	PrintResult(meanRows)
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

func RowMean() []SampleData {
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

	draftData := []SampleData{}
	for dec.More() {
		var m SampleData

		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		draftData = append(draftData, m)
	}
	Salary_Mean := 0.0
	Salary_Length := 0
	Rating_Mean := 0.0
	Rating_Length := 0
	Avg_Grade_Given_Mean := 0.0
	Avg_Grade_Given_Length := 0
	// Computing mean
	for i, m := range draftData {
		if m.Salary != nil {
			Salary_Mean += *m.Salary
			Salary_Length += i
		}
		if m.Rating != nil {
			Rating_Mean += *m.Rating
			Rating_Length += i
		}
		if m.Avg_Grade_Given != nil {
			Avg_Grade_Given_Mean += *m.Avg_Grade_Given
			Avg_Grade_Given_Length += i
		}
	}
	Salary_Mean = Salary_Mean / float64(Salary_Length)
	Rating_Mean = Rating_Mean / float64(Rating_Length)
	Avg_Grade_Given_Mean = Avg_Grade_Given_Mean / float64(Avg_Grade_Given_Length)

	// Saving mean
	finalData := []SampleData{}
	for _, m := range draftData {
		newData := SampleData{
			Avg_QPA_Given:   m.Avg_QPA_Given,
			Salary:          m.Salary,
			Children:        m.Children,
			Rating:          m.Rating,
			Avg_Grade_Given: m.Avg_Grade_Given,
		}
		if newData.Salary == nil {
			newData.Salary = new(float64)
			*newData.Salary = Salary_Mean
		}
		if newData.Rating == nil {
			newData.Rating = new(float64)
			*newData.Rating = Rating_Mean
		}
		if newData.Avg_Grade_Given == nil {
			newData.Avg_Grade_Given = new(float64)
			*newData.Avg_Grade_Given = Avg_Grade_Given_Mean
		}
		finalData = append(finalData, newData)
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
