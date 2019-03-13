package main

import(
  "fmt"
  "flag"  
)

func main(){
	url := flag.String("url", "", "Youtube Url")
	flag.Parse()
	
	if *url == ""{
	  fmt.Println("Please flag url should not be empty.")
	  return
	}
	fmt.Printf("User input is: %s \n", *url)
}
