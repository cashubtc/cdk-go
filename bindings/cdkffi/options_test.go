package cdk_ffi

import "testing"

func TestNewSendOptionsIsValid(t *testing.T) {
	options := NewSendOptions(true)
	if err := options.Validate(); err != nil {
		t.Fatalf("NewSendOptions returned invalid options: %v", err)
	}
	if !options.IncludeFee {
		t.Fatal("IncludeFee was not preserved")
	}
}

func TestNewReceiveOptionsIsValid(t *testing.T) {
	if err := NewReceiveOptions().Validate(); err != nil {
		t.Fatalf("NewReceiveOptions returned invalid options: %v", err)
	}
}

func TestZeroValueOptionsAreRejected(t *testing.T) {
	if err := (&SendOptions{}).Validate(); err == nil {
		t.Fatal("expected zero-value SendOptions to be rejected")
	}
	if err := (&ReceiveOptions{}).Validate(); err == nil {
		t.Fatal("expected zero-value ReceiveOptions to be rejected")
	}
}
