package cdk_ffi

import "fmt"

// NewSendOptions returns a SendOptions with the enum fields initialized to the
// defaults Rust expects. Constructing SendOptions directly with a zero value
// panics inside the uniffi converters before any wallet logic runs.
func NewSendOptions(includeFee bool) *SendOptions {
	return &SendOptions{
		AmountSplitTarget:       SplitTargetNone{},
		SendKind:                SendKindOnlineExact{},
		IncludeFee:              includeFee,
		P2pkLockedProofSendMode: P2pkLockedProofSendModeSwap,
	}
}

// NewReceiveOptions returns a ReceiveOptions with its required enum field set.
func NewReceiveOptions() *ReceiveOptions {
	return &ReceiveOptions{
		AmountSplitTarget: SplitTargetNone{},
	}
}

// Validate reports unset enum fields before they reach the FFI boundary.
func (o *SendOptions) Validate() error {
	if o == nil {
		return fmt.Errorf("send options are nil")
	}
	if o.AmountSplitTarget == nil {
		return fmt.Errorf("send options: AmountSplitTarget must be set (use NewSendOptions)")
	}
	if o.SendKind == nil {
		return fmt.Errorf("send options: SendKind must be set (use NewSendOptions)")
	}
	if o.P2pkLockedProofSendMode == 0 {
		return fmt.Errorf("send options: P2pkLockedProofSendMode must be set (use NewSendOptions)")
	}
	return nil
}

// Validate reports unset enum fields before they reach the FFI boundary.
func (o *ReceiveOptions) Validate() error {
	if o == nil {
		return fmt.Errorf("receive options are nil")
	}
	if o.AmountSplitTarget == nil {
		return fmt.Errorf("receive options: AmountSplitTarget must be set (use NewReceiveOptions)")
	}
	return nil
}
