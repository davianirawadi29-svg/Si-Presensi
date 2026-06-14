// This is a simple program to track absence of students in college. The user is assumed to be a lecturer or an academic administration staff in a university.
// There will be 3 data managed. Student's data, class schedule data and attendance log data. 

package main

import "fmt"

const STUDENT 50
const COURSE 20

type studentDetail struct {
	name string
	sid  int
	class string
}

type studentList struct {
	detail [STUDENT]studentDetail
	count  int
}

// type StudentList will be used as student's data. Count determines the amount of students. 

type course struct {
	name       string
	present    studentList
	sick       studentList
	permission studentList
	alpha      studentList
}

type class struct {
	subjects [COURSE]course
	subjectcount int
	day string
}

type classlist [7]class

// Class Schedule Data. The type Class lists out the courses each day. The type Classlist lists out the class schedule.
// Its assumed every week has the same schedule.

type log struct {
	student  studentDetail
	presence int
	absence  int
}

type loglist struct {
	datalog [STUDENT]log
	logcount int
}
// type Loglist is used as attendance log data. Students listed as present in a course will have their presence count added.
// Students listed as sick, permission or alpha in a course will have their absence count added.		

func main() {
	var students studentList
	var schedule classlist
	var attendancelog loglist
	var choice, index int
	var sorted bool
	var sortchoice, ordersort string
	index = -1

	students.count = 0
	attendancelog.logcount = 0

	schedule[0].day = "Monday"
	schedule[0].subjectcount = 0
	schedule[1].day = "Tuesday"
	schedule[1].subjectcount = 0
	schedule[2].day = "Wednesday"
	schedule[2].subjectcount = 0
	schedule[3].day = "Thursday"
	schedule[3].subjectcount = 0
	schedule[4].day = "Friday"
	schedule[4].subjectcount = 0
	schedule[5].day = "Saturday"
	schedule[5].subjectcount = 0
	schedule[6].day = "Sunday"
	schedule[6].subjectcount = 0

// Inserting default data. Schedule day names can't be changed.
	
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

// The main menu of the program. Choice is the input for selection the options.
// Index is returned after searching for a student and can be used to change or delete student data at that index. Index is -1 at default.

	for choice != -1 {
		switch choice {
		case 1:
			addStudent(&students, &attendancelog, &sorted)
		case 2:
			changeStudentData(&students, &attendancelog, &schedule, &sorted, &index)
		case 3:
			deleteStudentData(&students, &attendancelog, &schedule, &sorted, &index)
		case 4:
			searchStudentData(students, &index, &sorted, &ordersort)
			if index != -1 {
				fmt.Printf("\n")
				fmt.Printf("Student found at index: %d\n", index)
				fmt.Print("Press Enter to go back...")
				fmt.Scanln()
			} else {
				fmt.Printf("\n")
				fmt.Print("Student not found.")
				fmt.Print("Press Enter to go back...")
				fmt.Scanln()
			}
		case 5:
			addClassSchedule(&schedule)
		case 6:
			viewClassSchedule(schedule)
		case 7:
			recordAttendance(&schedule, &attendancelog)
		case 8:
			viewAttendanceLog(attendancelog)
		case 9:
			fmt.Print("Sort based on name, absence or sid: ")
			fmt.Scan(&sortchoice)
			switch sortchoice {
			case "name":
				Insertion(&students, &attendancelog, &sorted, &ordersort)
			case "absence":
				Selection(&students, &attendancelog, &sorted, &ordersort)
			case "sid":
				Binary(&students, &attendancelog, &sorted, &ordersort)
			default:
				fmt.Printf("Invalid choice. Press enter to go back...")
				fmt.Scanln()
				fmt.Scanln()
			}
		default:
			fmt.Print("Invalid choice. Please select a valid option: ")
			fmt.Scan(&choice)
			continue
		}
		clearscreen()

	// Switch case is used so the code looks cleaner. If the choice inputted is invalid, the program asks for another input without showing the menu again.

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
		fmt.Scan(&choice)
	}

// After the user has done something with the program, the menu pops up again. If the user inputs -1, the program stops.

}

func clearscreen() {
	fmt.Print("\033[H\033[2J")
}

// Clears the terminal of any text after the user selects an option and after using one of the options.

