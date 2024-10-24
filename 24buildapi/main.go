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

//model for course -file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB

var courses []Course

//middleware, helper -file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - Kdalearn.ng")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Chemistry", CoursePrice: 299, Author: &Author{Fullname: "Kingdavid Oshin", Website: "kaydeedev.ng"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Physics", CoursePrice: 350, Author: &Author{Fullname: "Jane Doe", Website: "janescihub.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Mathematics", CoursePrice: 399, Author: &Author{Fullname: "John Smith", Website: "mathmaster.io"}})
	courses = append(courses, Course{CourseId: "5", CourseName: "Biology", CoursePrice: 279, Author: &Author{Fullname: "Sarah Lee", Website: "bioexplore.org"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourses).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by David</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No Course found for ID provided")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what if {}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if course.IsEmpty() {
		http.Error(w, "Please send all required data", http.StatusBadRequest)
		return
	}

	// check for duplicate names
	for _, existingCourse := range courses {
		if course.CourseName == existingCourse.CourseName {
			http.Error(w, "Course name already exists", http.StatusBadRequest)
			return // Stop further execution if duplicate is found
		}
	}

	// generate a unique id, convert them to string
	// append course into courses

	// rand.Seed(time.Now().UnixNano()) //deprecated
	// course.CourseId = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100))
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(randGen.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from request

	params := mux.Vars(r)

	// loop through the values, find id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode((&course))
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: send a resonse when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove(index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// send a confirm or deny

			json.NewEncoder(w).Encode(fmt.Sprintf("Course with id - %s has been deleted", params["id"]))
			return
			break
		}
	}
}
