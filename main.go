package main

import (
  "os"
  "fmt"
)

func main(){
  org,err := os.ReadFile("./original.css")
  if err != nil {
    panic(err)
  }
  fmt.Print(string(org))
}

/**
* Steps, read in the file. Interprate the text until and score until { is 
* reached. Record the index of the line that is currently being read and 
* associate this with the score. Create a hash from 0 to the index of { and map this also to the line number
* 
* Run a sorting algorithm. 
*
* If next entry has the same hash combine the two. To do this we collect all of the styles
* and we hashify each line. If any of he lines are the same we only print one, if not we print styles from both.
* Once sorted create a new file and begin writing a copy of the newly sorted list. It moves to the next line number after the stack of {} is empty.
*
*/
