package main

import(
  "fmt"
  "flag"  
)

func main(){
	url := flag.String("url", "", "Youtube Url")
	flag.Parse()
	fmt.Printf("User input is: %s \n", *url)
}
