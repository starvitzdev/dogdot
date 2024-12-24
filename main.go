package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	url := "https://www.scrapethissite.com/pages/simple/"
	countries := []string{}
	newCountries := []string{}

	fmt.Println("starting main function")

	c := colly.NewCollector()

	c.OnHTML(".country-name", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		countries = append(countries, e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	fmt.Println(newCountries)
}
