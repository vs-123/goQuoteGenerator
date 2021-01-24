// Vahin Sharma
package main

import (
  "fmt"
  "os"
  "encoding/csv"
  "math/rand"
  "io"
  "log"
  "time"
)

func main() {
  var quotes, authors []string
  csvfile, err := os.Open("quotes.csv")
  if err != nil {
  	log.Fatalln("Couldn't open the csv file", err)
  }
  r := csv.NewReader(csvfile)
  r.Comma = ';'
  for {
  	record, err := r.Read()
  	if err == io.EOF {
  		break
  	}
  	if err != nil {
  		log.Fatal(err)
  	}
    quotes = append(quotes, record[0])
    authors = append(authors, record[1])
  }
  var chosenQuote int = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(quotes)-1)
  fmt.Println(quotes[chosenQuote], "-", authors[chosenQuote])
}
