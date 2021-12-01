package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	hclog "github.com/hashicorp/go-hclog"
)

//ExchangeRates is a Struct for exchange rate
type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

//NewRates is a holder for new rates taken from api
func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: l, rates: map[string]float64{}}

	err := er.getRates()

	return er, err
}

//getRates gets exchange rate info from api
func (e *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed API call: got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	md := &Cubes{}
	xml.NewDecoder(resp.Body).Decode(&md)

	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}
		e.rates[c.Currency] = r
	}
	return nil
}

// Cubes holds many cube
type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

// Cube holds info from exchange rate api
type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
