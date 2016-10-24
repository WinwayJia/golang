package html

import (
	"log"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

// ParseHTML parse html.
func ParseHTML(url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	selection := doc.Find("title")
	log.Println(selection.Text())

	class := doc.Find("h2").Text()
	fmt.Println(class, "\n")

	for _, v := range(doc.Find(".s_form").Nodes) {
		fmt.Println(v)
	}
	doc.Find(".s_form .s-p-top").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Find("image").Attr("src")
		if !exists {
			log.Println("index no exists", i)
			return
		}

		fmt.Println(i, href)
	})

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")


	fmt.Println("==================================================")
	doc, err = goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title a").Text()
		log.Println("第", i+1, "个帖子的标题：", title)
	})
	return nil
}

