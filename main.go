package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
)

func Pop(stack []string) []string {
	return stack[:len(stack)-1]
}

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	org, err := os.Open("./original.css")
	if err != nil {
		panic(err)
	}
	defer org.Close()

	scanner := bufio.NewScanner(org)

	lineNum := 1
	var attributeStack []string

	for scanner.Scan() {
		line := scanner.Text()
		lineHash := Hash(line)

		fmt.Println(attributeStack)

		if len(attributeStack) == 0 {
			selectorText := strings.Split(line, " ")
			for _, tag := range selectorText {
				fmt.Println(tag)
			}
		}
		if len(line) > 0 && line[len(line)-1:] == "{" {
			attributeStack = append(attributeStack, strconv.Itoa(lineNum)+"{")
		} else if line == "}" {
			attributeStack = Pop(attributeStack)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
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