func addStudent(students *studentList, attendancelog *loglist, sorted *bool) {
	var index int
	var choice string
	clearscreen()
	choice = "Yes"

	index = students.count
	for choice != "No" {
		if index < STUDENT {
			fmt.Print("Enter Student Name: ")
			fmt.Scan(&students.detail[index].name)
			fmt.Print("Enter Student ID: ")
			fmt.Scan(&students.detail[index].sid)
			fmt.Print("Enter Student Class: ")
			fmt.Scan(&students.detail[index].class)
			attendancelog.datalog[index].student.name = students.detail[index].name
			attendancelog.datalog[index].student.sid = students.detail[index].sid
			attendancelog.datalog[index].student.class = students.detail[index].class
			attendancelog.datalog[index].presence = 0
			attendancelog.datalog[index].absence = 0
	
			fmt.Print("Add another student? (Yes/No): ")
			fmt.Scan(&choice)
			index = index + 1
			fmt.Printf("\n")
		} else {
			fmt.Printf("Maximum capacity reached.\n")
			break
		}
	}
	attendancelog.logcount = index 
	students.count = index
	*sorted = false

	fmt.Println("Student(s) added successfully!")
	fmt.Print("Press Enter to go back...")
	fmt.Scanln()
	fmt.Scanln()
}

func changeStudentData(students *studentList, attendancelog *loglist, schedule *classlist, sorted *bool, index *int) {
	var j, k, l, day int
	var name string
	clearscreen()
	if *index == -1 {
		fmt.Print("No student selected. Please search for a student first.\n")
		fmt.Print("Press Enter to go back...")
		fmt.Scanln()
		return
	} else {
		fmt.Printf("Current Name: %s\n", students.detail[*index].name)
		fmt.Printf("Current SID: %d\n", students.detail[*index].sid)
		fmt.Printf("Current Class: %s\n", students.detail[*index].class)
		name = students.detail[*index].name
		fmt.Printf("\n")
		fmt.Print("Enter new name: ")
		fmt.Scan(&students.detail[*index].name)
		fmt.Print("Enter new SID: ")
		fmt.Scan(&students.detail[*index].sid)
		fmt.Print("Enter new Class: ")
		fmt.Scan(&students.detail[*index].class)
		attendancelog.datalog[*index].student.name = students.detail[*index].name
		attendancelog.datalog[*index].student.sid = students.detail[*index].sid
		attendancelog.datalog[*index].student.class = students.detail[*index].class
		attendancelog.datalog[*index].presence = 0
		attendancelog.datalog[*index].absence = 0

		day = 0
		for day < 7 {
			j = 0
			k = 0
			for j < schedule[day].subjectcount {
				for k < schedule[day].subjects[j].present.count {
					if schedule[day].subjects[j].present.detail[k].name == name {
						for l = k; l < schedule[day].subjects[j].present.count - 1; l++ {
							schedule[day].subjects[j].present.detail[l] = schedule[day].subjects[j].present.detail[l+1]
						}
						schedule[day].subjects[j].present.count--
					}
					k = k + 1
				}
				k = 0
				for k < schedule[day].subjects[j].sick.count {
					if schedule[day].subjects[j].sick.detail[k].name == name {
						for l = k; l < schedule[day].subjects[j].sick.count - 1; l++ {
							schedule[day].subjects[j].sick.detail[l] = schedule[day].subjects[j].sick.detail[l+1]
						}
						schedule[day].subjects[j].sick.count--
					}
					k = k + 1
				}
				k = 0
				for k < schedule[day].subjects[j].permission.count {
					if schedule[day].subjects[j].permission.detail[k].name == name {
						for l = k; l < schedule[day].subjects[j].permission.count - 1; l++ {
							schedule[day].subjects[j].permission.detail[l] = schedule[day].subjects[j].permission.detail[l+1]
						}
						schedule[day].subjects[j].permission.count--
					}
					k = k + 1
				}
				k = 0
				for k < schedule[day].subjects[j].alpha.count {
					if schedule[day].subjects[j].alpha.detail[k].name == name {
						for l = k; l < schedule[day].subjects[j].alpha.count - 1; l++ {
							schedule[day].subjects[j].alpha.detail[l] = schedule[day].subjects[j].alpha.detail[l+1]
						}
						schedule[day].subjects[j].alpha.count--
					}
					k = k + 1
				}
				k = 0
				j = j + 1
			}
			day = day + 1
		}
		
		*sorted = false
	}

	*index = -1
	fmt.Printf("\nStudent data updated, attendance log reset for this student.\n")
	fmt.Print("Press Enter to go back...")
	fmt.Scanln()
	fmt.Scanln()
}

