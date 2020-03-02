package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Jonss/go-wirecard-assinaturas/config"
	"github.com/Jonss/go-wirecard-assinaturas/customer"
)

func init() {
	config.WirecardConfig.Env = config.SANDBOX
	config.WirecardConfig.Token = ""
	config.WirecardConfig.Key = ""
}

func TestCreateCustomer(t *testing.T) {
	c := customer.Customer{
		Code:           "customer-wirecard-api-code-2",
		Email:          "myemail@myemail.com",
		Document:       "43363273002",
		FullName:       "My Fullname",
		PhoneAreaCode:  "11",
		PhoneNumber:    "999999998",
		BirthdateDay:   "02",
		BirthdateMonth: "05",
		BirthdateYear:  "1988",
		Address: customer.Address{
			Street:   "Av. Brigadeiro Faria Lima",
			Number:   "3064",
			District: "Itaim Bibi",
			City:     "São Paulo",
			State:    "SP",
			Country:  "BRA",
			Zipcode:  "01451001",
		},
		BillingInfo: customer.BillingInfo{
			CreditCard: customer.CreditCard{
				HolderName:      "My Fullname",
				Number:          "5555666677778884",
				ExpirationMonth: "12",
				ExpirationYear:  "20",
			},
		},
	}

	fmt.Println(c)

	m, err := c.Create()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(m)

}

func TestFindCustomer(t *testing.T) {
	c, err := customer.Find("customer-wirecard-api-code-2")
	if err != nil {
		fmt.Println(err.Error())
	}

	customerJSON, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(customerJSON))
}
