package main

import (
	"fmt"
	"go-testify-mock/src/modules"
)

func main() {
	scraping := &modules.Scraping{}
	scraping.Init()

	title, _ := scraping.GetPageTitle("https://qiita.com/")
	fmt.Sscanln(title)
}