func deleteStudentData(students *studentList, attendancelog *loglist, schedule *classlist, sorted *bool, index *int) {
	var j, k, l, day int
	var name string
	clearscreen()
	if *index == -1 {
		fmt.Print("No student selected. Please search for a student first.\n")
		fmt.Print("Press Enter to go back...")
		fmt.Scanln()
		return
	} else {
		name = students.detail[*index].name
		for i := *index; i < students.count - 1; i++ {
			students.detail[i] = students.detail[i+1]
			attendancelog.datalog[i] = attendancelog.datalog[i+1]
		}
		students.count = students.count - 1
		attendancelog.logcount = attendancelog.logcount - 1
		day = 0
		for day < 7 {
			j = 0
			k = 0
			for j < schedule[day].subjectcount {
				for k < schedule[day].subjects[j].present.count {
					if schedule[day].subjects[j].present.detail[k].name == name {
						for l = k; l < schedule[day].subjects[j].present.count - 1; l++ {
							schedule[day].subjects[j].present.detail[l] = schedule[day].subjects[j].present.detail[l+1]
						}
						schedule[day].subjects[j].present.count--
					}
					k = k + 1
				}
				k = 0
				for k < schedule[day].subjects[j].sick.count {
					if schedule[day].subjects[j].sick.detail[k].name == name {
						for l = k; l < schedule[day].subjects[j].sick.count - 1; l++ {
							schedule[day].subjects[j].sick.detail[l] = schedule[day].subjects[j].sick.detail[l+1]
						}
						schedule[day].subjects[j].sick.count--
					}
					k = k + 1
				}
				k = 0
				for k < schedule[day].subjects[j].permission.count {
					if schedule[day].subjects[j].permission.detail[k].name == name {
						for l = k; l < schedule[day].subjects[j].permission.count - 1; l++ {
							schedule[day].subjects[j].permission.detail[l] = schedule[day].subjects[j].permission.detail[l+1]
						}
						schedule[day].subjects[j].permission.count--
					}
					k = k + 1
				}
				k = 0
				for k < schedule[day].subjects[j].alpha.count {
					if schedule[day].subjects[j].alpha.detail[k].name == name {
						for l = k; l < schedule[day].subjects[j].alpha.count - 1; l++ {
							schedule[day].subjects[j].alpha.detail[l] = schedule[day].subjects[j].alpha.detail[l+1]
						}
						schedule[day].subjects[j].alpha.count--
					}
					k = k + 1
				}
				k = 0
				j = j + 1
			}
			day = day + 1
		}
		
		*sorted = false
	}
	*index = -1
	fmt.Printf("\nStudent data deleted successfully!\n")
	fmt.Print("Press Enter to go back...")
	fmt.Scanln()
	fmt.Scanln()
}

func searchStudentData(students studentList, index *int, sorted *bool, ordersort *string) {
	var choice, name string
	var sid, i, left, right int
	*index = -1
	clearscreen()
	fmt.Print("Search based on (name/sid): ")
	fmt.Scan(&choice)

	if choice == "name" {
		fmt.Print("Enter Student Name: ")
		fmt.Scan(&name)
		for i = 0; i < students.count && students.detail[i].name != name; i++ {
		}
	} else if choice == "sid" {
		if *sorted {
			fmt.Print("Enter Student ID: ")
			fmt.Scan(&sid)
			left = 0
			right = students.count - 1
			i = (left + right) / 2
			if *ordersort == "asc" {
				for left <= right {
					if students.detail[i].sid == sid {
						break
					} else if students.detail[i].sid > sid {
						right = i - 1
					} else {
						left = i + 1
					}
					i = (left + right) / 2
				}
			} else if *ordersort == "des" {
				for left <= right {
					if students.detail[i].sid == sid {
						break
					} else if students.detail[i].sid < sid {
						right = i - 1
					} else {
						left = i + 1
					}
					i = (left + right) / 2
				}
			}
		} else {
			fmt.Printf("Data is not sorted or empty. Can't search based on SID\n")
			return
		}
	} else {
		fmt.Printf("Invalid Search type! Searching failed!\n")
		return
	}
	if i < students.count {
		*index = i
	}
	fmt.Printf("\n")
}

