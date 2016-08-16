package myallies

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
)

const prefix = "http://myallies.com/api"



type NewsItems []struct {
	NewsID int `json:"NewsID"`
	Company struct {
		ID int `json:"ID"`
		Symbol string `json:"Symbol"`
		Name string `json:"Name"`
		ISIN string `json:"ISIN"`
		CIK int `json:"CIK"`
		TradeCount int `json:"TradeCount"`
		ViewCount int `json:"ViewCount"`
		Stock interface{} `json:"Stock"`
		Stream interface{} `json:"Stream"`
		LogoPath string `json:"LogoPath"`
	} `json:"Company"`
	Symbol string `json:"Symbol"`
	Title string `json:"Title"`
	Content interface{} `json:"Content"`
	Type int `json:"Type"`
	Created string `json:"Created"`
	URL string `json:"URL"`
	GeneratedURL string `json:"GeneratedURL"`
	Duration string `json:"Duration"`
	Comments []interface{} `json:"Comments"`
	Likes []interface{} `json:"Likes"`
	LikesCount int `json:"LikesCount"`
	DislikesCount int `json:"DislikesCount"`
	CommentsCount int `json:"CommentsCount"`
}


func (newsItems *NewsItems) Length() int {
	bytes, _ := json.Marshal(newsItems)
	return len(bytes)
}

func (newsItems *NewsItems) Encode() ([]byte, error) {
	bytes, _ := json.Marshal(newsItems)
	return bytes, nil
}



type StockQuote struct {
	StockID int `json:"StockID"`
	LastTradePriceOnly string `json:"LastTradePriceOnly"`
	ChangePercent string `json:"ChangePercent"`
	CompanyName string `json:"CompanyName"`
}

func (sq *StockQuote) Length() int {
	bytes, _ := json.Marshal(sq)
	return len(bytes)
}

func (sq *StockQuote) Encode() ([]byte, error) {
	bytes, _ := json.Marshal(sq)
	return bytes, nil
}

type NewsItemContent struct {
	ID int `json:"ID"`
	Title string `json:"Title"`
	Content string `json:"Content"`
	PublishDate string `json:"PublishDate"`
}



var httpClient = &http.Client{Timeout:time.Second*10, }


func FetchCompanyNews(companyName string) ( *NewsItems, error){
	endpoint := fmt.Sprintf("%s/news/%s",prefix,companyName)
	response, err := httpClient.Get(endpoint)
	if err != nil{
		return nil, err
	}
	newsItems := new(NewsItems)
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(newsItems); err != nil{
		return nil, err
	}
	return newsItems, nil
}

func FetchStockQuote(companyName string) ( *StockQuote, error){
	endpoint := fmt.Sprintf("%s/quote/%s",prefix,companyName)
	response, err := httpClient.Get(endpoint)
	if err != nil{
		return nil, err
	}
	newsItems := new(StockQuote)
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(newsItems); err != nil{
		return nil, err
	}
	return newsItems, nil
}


func FetchNewsItemContent (newsID int) (*NewsItemContent, error){
	endpoint := fmt.Sprintf("%s/newsitem/%d",prefix,newsID)
	fmt.Println(endpoint)
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Accept","application/json")
	response, err := httpClient.Do(req)
	if err != nil{
		return nil, err
	}
	nc := new(NewsItemContent)
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(nc); err != nil{
		return nil, err
	}
	return nc, nil
}

