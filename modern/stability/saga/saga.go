package saga

import (
	"errors"
	"fmt"
)

type Step struct {
	Name         string
	Action       func() error
	Compensation func() error
}

type Saga struct {
	Sequence []*Step
}

func (s *Saga) Execute() {
	for i, step := range s.Sequence {
		err := step.Action()
		if err != nil {
			fmt.Println("Error occurred in the step : ", step.Name)

			// we have to rollback everything
			for j := i; j >= 0; j-- {
				fmt.Println("Compensation executed for step : ", s.Sequence[j].Name)
				_ = s.Sequence[j].Compensation()
			}

			return
		}

		fmt.Println("Step executed successfully : ", step.Name)
	}
}

func AddOrder() error {
	fmt.Println("Order created")
	return nil
}

func CancelOrder() error {
	fmt.Println("Order cancelled")
	return nil
}

func DoEarmark() error {
	fmt.Println("Earmarked tokens")
	return nil
}

func ReleaseEarmark() error {
	fmt.Println("Released tokens")
	return nil
}

func PerformMatching() error {
	fmt.Println("Error occurred while attempting matching ")
	return errors.New("Error occurred while attempting matching ")
}

func DoRollback() error {
	fmt.Println("Perform rollback for matching ")
	return nil
}
