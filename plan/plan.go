package plan

import (
	"encoding/json"
	"fmt"

	"github.com/Jonss/go-wirecard-assinaturas/requests"
)

type paymentMethod string
type unit string
type status string

const (
	Active   status = "ACTIVE"
	Inactive status = "INACTIVE"
)

const (
	Day   unit = "DAY"
	Month unit = "MONTH"
	Year  unit = "YEAR"
)

const (
	CreditCard paymentMethod = "CREDIT_CARD"
	Boleto     paymentMethod = "BOLETO"
)

// Interval entity
type Interval struct {
	Length int  `json:"interval"`
	Unit   unit `json:"unit"`
}

// Trial entity
type Trial struct {
	Days         int  `json:"days"`
	Enabled      bool `json:"enabled"`
	HoldSetupFee bool `json:"hold_setup_fee"`
}

// Plan is the entity related with WirecardPlan
type Plan struct {
	Code          string        `json:"code,omitempty"`
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	Amount        int           `json:"amount,omitempty"`
	SetupFee      int           `json:"setup_fee,omitempty"`
	Interval      *Interval     `json:"interval,omitempty"`
	BillingCycles *int          `json:"billing_cycles,omitempty"`
	Trial         *Trial        `json:"trial,omitempty"`
	PaymentMethod paymentMethod `json:"payment_method,omitempty"`
	Status        status        `json:"status,omitempty"`
}

// Plans contains a slice of plan
type Plans struct {
	Plans []Plan `json:"plans"`
}

// Create a plan in Wirecard
func (p Plan) Create() (map[string]interface{}, error) {
	plan, _ := json.Marshal(p)
	resp, err := requests.Do(requests.POST, "/plans", plan)
	if err != nil || resp.StatusCode > 299 {
		return nil, fmt.Errorf("An error occurred on creating plan %s. StatusCode [%d]", p.Code, resp.StatusCode)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// List all plans of account
func List() (Plans, error) {
	resp, err := requests.Do(requests.GET, "/plans", nil)
	if err != nil || resp.StatusCode > 299 {
		return Plans{}, fmt.Errorf("An error occurred listing plans. StatusCode [%d]", resp.StatusCode)
	}

	var plans Plans
	json.NewDecoder(resp.Body).Decode(&plans)
	return plans, nil
}

// Find plan by code
func Find(code string) (Plan, error) {
	resp, err := requests.Do(requests.GET, "/plans/"+code, nil)
	if err != nil || resp.StatusCode > 299 {
		return Plan{}, fmt.Errorf("An error occurred finding plan with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	var p Plan
	json.NewDecoder(resp.Body).Decode(&p)
	return p, nil
}

// Activate plan by code
func Activate(code string) (string, error) {
	resp, err := requests.Do(requests.PUT, "/plans/"+code+"/activate", nil)
	if err != nil || resp.StatusCode > 299 {
		return "Erro na ativação do plano", fmt.Errorf("An error occurred activating plan with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	return "Plano ativado com sucesso", nil
}

// Inactivate plan by code
func Inactivate(code string) (string, error) {
	resp, err := requests.Do(requests.PUT, "/plans/"+code+"/inactivate", nil)
	if err != nil || resp.StatusCode > 299 {
		return "Erro na inativação do plano", fmt.Errorf("An error occurred inactivating plan with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	return "Plano desativado com sucesso", nil
}

// Update plan
func (p Plan) Update(code string) (string, error) {
	plan, _ := json.Marshal(p)

	resp, err := requests.Do(requests.PUT, "/plans/"+code, plan)
	if err != nil || resp.StatusCode > 299 {
		return "Plano não atualizado.", fmt.Errorf("An error occurred updating plan with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	return "Plano atualizado com sucesso", nil
}
