package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"

	"github.com/gosimple/slug"
)

var locations = [][]string{
	{
		"location_id",
		"location_name",
	},
}

var staff = [][]string{
	{
		"person_id",
		"person_number",
		"first_name",
		"middle_name",
		"last_name",
		"email_address",
		"sis_username",
		"location_id",
	},
}

var students = [][]string{
	{
		"person_id",
		"person_number",
		"first_name",
		"middle_name",
		"last_name",
		"grade_level",
		"email_address",
		"sis_username",
		"password_policy",
		"location_id",
	},
}

func writeCsv(data [][]string, filename string) {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer outFile.Close()

	w := csv.NewWriter(outFile)
	w.WriteAll(data) // calls Flush internally
	if err := w.Error(); err != nil {
		log.Fatalln(err)
	}
}

func generateTemplate() {
	records := [][]string{
		{
			"student_id",
			"first_name",
			"last_name",
			"class_id",
			"course_id",
			"roster_id",
			"teacher_id",
			"password_policy", // 4, 6, 8 or empty
			"location_name",
		},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln(err)
	}
}

func generateSixpack() {
	inFile, err := os.Open("template.csv")
	if err != nil {
		log.Fatalln(err)
	}

	r := csv.NewReader(inFile)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// drop header line
	for _, row := range records[1:] {
		locationID := slug.Make(row[8])
		studentID := row[0]
		locations = append(locations, []string{locationID, row[8]})
		staff = append(staff, []string{row[6], "", "Lehrer", row[6], "", "", locationID})
		students = append(students, []string{studentID, "", row[1], "", row[2], "", "", "", "", locationID})
	}
	// fmt.Println(locations)
	writeCsv(locations, "locations.csv")
	writeCsv(staff, "staff.csv")
	writeCsv(students, "students.csv")
}

func main() {

	var template bool

	flag.BoolVar(&template, "template", false, "Generate csv template")
	flag.BoolVar(&template, "t", false, "Generate csv template (shorthand).")

	flag.Parse()

	if template {
		generateTemplate()
	} else {
		generateSixpack()
	}
}
