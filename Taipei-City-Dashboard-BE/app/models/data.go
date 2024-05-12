// models/data.go

package models

// 假設您的數據結構如下
type Data struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// GetData fetches data from your data source
func GetData() ([]Data, error) {
	// 這裡是模擬數據，實際上您可能需要從數據庫中讀取數據
	mockData := []Data{
		{ID: 1, Value: "Example data 1"},
		{ID: 2, Value: "Example data 2"},
	}

	return mockData, nil
}
