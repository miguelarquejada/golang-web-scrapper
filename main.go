package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("g1.globo.com"))

	c.OnHTML(".feed-post-body-title h2", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("Título: %s\n\n", title)
	})

	// Manipular erros
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Erro:", err)
	})

	// Visitar a página principal do G1
	err := c.Visit("https://g1.globo.com/")
	if err != nil {
		log.Fatalf("Erro ao visitar página: %v", err)
	}
}
