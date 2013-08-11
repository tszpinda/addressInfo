package model

type Distance struct {
	Postcode1, Postcode2 string
	Meters               float64
}


type Address struct {
	HouseNumber string
	Street      string
	Town        string
	County      string
	Postcode    string
	Country     string
}

