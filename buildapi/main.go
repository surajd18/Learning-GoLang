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

// fake DB
var courses []Course

// middleware,helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == "";
	return c.CourseName == "";
}

func main() {
	fmt.Println("Hiiii");

	r := mux.NewRouter()

	//Seeding
	courses = append(courses, Course{CourseId: "2",CourseName: "ReactJs",CoursePrice: 299,Author: &Author{Fullname: "Suraj Divate",Website: "https://github.com/surajd18"}})
	courses = append(courses, Course{CourseId: "1",CourseName: "JavaScript",CoursePrice: 199,Author: &Author{Fullname: "ABC",Website: "https://github.com/surajd18"}})

	//Routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")

	//Listing to port
	log.Fatal(http.ListenAndServe(":4000",r))
}

//controller

func serveHome(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("<h1>Welcome Home</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all COurses");

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get one course");

	w.Header().Set("Content-Type","application/json");

	//Grab ID from Request
	params:= mux.Vars(r)

	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given ID")
	return;
}

func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create One Course..!!")

	w.Header().Set("Content-Type","application/json");

	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send Data");
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside the JSON")
		return;
	}

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses,course);
	json.NewEncoder(w).Encode(course)

	return;
}
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");

	params := mux.Vars(r);

	for index,course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"];
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return;
		}
	}
}

func deleteOneCourse (w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete Course")
	w.Header().Set("Content-Type","application/json");

	params := mux.Vars(r)

	for index,course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			break;
		}
	}
}