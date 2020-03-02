package subs

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Jonss/go-wirecard-assinaturas/subs/requests"
)

// Subscription represents a subscription on Wirecard
type Subscription struct {
	Code            string          `json:"code,"`
	Amount          int             `json:"amount"`
	Plan            Plan            `json:"plan"`
	Customer        Customer        `json:"customer"`
	PaymentMethod   PaymentMethod   `json:"payment_method"`
	ProRata         bool            `json:"pro_rata,omitempty"`
	BestInvoiceDate BestInvoiceDate `json:"best_invoice_date,omitempty"`
}

// BestInvoiceDate is used on proRata operation
type BestInvoiceDate struct {
	DayOfMonth  int `json:"day_of_month,omitempty"`
	MonthOfYear int `json:"month_of_year,omitempty"`
}

// Response represents a subscription response on Wirecard
type Response struct {
	Code            string          `json:"code"`
	Amount          int             `json:"amount"`
	Message         string          `json:"message"`
	Errors          []string        `json:"errors"`
	Plan            Plan            `json:"plan"`
	Status          string          `json:"status"`
	Customer        Customer        `json:"customer"`
	Invoice         Invoice         `json:"invoice"`
	Alerts          []string        `json:"alerts"`
	NextInvoiceDate NextInvoiceDate `json:"next_invoice_date"`
}

// Invoice represents an invoice on wirecard
type Invoice struct {
	Amount int           `json:"amount"`
	ID     int           `json:"id"`
	Status InvoiceStatus `json:"status"`
}

// InvoiceStatus represents status of an invoice
type InvoiceStatus struct {
	Description string `json:"description"`
	Code        int    `json:"code"`
}

// NextInvoiceDate represents the next date to charge an invoice
type NextInvoiceDate struct {
	Month int `json:"month"`
	Year  int `json:"year"`
	Day   int `json:"day"`
}

// CreateSubscription creates a subscription on wirecard
func (s Subscription) CreateSubscription() (*Response, error) {
	if s.PaymentMethod == ALL {
		return nil, fmt.Errorf("PaymentMethod must be %s or %s", string(CREDIT_CARD), string(BOLETO))
	}

	newCustomer := strconv.FormatBool(s.Customer.Email != "")

	code := s.Code
	subscription, _ := json.Marshal(s)
	resp, err := requests.Do(requests.POST, "/subscriptions?new_customer="+newCustomer, subscription)
	if err != nil || resp.StatusCode > 299 {
		return nil, fmt.Errorf("An error occurred creating subscription with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	var subsResponse *Response
	json.NewDecoder(resp.Body).Decode(subsResponse)

	return subsResponse, nil
}
