// This is a simple program to track absence of students in college. The user is assumed to be a lecturer or an academic administration staff in a university.
// There will be 3 data managed. Student's data, class schedule data and attendance log data. 

package main

import "fmt"

const NMAX = 999

type studentDetail struct {
	name string
	sid  int
	class string
}

type studentList struct {
	detail [NMAX]studentDetail
	count  int
}

// type StudentList will be used as student's data. Count determines the amount of students. 

type list [NMAX]studentDetail

type course struct {
	name       string
	present    list
	sick       list
	permission list
	alpha      list
}

type class struct {
	subjects [NMAX]course
	day string
}

type classlist [NMAX]class

// Class Schedule Data. The type Class lists out the courses each day. The type Classlist lists out the class schedule.

type log struct {
	student  studentDetail
	presence int
	absence  int
}

type loglist [NMAX]log

// type Loglist is used as attendance log data. Students listed as present in a course will have their presence count added.
// Students listed as sick, permission or alpha in a course will have their absence count added.		

func main() {
	var students studentList
	var schedule classlist
	var attendancelog loglist
	var choice, index int
	var sorted bool
	index = -1
	students.count = 0

	fmt.Println("== Welcome to SiPRESENSI - Student Attendance System ==")
	fmt.Println("Please select an option:")
	fmt.Println("1. Add Student")
	fmt.Println("2. Change Student Data")
	fmt.Println("3. Delete Student Data")
	fmt.Println("4. Search Student Data")
	fmt.Println("5. Add Class Schedule")
	fmt.Println("6. View Class Schedule")
	fmt.Println("7. Record Attendance")
	fmt.Println("8. View Attendance Log")
	fmt.Println("9. Sort Student Data")
	fmt.Println("-1. Exit")
	fmt.Scanln(&choice)

	for choice != -1 {
		switch choice {
		case 1:
			addStudent(&students, &attendancelog, &sorted)
		case 2:
			changeStudentData(&students, &attendancelog, &sorted, &index)
		case 3:
			deleteStudentData(&students, &attendancelog, &sorted, &index)
		case 4:
			searchStudentData(students, &index)
			fmt.Printf("\n")
			fmt.Printf("Student found at index: %d\n", index)
			fmt.Print("Press Enter to go back...")
			fmt.Scanln()
		case 5:
			addClassSchedule(&schedule)
		case 6:
			viewClassSchedule(&schedule)
		case 7:
			recordAttendance(&schedule, &attendancelog)
		case 8:
			viewAttendanceLog(&attendancelog)
		case 9:
			sortStudentData(&students, &attendancelog)
		default:
			fmt.Print("Invalid choice. Please select a valid option: ")
			fmt.Scan(&choice)
			continue
		}
		clearscreen()

		fmt.Println("== Welcome to SiPRESENSI - Student Attendance System ==")
		fmt.Println("Please select an option:")
		fmt.Println("1. Add Student")
		fmt.Println("2. Change Student Data")
		fmt.Println("3. Delete Student Data")
		fmt.Println("4. Search Student Data")
		fmt.Println("5. Add Class Schedule")
		fmt.Println("6. View Class Schedule")
		fmt.Println("7. Record Attendance")
		fmt.Println("8. View Attendance Log")
		fmt.Println("-1. Exit")
		fmt.Scanln(&choice)
	}

}

func addStudent(students *studentList, attendancelog *loglist, sorted *bool) {
	var index int
	var choice string
	clearscreen()
	choice = "Yes"

	index = students.count
	for choice != "No" {
		fmt.Print("Enter Student Name: ")
		fmt.Scan(&students.detail[index].name)
		fmt.Print("Enter Student ID: ")
		fmt.Scan(&students.detail[index].sid)
		fmt.Print("Enter Student Class: ")
		fmt.Scan(&students.detail[index].class)
		attendancelog[index].student.name = students.detail[index].name
		attendancelog[index].student.sid = students.detail[index].sid
		attendancelog[index].student.name = students.detail[index].class
		attendancelog[index].presence = 0
		attendancelog[index].absence = 0

		fmt.Print("Add another student? (Yes/No): ")
		fmt.Scan(&choice)
		index++
		fmt.Printf("\n")
	}
	students.count = index
	*sorted = false

	fmt.Println("Student(s) added successfully!")
	fmt.Print("Press Enter to go back...")
	fmt.Scanln()
}

func changeStudentData(students *studentList, attendancelog *loglist, sorted *bool, index *int) {
	clearscreen()
	if *index == -1 {
		fmt.Print("No student selected. Please search for a student first.\n")
		fmt.Print("Press Enter to go back...")
		fmt.Scanln()
		return
	} else {
		fmt.Printf("Current Name: %s\n", students.detail[*index].name)
		fmt.Printf("Current SID: %d\n", students.detail[*index].sid)
		fmt.Printf("\n")
		fmt.Print("Enter new name: ")
		fmt.Scan(&students.detail[*index].name)
		fmt.Print("Enter new SID: ")
		fmt.Scan(&students.detail[*index].sid)
		attendancelog[*index].student.name = students.detail[*index].name
		attendancelog[*index].student.sid = students.detail[*index].sid
		attendancelog[*index].presence = 0
		attendancelog[*index].absence = 0
		*sorted = false
	}

	fmt.Printf("\nStudent data updated, attendance log reset for this student.\n")
	fmt.Print("Press Enter to go back...")
	fmt.Scanln()
}

func deleteStudentData(students *studentList, attendancelog *loglist, sorted *bool, index *int) {
	if *index == -1 {
		fmt.Print("No student selected. Please search for a student first.\n")
		fmt.Print("Press Enter to go back...")
		fmt.Scanln()
		return
	} else {
		for i := *index; i < students.count-1; i++ {
			students.detail[i] = students.detail[i+1]
			attendancelog[i] = attendancelog[i+1]
		}
		students.count--
		*sorted = false
	}
	fmt.Printf("\nStudent data deleted successfully!\n")
	fmt.Print("Press Enter to go back...")
	fmt.Scanln()
}

func searchStudentData(students studentList, index *int) {
	var choice, name string
	var sid, i int
	clearscreen()
	fmt.Print("Search based on (name/sid): ")
	fmt.Scan(&choice)

	if choice == "name" {
		fmt.Print("Enter Student Name: ")
		fmt.Scan(&name)
		for i = 0; i < students.count && students.detail[i].name != name; i++ {
		}
	} else if choice == "sid" {
		fmt.Print("Enter Student ID: ")
		fmt.Scan(&sid)
		for i = 0; i < students.count && students.detail[i].sid != sid; i++ {
		}
	}
	if i < students.count {
		*index = i
	}
	fmt.Printf("\n")
}

func clearscreen() {
	fmt.Print("\033[H\033[2J")
}
