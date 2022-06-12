package modules

import (
	"github.com/sclevine/agouti"
)

type Scraping struct {
	// スクレイピングするページのタイトルの取得
	GetPageTitle func(url string) (string, error)
}

/*
初期化処理
*/
func (scraping *Scraping) Init() {
	// エラー処理は、めんどくさいので飛ばします。
	scraping.GetPageTitle = func(url string) (string, error) {
		options := agouti.ChromeOptions(
			"args", []string{
				"--headless",
				"--disable-gpu",
			})

		driver := agouti.ChromeDriver(options)
		defer driver.Stop()
		driver.Start()

		page, _ := driver.NewPage()
		page.Navigate(url)

		return page.Title()
	}
}
