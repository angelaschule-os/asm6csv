package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"

	"github.com/gosimple/slug"
)

// 6 arrays holding the data
// first line for the csv header
var classes = [][]string{
	{
		"class_id",
		"class_number",
		"course_id",
		"instructor_id",
		"instructor_id_2",
		"instructor_id_3",
		"location_id",
	},
}

var courses = [][]string{
	{
		"course_id",
		"course_number",
		"course_name",
		"location_id",
	},
}

var locations = [][]string{
	{
		"location_id",
		"location_name",
	},
}

var rosters = [][]string{
	{
		"roster_id",
		"class_id",
		"student_id",
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

// write data into csv file
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

// write a template file
// see: template/t flag with ">" pipe
// enter your data into the template file
// TODO: filename
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

// check for duplicates
func findInSlice(slice [][]string, id string) bool {
	for _, row := range slice[1:] {
		if row[0] == id {
			return true
		}
	}
	return false
}

func generateSixpack() {
	// read the template file
	// TODO: get filename from commandline or dropping file
	inFile, err := os.Open("template.csv")
	if err != nil {
		log.Fatalln(err)
	}

	r := csv.NewReader(inFile)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// parse the data
	// drop header line
	for _, row := range records[1:] {

		// columns which are needed several times
		studentID := row[0]
		locationID := slug.Make(row[8])

		if !findInSlice(classes, row[3]) {
			classes = append(classes, []string{row[3], "", row[4], "", "", "", locationID})
		}

		if !findInSlice(courses, row[4]) {
			courses = append(courses, []string{row[4], "", "", locationID})
		}

		if !findInSlice(locations, locationID) {
			locations = append(locations, []string{locationID, row[8]})
		}

		if !findInSlice(rosters, row[5]) {
			rosters = append(rosters, []string{row[5], row[3], studentID})
		}

		if !findInSlice(staff, row[6]) {
			staff = append(staff, []string{row[6], "", "Lehrer", row[6], "", "", locationID})
		}

		students = append(students, []string{studentID, "", row[1], "", row[2], "", "", "", row[7], locationID})
	}

	// write the 6 files for upload
	writeCsv(classes, "classes.csv")
	writeCsv(courses, "courses.csv")
	writeCsv(locations, "locations.csv")
	writeCsv(rosters, "rosters.csv")
	writeCsv(staff, "staff.csv")
	writeCsv(students, "students.csv")
}

func main() {

	// cmd flags
	var template bool

	flag.BoolVar(&template, "template", false, "Generate csv template")
	flag.BoolVar(&template, "t", false, "Generate csv template (shorthand).")

	flag.Parse()

	// flag switch
	if template {
		generateTemplate()
	} else {
		generateSixpack()
	}
}