func addClassSchedule(schedule *classlist) {
	var day, i int
	var choice string
	clearscreen()
	choice = ""
	day = 1
	i = schedule[day-1].subjectcount
	for day <= 7 && choice != "exit" {	
		fmt.Printf("Add class in day %d or exit (yes/no/exit)?\n", day)
		fmt.Scan(&choice)
		if choice == "yes" {
			for choice == "yes" {
				if i < COURSE {
					fmt.Print("Input the course name: ")
					fmt.Scan(&schedule[day-1].subjects[i].name)
					fmt.Printf("\n")
					fmt.Print("Add more course (yes/no)?\n")
					fmt.Scan(&choice)
					for choice != "no" && choice != "yes" {
						fmt.Printf("Invalid choice, Please select from yes/no.\n")
						fmt.Scan(&choice)
					}
					i = i + 1
				} else {
					fmt.Printf("Max course reached for this day\n")
					break
				}
			}
		} else if choice != "no" && choice != "exit" {
			fmt.Print("Invalid choice, please select between yes, no or exit (No capital letters!)\n")
			fmt.Scan(&choice)
			continue
		}
		schedule[day-1].subjectcount = i
		day = day + 1
		if day < 8 {
			i = schedule[day-1].subjectcount
		}
		clearscreen()
	}
	fmt.Printf("Press enter to go back...")
	fmt.Scanln()
	fmt.Scanln()
}

// Function to add class in schedule. It's assumed 2 subjects with the same name in the same day have different time.

