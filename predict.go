package go_rubbles

type ProductsPredict struct {
	Products []Product `json:"products"`
}

type Product struct {
	ProductId  string            `json:"product_id"`
	Dates      ProductDates      `json:"dates"`
	Parameters ProductParameters `json:"parameters"`
	Results    *[]PredictResult  `json:"results,omitempty"`
}

type ProductParameters struct {
	Client      string `json:"client"`
	Type        string `json:"type"`
	Price       string `json:"price"`
	DiscountPpd string `json:"discount_ppd"`
	DiscountOff string `json:"discount_off"`
	DiscountOn  string `json:"discount_on"`
	ShelfPrice  string `json:"shelf_price"`
}

type ProductDates struct {
	ShipmentDateFrom string `json:"shipment_date_from"`
	ShipmentDateTo   string `json:"shipment_date_to"`
	ShelfDateFrom    string `json:"shelf_date_from"`
	ShelfDateTo      string `json:"shelf_date_to"`
}

type PredictResult struct {
	Predict   float64 `json:"predict"`
	TimeStamp int     `json:"time_stamp"`
}

type PredictResponse struct {
	Result PredictResults
	Id     string
}

type PredictResults struct {
	Output   ProductsPredict `json:"output"`
	Ready    bool            `json:"ready"`
	ErrorMsg string          `json:"error_msg"`
}
