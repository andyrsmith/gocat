package main

import (
    "fmt"
    "os"
    "bufio"
)
//error handling if no args
//error handling if no file exists
func main() {
    filename := os.Args[1]
    data, err := os.Open(filename)

    if err != nil{
        panic(err)
    }
    
    fileScanner := bufio.NewScanner(data)

    fileScanner.Split(bufio.ScanLines)

    lineNumber := 1

    for fileScanner.Scan() {
        fmt.Printf("%d: ", lineNumber)
        fmt.Println(fileScanner.Text())
        lineNumber += 1
    }

    data.Close()

}
