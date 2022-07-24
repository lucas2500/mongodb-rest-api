package entities

type Hello struct {
	Ping string `json:"Ping"`
}

// Response structs
type ProductResponse struct {
	Code    string `json:"Code"`
	Success bool   `json:"Success"`
}

type Result struct {
	Products []ProductResponse `json:"Products"`
	Message  string            `json:"Message"`
}

// Request structs
type Product struct {
	Name    string `json:"Name"`
	Code    string `json:"Code"`
	Barcode string `json:"Barcode"`
	Active  bool   `json:"Active"`
}

type ReqBody struct {
	Products []Product `json:"Products"`
}
