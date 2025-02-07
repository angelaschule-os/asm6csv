package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	// used for computing of ids
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

// write data into filename in csv format (comma seperated)
func writeCsv(data [][]string, filename string) {
	// TODO: check if file exists
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	// make sure that file is closed
	defer outFile.Close()

	w := csv.NewWriter(outFile)
	w.WriteAll(data) // calls Flush internally
	if err := w.Error(); err != nil {
		log.Fatalln(err)
	}
}

// write a template to filename
// enter your data into the template file
func generateTemplate(filename string) {
	// columns header, add one row per student
	records := [][]string{
		{
			"student_id",
			"first_name",
			"last_name",
			"class_id",
			"course_id",
			"teacher_id",
			"password_policy", // 4, 6, 8 or empty
			"location_name",
		},
	}
	writeCsv(records, filename)
}

// idexes if template fields
const (
	STUDENTID = 0
	FIRSTNAME = 1
	LASTNAME  = 2
	CLASSID   = 3
	COURSEID  = 4
	TEACHERID = 5
	PWPOLICY  = 6
	LOCATION  = 7
)

// check for duplicates
func findInSlice(slice [][]string, id string) bool {
	for _, row := range slice[1:] {
		if row[0] == id {
			return true
		}
	}
	return false
}

func generateSixpack(filename string) {
	// read the template file
	inFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	r := csv.NewReader(inFile)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// parse all the data to arrays
	// drop header line
	for _, row := range records[1:] {

		// columns which are needed several times
		studentID := row[STUDENTID]
		locationID := slug.Make(row[LOCATION])

		// no duplicates
		if !findInSlice(classes, row[CLASSID]) {
			classes = append(classes, []string{row[CLASSID], "", row[COURSEID], "", "", "", locationID})
		}

		if !findInSlice(courses, row[COURSEID]) {
			courses = append(courses, []string{row[COURSEID], "", "", locationID})
		}

		if !findInSlice(locations, locationID) {
			locations = append(locations, []string{locationID, row[LOCATION]})
		}

		// computing generic roster_id from class_id and student_id
		rosterID := row[CLASSID] + "-" + studentID
		if !findInSlice(rosters, rosterID) {
			rosters = append(rosters, []string{rosterID, row[CLASSID], studentID})
		}

		if !findInSlice(staff, row[TEACHERID]) {
			staff = append(staff, []string{row[TEACHERID], "", "Lehrer", row[TEACHERID], "", "", locationID})
		}

		students = append(students, []string{studentID, "", row[FIRSTNAME], "", row[LASTNAME], "", "", "", row[PWPOLICY], locationID})
	}

	// write the arrays into 6 files for upload
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

	// pretty usage message
	flag.Usage = func() {
		fmt.Printf("Usage: asm6csv [options] filename.csv\nOptions:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// flag switch
	if template {
		// check is a filename is passed on the commandline
		if len(flag.Args()) > 0 {
			generateTemplate(flag.Args()[0])
		} else {
			log.Fatalln("No filename given.")
		}
	} else {
		// check is a filename is passed on the commandline
		if len(flag.Args()) > 0 {
			generateSixpack(flag.Args()[0])
		} else {
			log.Fatalln("No filename given.")
		}
	}
}
