package subs

import (
	"encoding/json"
	"fmt"

	"github.com/Jonss/go-wirecard-assinaturas/subs/requests"
)

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
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
	Status        status        `json:"status,omitempty"`
}

// Plans contains a slice of plan
type Plans struct {
	Plans []Plan `json:"plans"`
}

// CreatePlan a plan in Wirecard
func (p Plan) CreatePlan() (map[string]interface{}, error) {
	plan, _ := json.Marshal(p)
	resp, err := requests.Do(requests.POST, "/plans", plan)
	if err != nil || resp.StatusCode > 299 {
		return nil, fmt.Errorf("An error occurred on creating plan %s. StatusCode [%d]", p.Code, resp.StatusCode)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// ListPlans all plans of account
func ListPlans() (Plans, error) {
	resp, err := requests.Do(requests.GET, "/plans", nil)
	if err != nil || resp.StatusCode > 299 {
		return Plans{}, fmt.Errorf("An error occurred listing plans. StatusCode [%d]", resp.StatusCode)
	}

	var plans Plans
	json.NewDecoder(resp.Body).Decode(&plans)
	return plans, nil
}

// FindPlan plan by code
func FindPlan(code string) (Plan, error) {
	resp, err := requests.Do(requests.GET, "/plans/"+code, nil)
	if err != nil || resp.StatusCode > 299 {
		return Plan{}, fmt.Errorf("An error occurred finding plan with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	var p Plan
	json.NewDecoder(resp.Body).Decode(&p)
	return p, nil
}

// ActivatePlan plan by code
func ActivatePlan(code string) (string, error) {
	resp, err := requests.Do(requests.PUT, "/plans/"+code+"/activate", nil)
	if err != nil || resp.StatusCode > 299 {
		return "Erro na ativação do plano", fmt.Errorf("An error occurred activating plan with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	return "Plano ativado com sucesso", nil
}

// InactivatePlan by code
func InactivatePlan(code string) (string, error) {
	resp, err := requests.Do(requests.PUT, "/plans/"+code+"/inactivate", nil)
	if err != nil || resp.StatusCode > 299 {
		return "Erro na inativação do plano", fmt.Errorf("An error occurred inactivating plan with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	return "Plano desativado com sucesso", nil
}

// UpdatePlan by code
func (p Plan) UpdatePlan(code string) (string, error) {
	plan, _ := json.Marshal(p)

	resp, err := requests.Do(requests.PUT, "/plans/"+code, plan)
	if err != nil || resp.StatusCode > 299 {
		return "Plano não atualizado.", fmt.Errorf("An error occurred updating plan with code %s. StatusCode [%d]", code, resp.StatusCode)
	}

	return "Plano atualizado com sucesso", nil
}
