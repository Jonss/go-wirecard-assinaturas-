package subs

import (
	"encoding/json"
	"fmt"

	"github.com/Jonss/go-wirecard-subs/subs/requests"
)

// Customer represents a customer on wirecard
type Customer struct {
	Code           string       `json:"code"`
	Email          string       `json:"email,omitempty"`
	FullName       string       `json:"fullname,omitempty"`
	Document       string       `json:"cpf,omitempty"`
	PhoneAreaCode  string       `json:"phone_area_code,omitempty"`
	PhoneNumber    string       `json:"phone_number,omitempty"`
	BirthdateDay   string       `json:"birthdate_day,omitempty"`
	BirthdateMonth string       `json:"birthdate_month,omitempty"`
	BirthdateYear  string       `json:"birthdate_year,omitempty"`
	Address        *Address     `json:"address,omitempty"`
	BillingInfo    *BillingInfo `json:"billing_info,omitempty"`
	CreationDate   string       `json:"creation_date,omitempty"`
	CreationTime   string       `json:"creation_time,omitempty"`
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

// CreateCustomer on wirecard
func (c Customer) CreateCustomer() (map[string]interface{}, error) {
	customer, _ := json.Marshal(c)
	resp, err := requests.Do(requests.POST, "/customers?new_vault=true", customer)
	if err != nil || resp.StatusCode > 299 {
		return nil, fmt.Errorf("An error occurred creating customer %s. StatusCode [%d]", c.Code, resp.StatusCode)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// FindCustomer by code
func FindCustomer(code string) (*Customer, error) {
	resp, err := requests.Do(requests.GET, "/customers/"+code, nil)
	if err != nil || resp.StatusCode > 299 {
		return nil, fmt.Errorf("An error occurred finding customer with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	var c *Customer
	json.NewDecoder(resp.Body).Decode(c)
	return c, nil
}
