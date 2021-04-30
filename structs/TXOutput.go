package structs

type TXOutput struct {
	Value        int
	ScriptPubKey string // receiver's address
}

// CanBeUnlockedWith checks if this TXOutput was to be sent to a specific user
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

