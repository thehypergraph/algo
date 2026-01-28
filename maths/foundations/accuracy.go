package foundations

import "fmt"

type Accuracy struct {
	LowerBound float64
	UpperBound float64
}

func (a Accuracy) Print() {
	fmt.Printf("%f <=x < %f", a.LowerBound, a.UpperBound)
}

// This is to practice the Degrees of Accuracy concept I learned
// on MathAcademy recently.
func DegreesOfAccuracy(target float64, rng float64) Accuracy {

	measure := rng / 2
	lowerBound := target - measure
	upperBound := target + measure

	accuracy := Accuracy{
		LowerBound: lowerBound,
		UpperBound: upperBound,
	}

	return accuracy
}
