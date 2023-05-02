package Models

type SymbolInterval struct {
	Symbol   string `json:"symbol" xml:"symbol" form:"symbol"`
	Interval string `json:"interval" xml:"interval" form:"interval"`
}
