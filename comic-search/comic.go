package main

// Comic struct represents a single comic entry
type Comic struct {
	Month      int    `json:"month,string"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       int    `json:"year,string"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        int    `json:"day,string"`
}
