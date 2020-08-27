package main

import "fmt"
import "regexp"

var text string = "You know that you murdered people with a spoon, spoon can be dangerous since it made of stainless steel"

func main() {


  fmt.Print("Text:\n", text, "\n\n")
  fmt.Print("Edit Text above !!\n\n")

  var search, replace, newString string

  fmt.Print("Search the Word : ")
  fmt.Scanln(&search)
  fmt.Println()

  var regex, err = regexp.Compile(`(?i)`+search)
  // fmt.Println(regex)

  if err != nil {
    fmt.Println(err)
  }

  var isExist = regex.MatchString(text)
  // fmt.Println(isExist)
  if !isExist {
    fmt.Printf("\"%v\" not found in the text\n", search)
  } else {
    fmt.Print("Word for replacement: ")
    fmt.Scanln(&replace)
    newString = regex.ReplaceAllString(text, replace)
    fmt.Printf("\nNew Text:\n%s\n", newString)
  }
}
