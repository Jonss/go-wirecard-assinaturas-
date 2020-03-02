package customer

import (
	"encoding/json"
	"fmt"

	"github.com/Jonss/go-wirecard-assinaturas/config"
	"github.com/Jonss/go-wirecard-assinaturas/requests"
)

func init() {
	config.WirecardConfig.Env = config.SANDBOX
	config.WirecardConfig.Token = ""
	config.WirecardConfig.Key = ""
}

// Customer represents a customer on wirecard
type Customer struct {
	Code           string      `json:"code"`
	Email          string      `json:"email,omitempty"`
	FullName       string      `json:"fullname,omitempty"`
	Document       string      `json:"cpf,omitempty"`
	PhoneAreaCode  string      `json:"phone_area_code,omitempty"`
	PhoneNumber    string      `json:"phone_number,omitempty"`
	BirthdateDay   string      `json:"birthdate_day,omitempty"`
	BirthdateMonth string      `json:"birthdate_month,omitempty"`
	BirthdateYear  string      `json:"birthdate_year,omitempty"`
	Address        Address     `json:"address,omitempty"`
	BillingInfo    BillingInfo `json:"billing_info,omitempty"`
	CreationDate   string      `json:"creation_date,omitempty"`
	CreationTime   string      `json:"creation_time,omitempty"`
}

// Address represents a address on wirecard
type Address struct {
	Street     string `json:"street,omitempty"`
	Number     string `json:"number,omitempty"`
	Complement string `json:"complement,omitempty"`
	District   string `json:"district,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	Zipcode    string `json:"zipcode,omitempty"`
}

// BillingInfo represents payment info of customer
type BillingInfo struct {
	CreditCard  CreditCard   `json:"credit_card,omitempty"`
	CreditCards []CreditCard `json:"credit_cards,omitempty"`
}

// CreditCard represents a credit card
type CreditCard struct {
	HolderName      string `json:"holder_name,omitempty"`
	Number          string `json:"number,omitempty"`
	ExpirationMonth string `json:"expiration_month,omitempty"`
	ExpirationYear  string `json:"expiration_year,omitempty"`
	Vault           string `json:"vault,omitempty"`
	FirstSixDigits  string `json:"first_six_digits,omitempty"`
	LastFourDigits  string `json:"last_four_digits,omitempty"`
	Brand           string `json:"brand,omitempty"`
}

// Create a customer on wirecard
func (c Customer) Create() (map[string]interface{}, error) {
	customer, _ := json.Marshal(c)
	resp, err := requests.Do(requests.POST, "/customers?new_vault=true", customer)
	if err != nil || resp.StatusCode > 299 {
		return nil, fmt.Errorf("An error occurred creating customer %s. StatusCode [%d]", c.Code, resp.StatusCode)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// Find customer by code
func Find(code string) (Customer, error) {
	resp, err := requests.Do(requests.GET, "/customers/"+code, nil)
	if err != nil || resp.StatusCode > 299 {
		return Customer{}, fmt.Errorf("An error occurred finding customer with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	var c Customer
	json.NewDecoder(resp.Body).Decode(&c)
	return c, nil
}
