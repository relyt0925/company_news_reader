package myallies

import (
	"testing"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func TestNewsItems(t *testing.T) {
	fileBytes, e := ioutil.ReadFile("./latestNewsResponse.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		panic(e)
	}
	fmt.Printf("%s\n", string(fileBytes))
	fmt.Println()

    	var newsItems NewsItems
    	json.Unmarshal(fileBytes, &newsItems)
	if len(newsItems) != 50{
		t.Error("DIDN'T READ ALL ELEMENTS IN FILE")
	}
}

func TestFetchCompanyNews(t *testing.T) {
	resp, err := FetchCompanyNews("MSFT")
	if err != nil{
		t.Error(err)
	}
	fmt.Println((*resp)[49])
}

func TestFetchStockQuote(t *testing.T) {
	resp, err := FetchStockQuote("MSFT")
	if err != nil{
		t.Error(err)
	}
	fmt.Println(*resp)
}

func TestFetchNewsItemContent(t *testing.T) {
	resp, err := FetchNewsItemContent(366183)
	if err != nil{
		t.Error(err)
	}
	fmt.Println(*resp)
}