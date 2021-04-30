package structs

type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string // sender's address
}

// CanUnlockOutputWith checks if an user can use this TXInput to create a TXOutput
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}