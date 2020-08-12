package measurements

import "testing"

func TestGram(t *testing.T) {
	gram := New(10)

	if gram.Amount != 10 {
		t.Errorf("New gram was created with incorrect value. Expect %d, get %d", 10, gram.Amount)
	}
}
