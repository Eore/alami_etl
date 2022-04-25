package processor

import "fmt"

type Data struct {
	ID               int
	Name             string
	Age              int
	Balanced         float64
	PreviousBalanced float64
	AverageBalanced  float64
	FreeTransfer     int
}

type Datas []*Data

func (r *Data) CalculateAverage() float64 {
	avg := (r.Balanced + r.PreviousBalanced) / 2
	r.AverageBalanced = avg
	return avg
}

type Benefit struct {
	FreeTransfer int
	Balanced     float64
}

func (r *Data) CalculateBenefit() *Benefit {
	freeTransferCount := 5
	balancedBonus := 25

	switch {
	case r.Balanced >= 100 && r.Balanced <= 150:
		r.FreeTransfer = freeTransferCount
		return &Benefit{FreeTransfer: freeTransferCount}
	case r.Balanced > 150:
		r.Balanced += float64(balancedBonus)
		return &Benefit{Balanced: float64(balancedBonus)}
	}

	return nil
}

func (r *Data) AddBalance(amount float64) *Data {
	r.Balanced += amount
	return r
}

func (r Datas) BonusDisburst(budget float64, countPeople int) {
	bonusPerPerson := budget / float64(countPeople)

	for i, data := range r {
		if i > countPeople {
			break
		}
		data.Balanced += bonusPerPerson
	}
}

func (r Data) ToRowCSVString(workerID int) string {
	return fmt.Sprintf("%d;%s;%0.0f;%d;%d;%0.0f;%.2f;%d;%d\n", r.ID, r.Name, r.Balanced, workerID, workerID, r.PreviousBalanced, r.AverageBalanced, workerID, r.FreeTransfer)
}
