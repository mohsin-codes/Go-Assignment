package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)


// func healthHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/health" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprint(w, "Content Not Found : ", http.StatusNoContent)
// }

//////////////////////////////////////////////////Manipulation Methods///////////////////////////////////////////////////
func concat(str string) string {
	var emptySlice []string
	for i := 0; i < len(str); i++ {
		if int(str[i]) >= 65 && int(str[i]) <= 90 || int(str[i]) >= 97 && int(str[i]) <= 122 {
			emptySlice = append(emptySlice, string(str[i]))
		}
	}
	retString := strings.Join(emptySlice, "")
	return strings.ToLower(retString)
}

func dup_count(list []string) map[string]int {
	duplicate_frequency := make(map[string]int)
	for _, item := range list {
		_, exist := duplicate_frequency[item]
		if exist {
			duplicate_frequency[item] += 1
		} else {
			duplicate_frequency[item] = 1
		}
	}
	return duplicate_frequency
}

func vowelCount(str string) map[string]int {
	var LowerCase []string
	var UpperCase []string
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "a" || string(str[i]) == "e" || string(str[i]) == "i" || string(str[i]) == "o" || string(str[i]) == "u" {
			LowerCase = append(LowerCase, string(str[i]))
		} else if string(str[i]) == "A" || string(str[i]) == "E" || string(str[i]) == "I" || string(str[i]) == "O" || string(str[i]) == "U" {
			UpperCase = append(UpperCase, string(str[i]))
		}
	}
	for i := 0; i < len(UpperCase); i++ {
		LowerCase = append(LowerCase, UpperCase[i])
	}
	FinalMap := dup_count(LowerCase)
	return FinalMap
}

func convertToBytes(str string, b *[]byte) {
	byteSlice := []byte(str)
	*b = byteSlice
	fmt.Println(b)
}

func asciiTotal(str string) int {
	Total := 0
	for i := 0; i < len(str); i++ {
		Total += int(str[i])
	}
	return Total
}

func manipulationHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() error = %v", err)
		return
	}
	// fmt.Fprintf(w, "POST request Succesful")
	string := r.FormValue("string")
	radioVal := r.FormValue("group1")
	stringRes := ""
	var Total int
	b := []byte{0}
	if radioVal == "concat" {
		stringRes = concat(string)
		defer fmt.Fprintf(w, "Concatenated String : %v", stringRes)
	} else if radioVal == "vowelCount" {
		finalMap := vowelCount(string)
		defer fmt.Fprintf(w, "Vowel Count : %v", finalMap)
	} else if radioVal == "asciiTotal" {
		Total = asciiTotal(string)
		defer fmt.Fprintf(w, "ASCII Total : %v", Total)
	} else if radioVal == "convertToBytes" {
		convertToBytes(string, &b)
		fmt.Println(b)
		defer fmt.Fprintf(w, "Converted to Bytes: %v", b)
	}
	fmt.Fprintf(w, "Inputted String: %s\n", string)
}

///////////////////////////////////////////////////////Area Methods//////////////////////////////////////////////////////////
func areaRect(s string) float64 {
	dim := strings.Split(s, ",")
	i1, err := strconv.ParseFloat(dim[0], 32)
	if err != nil {
		fmt.Println(err)
	}
	i2, err := strconv.ParseFloat(dim[1], 32)
	if err != nil {
		fmt.Println(err)
	}
	area := i1 * i2
	return area
}

func areaCircle(s string) float64 {
	pi := 3.14159265
	dim := strings.Split(s, ",")
	i1, err := strconv.ParseFloat(dim[0], 32)
	if err != nil {
		fmt.Println(err)
	}
	var area float64 = i1 * pi
	return area
}

func areaSquare(s string) float64 {
	dim := strings.Split(s, ",")
	i1, err := strconv.ParseFloat(dim[0], 32)
	if err != nil {
		fmt.Println(err)
	}
	var area float64 = i1 * i1
	return area
}

func areaAll(s string) (float64, float64, float64) {
	pi := 3.14159265
	dim := strings.Split(s, ",")
	i1, err := strconv.ParseFloat(dim[0], 32)
	if err != nil {
		fmt.Println(err)
	}
	i2, err := strconv.ParseFloat(dim[1], 32)
	if err != nil {
		fmt.Println(err)
	}

	i3, err := strconv.ParseFloat(dim[2], 32)
	if err != nil {
		fmt.Println(err)
	}

	i4, err := strconv.ParseFloat(dim[3], 32)
	if err != nil {
		fmt.Println(err)
	}

	area1 := i1 * i2
	area2 := i3 * pi
	area3 := i4 * i4
	return area1, area2, area3
}

func areaHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() error = %v", err)
		return
	}
	// fmt.Fprintf(w, "POST request Succesful")
	string := r.FormValue("string")
	radioVal := r.FormValue("group1")
	var area float64
	var area1, area2, area3 float64
	if radioVal == "rect" {
		s := strings.Split(string, ",")
		if len(s) != 2 {
			fmt.Fprintf(w, "Wrong Number of Inputs! Try Again! : Error Code %v", http.StatusNotAcceptable)
		} else {
			area = areaRect(string)
			fmt.Fprintf(w, "Area of Rectangle : %v", area)
		}
	} else if radioVal == "circle" {
		s := strings.Split(string, ",")
		if len(s) != 1 {
			fmt.Fprintf(w, "Wrong Number of Inputs! Try Again! : Error Code %v", http.StatusNotAcceptable)
		} else {
			area = areaCircle(string)
			fmt.Fprintf(w, "Area of Cirlce : %v", area)
		}
	} else if radioVal == "square" {
		s := strings.Split(string, ",")
		if len(s) != 1 {
			fmt.Fprintf(w, "Wrong Number of Inputs! Try Again! : Error Code %v", http.StatusNotAcceptable)
		} else {
			area = areaSquare(string)
			fmt.Fprintf(w, "Area of Cirlce : %v", area)
		}
	} else if radioVal == "all" {
		s := strings.Split(string, ",")
		if len(s) != 4 {
			fmt.Fprintf(w, "Wrong Number of Inputs! Try Again! : Error Code %v", http.StatusNotAcceptable)
		} else {
			area1, area2, area3 = areaAll(string)
			fmt.Fprintf(w, "Area of Rectangle : %v\n", area1)
			fmt.Fprintf(w, "Area of Circle : %v\n", area2)
			fmt.Fprintf(w, "Area of Square : %v", area3)
		}
	}

}

func main() {
	c := make(chan string)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/health", SiteCheck("localhost:8080/health",c))
	http.HandleFunc("/string/manipulation", manipulationHandler)
	http.HandleFunc("/calculate/area", areaHandler)

	for {
		SiteCheck("localhost:8080/health",c)
	}
	_, err := http.Get("http://localhost:8080")
	for link := range c{
		if err != nil{
			go func(l string){
				time.Sleep(5 * time.Second)
				SiteCheck(l,c)
			}(link)
		}
	}
		

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func SiteCheck(link string, c chan string) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		if r.URL.Path != "/health" {
			http.Error(w, "404 not found", http.StatusNotFound)
			fmt.Println("404 not found!")
			c <- link
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method is not supported", http.StatusNotFound)
			fmt.Println("404 not found!")
			c <-link 
			return
		}
		fmt.Fprint(w, "Content Not Found : ", http.StatusNoContent)
		c <- link
	}
}