package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SingleItem struct {
	Field    string  `json:"field"`
	Hour     int     `json:"hour"`
	Minute   int     `json:"minute"`
	ItemCode string  `json:"item_code"`
	Price    float64 `json:"price"`
	Quantity int     `json:"qty"`
}

type RawItems struct {
	Items            []SingleItem `json:"items"`
	TotalRecordCount int          `json:"total_record_count"`
	Start            int          `json:"start"`
}

type Calculator struct{}

func (c *Calculator) SomeComplexAggregationFunction(startDate, endDate time.Time, field string) (float64, error) {
	convertedStartTime := startDate.Format("2006-02-01")
	convertedEndTime := startDate.Format("2006-02-01")
	rawResp, err := http.Get(fmt.Sprintf("http://example-data-server/api/data-archive/v1/retail?field=%v&start-date=%v&end-date=%v", field, convertedStartTime, convertedEndTime))
	if err != nil {
		return 0.0, err
	}
	if rawResp.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("unexpected status code")
	}
	raw, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return 0.0, err
	}
	var items RawItems
	err = json.Unmarshal(raw, &items)
	if err != nil {
		return 0, err
	}
	// Pretend this is some complex calculation
	summer := 0.0
	for k, v := range items.Items {
		fmt.Printf("processing current item: %v", k)
		summer = float64(v.Quantity)*v.Price + summer
	}
	return summer, nil
}

func main() {}
