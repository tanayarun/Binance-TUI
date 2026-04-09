package binance

import "testing"

func TestSign(t *testing.T) {
	result := sign("timestamp=123456789", "mysecret")
	expected := "fefb2d43c59b7aa0444091d9b00a80b71991465d116ab32c3c398202606f5685"
	if result != expected {
		t.Errorf("expected %s but got %s", expected, result)
	}
}

func TestSignWrongSecret(t *testing.T) {
	result := sign("timestamp=123456789", "wrongsecret")
	expected := "fefb2d43c59b7aa0444091d9b00a80b71991465d116ab32c3c398202606f5685"
	if result == expected {
		t.Error("signatures should not match with different secret keys", expected, result)
	}
}