func viewClassSchedule(schedule classlist) {
	var day, i int
	var choice string
	clearscreen()
	choice = ""
	day = 1
	i = 0

	for day <= 7 && choice != "exit" {
		fmt.Printf("View schedule in day %d or exit (yes/no/exit)?\n", day)
		fmt.Scan(&choice)
		if choice == "yes" {
			for i < schedule[day-1].subjectcount {
				fmt.Printf("Subject %d: %s\n", i+1, schedule[day-1].subjects[i].name)
				i = i + 1
			}
		} else if choice != "no" && choice != "exit" {
			fmt.Printf("Choice invalid! Please selection between yes, no or exit (No capital letters!)\n")
			fmt.Scan(&choice)
			continue
		}
		i = 0
		day = day + 1
	}

	fmt.Printf("Press enter to go back...")
	fmt.Scanln()
	fmt.Scanln()
}
// Fixing Record Attendance
func recordAttendance(schedule *classlist, attendancelog *loglist) {
	var subjectchoice, choice, attendance, inputchoice, name, optioncase string
	var i, day, j, k int
	i = 0
	day = 1
	k = 0
	inputchoice = "yes"
	fmt.Printf("Record Attendance in day %d? (yes/no/exit)\n", day)
	fmt.Scan(&choice)
	for choice != "exit" && day <= 7 {
		if choice == "yes" {
			for i < schedule[day-1].subjectcount {
				fmt.Printf("Add attendance for %s? (yes/no)\n", schedule[day-1].subjects[i].name)
				fmt.Scan(&subjectchoice)
				if subjectchoice == "yes" {
					fmt.Printf("Select from the options (present/sick/permission/alpha)\n")
					fmt.Scan(&attendance)
					switch attendance {
						case "present":
							inputchoice = "yes"
							j = schedule[day-1].subjects[i].present.count
							for inputchoice != "no" {
								if j < STUDENT {
									k = 0
									fmt.Print("Input student name present: ")
									fmt.Scan(&name)
									for k < attendancelog.logcount {
										if attendancelog.datalog[k].student.name == name {
											schedule[day-1].subjects[i].present.detail[j].name = attendancelog.datalog[k].student.name 
											schedule[day-1].subjects[i].present.detail[j].sid = attendancelog.datalog[k].student.sid
											schedule[day-1].subjects[i].present.detail[j].class = attendancelog.datalog[k].student.class
											schedule[day-1].subjects[i].present.count++
											attendancelog.datalog[k].presence = attendancelog.datalog[k].presence + 1
											break
										}
										k++
									}
									if k >= attendancelog.logcount {
										fmt.Printf("Student not found! Type exit if you don't know student name or anything if you want to continue.\n")
										fmt.Scan(&optioncase)
										if optioncase == "exit" {
											return
										} else {
											continue
										}
									}
								} else {
									fmt.Print("Maximum student capacity reached. Press enter to go back to menu..")
									fmt.Scanln()
									fmt.Scanln()
									return
								}
								fmt.Printf("\nAdd more students? (yes/no)\n")
								fmt.Scan(&inputchoice)
								j = j + 1
								k = 0
								for inputchoice != "no" && inputchoice != "yes" {
									fmt.Printf("Invalid option. Select yes or no\n")
									fmt.Scan(&inputchoice)
								}
							}
						case "sick":
							inputchoice = "yes"
							j = schedule[day-1].subjects[i].sick.count
							for inputchoice != "no" {
								if j < STUDENT {
									k = 0
									fmt.Print("Input student name that's sick: ")
									fmt.Scan(&name)
									for k < attendancelog.logcount {
										if attendancelog.datalog[k].student.name == name {
											schedule[day-1].subjects[i].sick.detail[j].name = attendancelog.datalog[k].student.name
											schedule[day-1].subjects[i].sick.detail[j].sid = attendancelog.datalog[k].student.sid
											schedule[day-1].subjects[i].sick.detail[j].class = attendancelog.datalog[k].student.class
											schedule[day-1].subjects[i].sick.count++
											attendancelog.datalog[k].absence = attendancelog.datalog[k].absence + 1
											break
										}
										k++
									}
									if k >= attendancelog.logcount {
										fmt.Printf("Student not found! Type exit if you don't know student name or anything if you want to continue.\n")
										fmt.Scan(&optioncase)
										if optioncase == "exit" {
											return
										} else {
											continue
										}
									}
								} else {
									fmt.Print("Maximum student capacity reached. Press enter to go back...")
									fmt.Scanln()
									fmt.Scanln()
									return
								}
								fmt.Printf("\nAdd more students? (yes/no)\n")
								fmt.Scan(&inputchoice)
								j = j + 1
								k = 0
								for inputchoice != "no" && inputchoice != "yes" {
									fmt.Printf("Invalid option. Select yes or no\n")
									fmt.Scan(&inputchoice)
								}
							}
						case "permission":
							inputchoice = "yes"
							j = schedule[day-1].subjects[i].permission.count
							for inputchoice != "no" {
								if j < STUDENT {
									k = 0
									fmt.Print("Input student name with permission: ")
									fmt.Scan(&name)
									for k < attendancelog.logcount {
										if attendancelog.datalog[k].student.name == name {
											schedule[day-1].subjects[i].permission.detail[j].name = attendancelog.datalog[k].student.name
											schedule[day-1].subjects[i].permission.detail[j].sid = attendancelog.datalog[k].student.sid
											schedule[day-1].subjects[i].permission.detail[j].class = attendancelog.datalog[k].student.class
											schedule[day-1].subjects[i].permission.count++
											attendancelog.datalog[k].absence = attendancelog.datalog[k].absence + 1
											break
										}
										k++
									}
									if k >= attendancelog.logcount {
										fmt.Printf("Student not found! Type exit if you don't know student name or anything if you want to continue.\n")
										fmt.Scan(&optioncase)
										if optioncase == "exit" {
											return
										} else {
											continue
										}
									}
								} else {
									fmt.Print("Maximum student capacity reached. Press enter to go back...")
									fmt.Scanln()
									fmt.Scanln()
									return
								}
								fmt.Printf("\nAdd more students? (yes/no)\n")
								fmt.Scan(&inputchoice)
								j = j + 1
								k = 0
								for inputchoice != "no" && inputchoice != "yes" {
									fmt.Printf("Invalid option. Select yes or no\n")
									fmt.Scan(&inputchoice)
								}
							} 
						case "alpha":
							inputchoice = "yes"
							j = schedule[day-1].subjects[i].alpha.count
							for inputchoice != "no" {
								if j < STUDENT {
									k = 0
									fmt.Print("Input student name that's alpha: ")
									fmt.Scan(&name)
									for k < attendancelog.logcount {
										if attendancelog.datalog[k].student.name == name {
											schedule[day-1].subjects[i].alpha.detail[j].name = attendancelog.datalog[k].student.name
											schedule[day-1].subjects[i].alpha.detail[j].sid = attendancelog.datalog[k].student.sid
											schedule[day-1].subjects[i].alpha.detail[j].class = attendancelog.datalog[k].student.class
											schedule[day-1].subjects[i].alpha.count++
											attendancelog.datalog[k].absence = attendancelog.datalog[k].absence + 1
											break
										}
										k++
									}
									if k >= attendancelog.logcount {
										fmt.Printf("Student not found! Type exit if you don't know student name or anything if you want to continue.\n")
										fmt.Scan(&optioncase)
										if optioncase == "exit" {
											return
										} else {
											continue
										}
									}
								} else {
									fmt.Print("Maximum student reached. Press enter to go back...")
									fmt.Scanln()
									fmt.Scanln()
									return
								}
								fmt.Printf("\nAdd more students? (yes/no)\n")
								fmt.Scan(&inputchoice)
								j = j + 1
								k = 0
								for inputchoice != "no" && inputchoice != "yes" {
									fmt.Printf("Invalid option. Select yes or no\n")
									fmt.Scan(&inputchoice)
								}
							}
						default:
							fmt.Printf("Invalid option!\n")
							continue
						}
				} else if subjectchoice != "no" {
					fmt.Printf("Invalid Choice!\n")
					continue
				}
				i = i + 1
			}
		} else if choice != "no" {
			fmt.Printf("Invalid choice! Please select between yes/no/exit\n")
			fmt.Scan(&choice)
			continue
		}
		if day < 7 {
			day = day + 1
			fmt.Printf("Record Attendance in day %d? (yes/no/exit)\n", day)
			fmt.Scan(&choice)
			i = 0
			k = 0
			inputchoice = "yes"
		} else {
			break
		}
	}
	fmt.Print("Press enter to go back...")
	fmt.Scanln()
	fmt.Scanln()
}

