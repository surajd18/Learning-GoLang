package main

import "fmt"

// func main() {
// 	fmt.Println("variables");
// }

//Conversion

// func main(){
// 	fmt.Println("Please rate to our app!!");
// 	reader := bufio.NewReader(os.Stdin);
// 	input,_:=reader.ReadString('\n');
// 	fmt.Println("Thanks",input);

// 	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64);

// 	if err !=nil{
// 		fmt.Println(err);
// 	}else{
// 		fmt.Println("Added 1 to your ratting :",numRating+1);
// 	}
// }

//pointer

// func main() {
// 	myNumber := 23;

// 	var ptr = &myNumber;

// 	fmt.Println("Value of actual pointer is ",ptr); //Address of the memory
// 	fmt.Println("Value of actual pointer is ",*ptr); //value which is stored in that address

// 	*ptr = *ptr+2;

// 	fmt.Println("New value is: ",myNumber);   //25
// }

//Map

// func main() {
// 	languages := make(map[string]string);

// 	languages["JS"] = "JavaScript"
// 	languages["RB"] = "Ruby"
// 	languages["PY"] = "Python"

// 	fmt.Println("List of all Languanges: ",languages);
// 	fmt.Println("JS shorts for: ",languages["JS"]);

// 	delete(languages,"RB");
// 	fmt.Println("List of all Languanges: ",languages);

// 	//Loops are interesting in golang
// 	for key,value := range languages {
// 		fmt.Printf("For key %v, value is %v\n",key, value)
// 	}
// }

//Loops with different types

// func main(){
//     days := []string{"Sunday","Monday","Tuesday","Thursday","Friday","Saturday"}

//     // for d:=0;d< len(days);d++{
//     //     fmt.Println(days[d])
//     // }

//     // for i:= range days {
//     //     fmt.Println(days[i])
//     // }

//     for index, day := range days {
//         fmt.Printf("Index is %v and value is %v\n",index, day)
//     }
// }

//Defer in go
//In Go, defer is a keyword used to ensure that a function call is executed at the end of the surrounding function's execution, just before the function returns.
//Deferred calls are executed in LIFO (Last In, First Out) order. This means the most recent defer statement is executed first.

func main() {
	defer fmt.Println("World!")
	defer fmt.Println("one!")
	defer fmt.Println("two!")

	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
        defer fmt.Print(i);
	}
}

// output of this function is hello, 43210,two,one,world
