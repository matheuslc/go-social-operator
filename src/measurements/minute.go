package measurements

import (
	"fmt"
	"time"
)

// Gram defines the gram unit type
// e.g 1.0
type Minute int

func (minute Minute) ValueOf() string {
	if m, err := time.ParseDuration(fmt.Sprintf("%d", minute)); err != nil {
		return m.String()
	}

	return ""
}
