package model

type WebPage struct {
	URL      string `gorm:"column:url"`
	HTML     string `gorm:"column:html"`
	Title    string `gorm:"column:title"`
	P_Tag    string `gorm:"column:p_tag"`
	A_Tag    string `gorm:"column:a_tag"`
	Category string `gorm:"column:category"`
}

func (w WebPage) TableName() string {
	return "webpage"
}

// func NewWebPage(url string, html string, title string, pTag string, aTag string, category string) WebPage {
// 	return WebPage{
// 		URL:      url,
// 		HTML:     html,
// 		Title:    title,
// 		P_Tag:    pTag,
// 		A_Tag:    aTag,
// 		Category: category,
// 	}
// }
