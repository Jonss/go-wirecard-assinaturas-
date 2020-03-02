package tests

import (
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

func TestCreatePlan(t *testing.T) {
	plan := &subs.Plan{
		Code:          "plan_code",
		Name:          "plan_name",
		Description:   "This is a plan test",
		Amount:        1096,
		PaymentMethod: subs.CREDIT_CARD,
		Status:        subs.Active,
	}

	planJSON, err := plan.CreatePlan()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(planJSON)
}

func TestGetPlans(t *testing.T) {
	plans, err := subs.ListPlans()
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range plans.Plans {
		fmt.Println(p.Code)
	}

}

func TestGetPlan(t *testing.T) {
	plan, err := subs.FindPlan("plan_code") // change it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(plan)
}

func TestActivatePlan(t *testing.T) {
	message, err := subs.ActivatePlan("plan_code") // change it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(message)
}

func TestInactivatePlan(t *testing.T) {
	message, err := subs.InactivatePlan("plan_code") // change it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(message)
}

func TestUpdatePlan(t *testing.T) {
	plan := &subs.Plan{
		Amount: 99999,
		Name:   "Plano topperson",
	}

	planJSON, err := plan.UpdatePlan("plan_code") // change it
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(planJSON)
}
