package subs

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
