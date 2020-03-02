package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Jonss/go-wirecard-assinaturas/subs"
	"github.com/Jonss/go-wirecard-assinaturas/subs/config"
)

func init() {
	config.WirecardConfig.Env = config.SANDBOX
	config.WirecardConfig.Token = ""
	config.WirecardConfig.Key = ""
}

func TestCreateSubscription(t *testing.T) {
	s := subs.Subscription{
		Code:          "subscription-wirecard-api-code-7",
		Amount:        10000,
		PaymentMethod: subs.CREDIT_CARD,
		ProRata:       true,
		BestInvoiceDate: subs.BestInvoiceDate{
			DayOfMonth: 20,
		},
		Customer: subs.Customer{
			Code: "customer-wirecard-api-code-2",
		},
		Plan: subs.Plan{
			Code: "code-1014",
		},
	}

	subscription, err := s.CreateSubscription()
	if err != nil {
		fmt.Println(err.Error())
	}

	subscriptionJSON, err := json.Marshal(subscription)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(subscriptionJSON))
}
