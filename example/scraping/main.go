// https://zetcode.com/golang/goquery/
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	years := []int{2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024}
	for _, year := range years {
		url := fmt.Sprintf("https://www.sbineotrade.jp/ipo/record-%v.html", year)
		doc := fetchPage(url)
		header := getHeader(doc)
		body := getBody(doc)
		outputCsv(year, body, header)
	}
}

func fetchPage(webPage string) *goquery.Document {
	resp, err := http.Get(webPage)
	if err != nil {
		log.Printf("failed to get html: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("failed to load html: %s", err)
	}
	return doc
}

func getHeader(doc *goquery.Document) []string {
	table := doc.Find("div.table-responsive > table.text-right")
	thead := table.Find("thead > tr > th")
	var header []string
	thead.Each(func(i int, s *goquery.Selection) {
		header = append(header, s.Text())
	})
	return header
}

func getBody(doc *goquery.Document) [][]string {
	table := doc.Find("div.table-responsive > table.text-right")
	var body [][]string
	tbody := table.Find("tbody > tr")
	tbody.Each(func(i int, s *goquery.Selection) {
		var row []string
		s.Find(("td")).Each(func(j int, s *goquery.Selection) {
			row = append(row, s.Text())
		})
		body = append(body, row)
	})
	return body
}

// Output:
func outputCsv(year int, body [][]string, header []string) {
	file, err := os.Create(fmt.Sprintf("./file-%v.csv", year))
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	if err := w.Write(header); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	for _, b := range body {
		if err := w.Write(b); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
