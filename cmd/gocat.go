package main

import (
    "fmt"
    "os"
    "bufio"
)
//error handling if no args
//error handling if no file exists
func main() {

    if len(os.Args) <= 1 {
        fmt.Println("No arguments provided!")
        return
    }

    filename := os.Args[1]

    data, err := os.Open(filename)

    if err != nil{
        fmt.Println("No such file exists!")
        return
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
