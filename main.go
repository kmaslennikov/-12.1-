package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
  "strings"
)

func main() {
	file, err := os.Create("messages.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 1; ; i++ {
		var line string
		fmt.Println("Введите сообщение:")

		reader := bufio.NewReader(os.Stdin)
		line, err = reader.ReadString('\n')
    
		if err != nil {
			fmt.Println(err)
			return
		}
    line = strings.Trim(line, "\n ")
     if line == "exit" {
			return      
		}

		t := time.Now()
		timeToString := t.Format("2006-01-02 15:04:05")

		file.WriteString(fmt.Sprintf("%d. %v: %v\n", i, timeToString, line))
	}

}
