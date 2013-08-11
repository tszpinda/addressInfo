package address

import (
	"code.google.com/p/gorest"
	"fmt"
	"strings"
	address "github.com/tszpinda/addressInfo/address/google"
	model "github.com/tszpinda/addressInfo/model"
)

type AddressService struct {
	gorest.RestService `root:"/ds/" consumes:"application/json" produces:"application/json"`
	getAddress         gorest.EndPoint `method:"GET" path:"/address/{postcode:string}" output:"Address"`
}

func (serv AddressService) GetAddress(postcode string) model.Address {
	fmt.Println("incoming GetAddress request: ", postcode)

	a, err := address.GetAddress(postcode)
	if err != nil {
		if strings.ContainsAny(err.Error(), "Unable to find address") {
			serv.ResponseBuilder().SetResponseCode(404).WriteAndOveride([]byte("Postcode '" + postcode + "' not found"))
		}else{
			serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte(err.Error()))
		}
	}
	return a
}
