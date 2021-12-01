package model

type ApiRequest struct {
	Canal    string       `json:"Canal"`
	Datos    DatosRequest `json:"Datos"`
	SubCanal string       `json:"SubCanal"`
}

type DatosRequest struct {
	Assets    []string `json:"Assets"`
	Currency  string   `json:"Currency"`
	EndDate   string   `json:"EndDate"`
	StartDate string   `json:"StartDate"`
	Use       string   `json:"Use"`
}

type ApiResponse struct {
	Datos          []Price `json:"Datos"`
	Codigo         int     `json:"Codigo"`
	Mensaje        string  `json:"Mensaje"`
	MensajeTecnico string  `json:"MensajeTecnico"`
}

type Price struct {
	Price     string `json:"Price"`
	Asset     string `json:"Asset"`
	Currency  string `json:"Currency"`
	Use       string `json:"Use"`
	PriceDate string `json:"PriceDate"`
}
