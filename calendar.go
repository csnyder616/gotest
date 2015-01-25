package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "errors"
    //"github.com/organsnyder/ical-go"
)

func main() {
    result, err := processCalendar("test.ics")
    if (err != nil) {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}

func processCalendar(filename string) (cal string, err error) {
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    scanner.Scan() // First line
    key, value := processLine(scanner.Text())
    if key != "BEGIN" || value != "VCALENDAR" {
        return "", errors.New("File doesn't begin with required “BEGIN: VCALENDAR” line")
    }
    for scanner.Scan() {
        line := scanner.Text()
        if (strings.HasPrefix(line, "BEGIN:")) {
            
        }
        if err := scanner.Err(); err != nil {
            return "", err
            // TODO handle error
        }
    }
    return "", nil
}

func processLine(line string) (key string, value string) {
    tokens := strings.SplitN(line, ":", 2)
    return strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1])
}