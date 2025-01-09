package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses,"","\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",finalJson);
}

func DecodeJson(){
	jsonDataFromweb := []byte(`
		{
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",
                "tags": ["web-dev","js"]
        }
	`)

	var lcocourse course;
	checkValid :=json.Valid(jsonDataFromweb);

	if checkValid {
		fmt.Println("Json was Valid");
		json.Unmarshal(jsonDataFromweb,&lcocourse);
		fmt.Printf("%#v\n",lcocourse);
	}else{
		fmt.Println("Json is not valid")
	}
}