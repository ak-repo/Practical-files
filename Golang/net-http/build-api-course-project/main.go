package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course-file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// author model
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

//middleware/helper

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

//controller

// home
func HandleHome(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("<h1> WELCOME TO HOMEPAGE</h1>"))

}

// get all courses
func handleAllCourses(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("get all course")
	writer.Header().Set("Content-Type", "application/json")

	//passing json value
	json.NewEncoder(writer).Encode(courses)

}

// get one course
func handleGetOneCourse(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(request)
	fmt.Println(params)

	//find matching id course
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(writer).Encode(course)
			return
		}
	}
	// not found
	writer.WriteHeader(http.StatusNotFound)
	json.NewEncoder(writer).Encode("No course found given id")

}

// create course
func handleCreateCourse(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	//check body is empty
	if request.Body == nil {
		writer.WriteHeader(http.StatusNotFound)

		json.NewEncoder(writer).Encode("Please sent some data")
		return
	}

	// what if its {}
	var course Course
	_ = json.NewDecoder(request.Body).Decode(&course)
	if course.IsEmpty() {
		writer.WriteHeader(http.StatusNotFound)

		json.NewEncoder(writer).Encode("Please sent some data")
		return
	}

	// generate unique id and conv into string

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	//append course into courses
	courses = append(courses, course)
	json.NewEncoder(writer).Encode(course)

}

// update course info
func handleUpdateCourse(writer http.ResponseWriter, request *http.Request) {

	//grab id from request
	params := mux.Vars(request)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			//new data
			var course Course
			json.NewDecoder(request.Body).Decode(&course)
			course.CourseId = params["id"]

			//removing old and adding new
			// courses = append(courses[:index], append([]Course{course}, courses[index+1:]...)...)  one line code
			courses = append(courses[:index], courses[index+1:]...)
			courses = append(courses, course)
			//response
			json.NewEncoder(writer).Encode(course)
			return

		}
	}

	// if no data found
	json.NewEncoder(writer).Encode("Course updation failed")

}

// delete course
func handleDeleteCourse(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			//response
			json.NewEncoder(writer).Encode("course deleted successfully")
			return

		}

	}

	//not found
	writer.WriteHeader(http.StatusNotFound)

	json.NewEncoder(writer).Encode("something went wrong! course not found")
}

// main function
func main() {

	//router
	router := mux.NewRouter()

	//dummy data
	courses = append(courses, Course{CourseId: "1", CoursePrice: 200, CourseName: "react", Author: &Author{Fullname: "ak", Website: "ak.com"}})
	courses = append(courses, Course{CourseId: "2", CoursePrice: 300, CourseName: "python", Author: &Author{Fullname: "ak", Website: "ak.com"}})
	courses = append(courses, Course{CourseId: "3", CoursePrice: 700, CourseName: "golang", Author: &Author{Fullname: "ak", Website: "ak.com"}})

	//routes
	router.HandleFunc("/", HandleHome).Methods("GET")
	router.HandleFunc("/courses", handleAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", handleGetOneCourse).Methods("GET")
	router.HandleFunc("/course", handleCreateCourse).Methods("POST")
	router.HandleFunc("/course/{id}", handleUpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", handleDeleteCourse).Methods("DELETE")

	//server
	log.Fatal(http.ListenAndServe(":8080", router))

}
