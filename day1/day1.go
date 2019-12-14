package main

import (

    "fmt"
    "bufio"
    "log"
    "os"
    "strconv"
)

func main() {
    fmt.Println("Day 1 assignment")

    file, err := os.Open("./input.txt")
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }
    var sum int
    sum = 0
    s := bufio.NewScanner(file)
    for s.Scan() {
        line := s.Text()
        n, err := strconv.Atoi(line)
        if err != nil {
            log.Fatal(err)
        }
	sum = sum + ((n/3)-2)
        fmt.Printf(" > Read %d \n", n)
    }
    fmt.Printf("Total sum: %d\n", sum)
}