func viewAttendanceLog(attendancelog loglist) {
	var i int
	clearscreen()
	if attendancelog.logcount > 0 {
		for i = 0; i < attendancelog.logcount; i++ {
			fmt.Printf("Student Name: %s\n", attendancelog.datalog[i].student.name)
			fmt.Printf("Presence: %d\n", attendancelog.datalog[i].presence)
			fmt.Printf("Absence: %d\n", attendancelog.datalog[i].absence)
			fmt.Println()
		}
		fmt.Printf("\nData Shown, Please press enter to go back...")
		fmt.Scanln()
		fmt.Scanln()
	} else {
		fmt.Print("No attendance log available. Press enter to go back...")
		fmt.Scanln()
		fmt.Scanln()
	}
}

func Insertion(students *studentList, attendancelog *loglist, sorted *bool, ordersort *string) {
	var temp1 studentDetail
	var temp2 log
	var i, j int
	var choice string
	clearscreen()
	if students.count > 0 {
		fmt.Printf("Sort ascendingly or descendingly? (asc/des)\n")
		fmt.Scan(&choice)
		if choice == "asc" {
			for i = 1; i < students.count; i++ {
				temp1 = students.detail[i]
				temp2 = attendancelog.datalog[i]
				j = i - 1
				for j >= 0 && students.detail[j].name > temp1.name {
					students.detail[j+1] = students.detail[j]
					attendancelog.datalog[j+1] = attendancelog.datalog[j]
					j = j - 1
				}
				students.detail[j+1] = temp1
				attendancelog.datalog[j+1] = temp2
			}
			*sorted = false
			*ordersort = ""
			fmt.Printf("Data successfully sorted ascendingly!\n")
		} else if choice == "des" {
			for i = 1; i < students.count; i++ {
				temp1 = students.detail[i]
				temp2 = attendancelog.datalog[i]
				j = i - 1
				for j >= 0 && students.detail[j].name < temp1.name {
					students.detail[j+1] = students.detail[j]
					attendancelog.datalog[j+1] = attendancelog.datalog[j]
					j = j - 1
				}
				students.detail[j+1] = temp1
				attendancelog.datalog[j+1] = temp2
			}
			*sorted = false
			*ordersort = ""
			fmt.Printf("Data successfully sorted descendingly!\n")
		} else {
			fmt.Printf("Invalid choice!\n")
		}
	} else {
		fmt.Printf("No students available.\n")
	}
	fmt.Print("Press enter to go back...")
	fmt.Scanln()
	fmt.Scanln()
}

