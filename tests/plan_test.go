package tests

import (
	"fmt"
	"testing"

	"github.com/Jonss/go-wirecard-assinaturas/config"
	"github.com/Jonss/go-wirecard-assinaturas/plan"
)

func init() {
	config.WirecardConfig.Env = config.SANDBOX
	config.WirecardConfig.Token = ""
	config.WirecardConfig.Key = ""
}

func TestCreatePlan(t *testing.T) {
	plan := &plan.Plan{
		Code:          "plan_code",
		Name:          "plan_name",
		Description:   "This is a plan test",
		Amount:        1096,
		PaymentMethod: plan.CreditCard,
		Status:        plan.Active,
	}

	planJSON, err := plan.Create()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(planJSON)
}

func TestGetPlans(t *testing.T) {
	plans, err := plan.List()
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range plans.Plans {
		fmt.Println(p.Code)
		fmt.Println("-----------------------")
	}

}

func TestGetPlan(t *testing.T) {
	plan, err := plan.Find("plan_code") // change it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(plan)
}

func TestActivatePlan(t *testing.T) {
	message, err := plan.Activate("plan_code") // change it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(message)
}

func TestInactivatePlan(t *testing.T) {
	message, err := plan.Inactivate("plan_code") // change it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(message)
}

func TestUpdatePlan(t *testing.T) {
	plan := &plan.Plan{
		Amount: 99999,
		Name:   "Plano topperson",
	}

	planJSON, err := plan.Update("plan_code") // change it
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(planJSON)
}
