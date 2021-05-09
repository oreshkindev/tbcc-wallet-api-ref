package model

import (
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/util"
	"gopkg.in/guregu/null.v4"
)

// User struct
type User struct {
	ID          util.UUID `json:"id" sql:",type:uuid"`
	Useraddress []string  `json:"useraddress"`
	Accounttype string    `json:"accounttype"`
	Smartcard   bool      `json:"smartcard"`
	VpnKeys     []VpnKey  `json:"vpn_keys"`
}

// UserMigrate struct
type UserMigrate struct {
	ID            int
	Address       string
	PaidFee       null.Float
	PaidSmartcard null.Float
}