func Selection(students *studentList, attendancelog *loglist, sorted *bool, ordersort *string) {
	var index, i, j int
	var temp1 studentDetail
	var temp2 log
	var choice string
	clearscreen()
	if students.count > 0 {
		fmt.Printf("Sort ascendingly or descendingly? (asc/des)\n")
		fmt.Scan(&choice)
		if choice == "asc" {
			for i = 0; i < students.count - 1; i++ {
				index = i
				for j = i+1; j < students.count; j++ {
					if attendancelog.datalog[j].absence < attendancelog.datalog[index].absence {
						index = j
					}
				}
				if index != i {
					temp1 = students.detail[i]
					temp2 = attendancelog.datalog[i]
					students.detail[i] = students.detail[index]
					attendancelog.datalog[i] = attendancelog.datalog[index]
					students.detail[index] = temp1
					attendancelog.datalog[index] = temp2
				}
			}
			*sorted = false
			*ordersort = ""
			fmt.Printf("Data successfully sorted ascendingly!\n")
		} else if choice == "des" {
			for i = 0; i < students.count - 1; i++ {
				index = i
				for j = i+1; j < students.count; j++ {
					if attendancelog.datalog[j].absence > attendancelog.datalog[index].absence {
						index = j
					}
				}
				if index != i {
					temp1 = students.detail[i]
					temp2 = attendancelog.datalog[i]
					students.detail[i] = students.detail[index]
					attendancelog.datalog[i] = attendancelog.datalog[index]
					students.detail[index] = temp1
					attendancelog.datalog[index] = temp2
				}
			}
			*sorted = false
			*ordersort = ""
			fmt.Printf("Data successfully sorted descendingly!\n")
		} else {
			fmt.Printf("Invalid choice!\n")
		}
	} else {
		fmt.Printf("No students available.\n")
	}
	fmt.Print("Press enter to go back...")
	fmt.Scanln()
	fmt.Scanln()
}

func Binary(students *studentList, attendancelog *loglist, sorted *bool, ordersort *string) {
	var temp1 studentDetail
	var temp2 log
	var i, j int
	var choice string
	clearscreen()
	if students.count > 0 {
		fmt.Printf("Sort ascendingly or descendingly? (asc/des)\n")
		fmt.Scan(&choice)
		if choice == "asc" {
			for i = 1; i < students.count; i++ {
				temp1 = students.detail[i]
				temp2 = attendancelog.datalog[i]
				j = i - 1
				for j >= 0 && students.detail[j].sid > temp1.sid {
					students.detail[j+1] = students.detail[j]
					attendancelog.datalog[j+1] = attendancelog.datalog[j]
					j = j - 1
				}
				students.detail[j+1] = temp1
				attendancelog.datalog[j+1] = temp2
			}
			*sorted = true
			*ordersort = "asc"
			fmt.Printf("Data successfully sorted ascendingly!\n")
		} else if choice == "des" {
			for i = 1; i < students.count; i++ {
				temp1 = students.detail[i]
				temp2 = attendancelog.datalog[i]
				j = i - 1
				for j >= 0 && students.detail[j].sid < temp1.sid {
					students.detail[j+1] = students.detail[j]
					attendancelog.datalog[j+1] = attendancelog.datalog[j]
					j = j - 1
				}
				students.detail[j+1] = temp1
				attendancelog.datalog[j+1] = temp2
			}
			*sorted = true
			*ordersort = "des"
			fmt.Printf("Data successfully sorted descendingly!\n")
		} else {
			fmt.Printf("Invalid choice!\n")
		}
	} else {
		fmt.Printf("No students available.\n")
	}
	fmt.Print("Press enter to go back...")
	fmt.Scanln()
	fmt.Scanln()
}
			
	
	
				
	
		
			
		
	


	
	
