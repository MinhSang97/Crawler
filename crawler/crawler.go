// crawler.go
package crawler

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly/v2"
)

// Crawler is now modified to return a string slice
func Crawler() ([]string, []string, []string, []string, []string, error) {

	// crawler := make(chan bool, 5)

	// numGoroutine := 5

	// Bắt đầu việc xử lý trang web khi tìm thấy các thẻ p
	var pTag []string
	var aTag []string
	var collectedTitles []string
	var collectedURLs []string
	var collectedHtml []string

	// Tạo một bộ lọc mới của Colly
	c := colly.NewCollector()

	// Create an OnHTML handler for the "p" tag
	c.OnHTML("p", func(e *colly.HTMLElement) {
		p_tag := e.Text
		pTag = append(pTag, p_tag)
		fmt.Println("Paragraph 'p' tag:", p_tag)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		a_tag := e.Text
		aTag = append(aTag, a_tag)
		fmt.Println("Paragraph 'a' tag:", a_tag)
	})

	// Crawl title
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		collectedTitles = append(collectedTitles, title) // Thêm title vào slice
		fmt.Println("Title:", title)
	})

	// Đăng ký hàm callback cho thẻ <a> có thuộc tính href
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Lấy đường dẫn URL từ thuộc tính href
		link := e.Attr("href")

		// Kiểm tra xem đường dẫn URL thuộc cùng domain không
		if isSameDomain(link, "https://dbeaver.io/") {
			// Thêm đường dẫn URL vào slice
			collectedURLs = append(collectedURLs, link)

			// Ghé thăm trang mới
			e.Request.Visit(link)

			// Hiển thị đường dẫn URL
			fmt.Println("url:", link)
		}
	})

	// Xử lý lỗi
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Crawl full HTML
	c.OnResponse(func(r *colly.Response) {
		html := string(r.Body)
		collectedHtml = append(collectedHtml, html)
		fmt.Println("Full HTML:", html)
	})

	// Truy cập trang web ban đầu
	err := c.Visit("https://dbeaver.io/")
	if err != nil {
		fmt.Println("Error visiting website:", err)
		return nil, nil, nil, nil, nil, err
	}

	// Chạy bộ lọc
	c.Wait()

	// Return the collected paragraphs
	return pTag, aTag, collectedURLs, collectedTitles, collectedHtml, nil

}

// Hàm kiểm tra xem một đường dẫn URL có thuộc cùng domain không
func isSameDomain(link, baseDomain string) bool {
	u, err := url.Parse(link) // Replace urlpkg with net/url here
	if err != nil {
		return false
	}

	base, err := url.Parse(baseDomain) // Replace urlpkg with net/url here
	if err != nil {
		return false
	}

	return u.Hostname() == base.Hostname()
}
