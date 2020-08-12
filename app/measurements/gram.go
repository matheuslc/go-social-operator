package measurements

// Gram type
type Gram struct {
	Amount int
}

// New creates a new Gram type literal.
// Rules to validate if a type Gram can exists has to be made here
func New(amount int) Gram {
	return Gram{Amount: amount}
}
