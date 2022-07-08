package chef

import (
	"testing"
)

func TestNewChefProfessional(t *testing.T) {
	n := Name("carmel")
	r := Role("professional")

	c, err := NewChef(r, n, "example@example.org")

	if err != nil {
		t.Errorf("Cannot create a valid chef with valid fields. Some rule changed? Error: %s", err)
	}

	if c.Name != n {
		t.Errorf("Chef was created with wrong name. Expected: %s. Get: %s", n, c.Name)
	}

	if c.RoleType != r {
		t.Errorf("Chef was created with wrong role. Expected: %s. Get: %s", r, c.RoleType)
	}
}

func TestNewChefAmauter(t *testing.T) {
	n := Name("carmel")
	r := Role("amauter")

	c, err := NewChef(r, n, "example@example.org")

	if err != nil {
		t.Errorf("Cannot create a valid chef with valid fields. Some rule changed? Error: %s", err)
	}

	if c.Name != n {
		t.Errorf("Chef was created with wrong name. Expected: %s. Get: %s", n, c.Name)
	}

	if c.RoleType != r {
		t.Errorf("Chef was created with wrong role. Expected: %s. Get: %s", r, c.RoleType)
	}
}
