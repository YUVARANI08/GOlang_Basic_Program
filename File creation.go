package main

import (
	"log"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := os.Create("sample.txt") //create file
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("hi welcome dois is my newfile") //write into file
	file.Close()
	
      stream,err:=ioutil.ReadFile("sample.txt") //reading from the file
      if err != nil {
		log.Fatal(err)
	}
      s1:=string(stream)
     fmt.Println(s1)