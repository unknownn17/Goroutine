package main

import (
    "fmt"
    "log"
    "os"
)
func file1(ch1 chan string) {
    f, err := os.OpenFile("file1.txt", os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    _, err = f.WriteString("Hello")
    if err != nil {
        log.Fatal(err)
    }
    dat, err := os.ReadFile("file1.txt")
    if err != nil {
        log.Fatal(err)
    }
    ch1 <- string(dat)
}

func file2(ch2 chan string) {
    f, err := os.Create("file2.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close() 
	_, err = f.WriteString("World")
    if err != nil {
        log.Fatal(err)
    }
    dat, err := os.ReadFile("file2.txt")
    if err != nil {
        log.Fatal(err)
    }
    ch2 <- string(dat)
}

func main() {
    ch1 := make(chan string,1)
    ch2 := make(chan string,1)
	res := make(chan string,3)

    go file1(ch1)
    go file2(ch2)
	res <- <-ch1 +" "+ <-ch2

fmt.Println(<-res)
}
