package ptr

import "testing"

func TestOrElse(t *testing.T) {
	if OrElse((*string)(nil), "a") != "a" {
		t.Errorf("OrElse() mismatch")
	}

	if OrElse(ToPtr("b"), "a") != "b" {
		t.Errorf("OrElse() mismatch")
	}
}

func TestOrElseFunc(t *testing.T) {
	if OrElseF((*string)(nil), func() string { return "a" }) != "a" {
		t.Errorf("OrElseF() mismatch")
	}

	if OrElseF(ToPtr("b"), func() string { return "a" }) != "b" {
		t.Errorf("OrElseF() mismatch")
	}

	// Verify that the function is not called if the value is not nil
	called := false
	if OrElseF(ToPtr("b"), func() string {
		called = true
		return "a"
	}) != "b" {
		t.Errorf("OrElseF() mismatch")
	}

	if called {
		t.Errorf("OrElseF() function was called")
	}
}

func TestToPtr(t *testing.T) {
	if *ToPtr("a") != "a" {
		t.Errorf("ToPtr() mismatch")
	}

	// check that the pointer is of right type (i.e. this test compiles)
	//goland:noinspection ALL
	var s *string = ToPtr("a")
	if *s != "a" {
		t.Errorf("ToPtr() mismatch")
	}
}
