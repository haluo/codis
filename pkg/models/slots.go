package models

const MaxSlotNum = 1024

type SlotInfo struct {
	Target      string `json:"target"`
	MigrateFrom string `json:"migrate_from,omitempty"`
	Locked      bool   `json:"locked,omitempty"`
}
