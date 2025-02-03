// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type AccountStatus string

const (
	AccountStatusACTIVE AccountStatus = "ACTIVE"
	AccountStatusCLOSED AccountStatus = "CLOSED"
)

func (e *AccountStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccountStatus(s)
	case string:
		*e = AccountStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for AccountStatus: %T", src)
	}
	return nil
}

type NullAccountStatus struct {
	AccountStatus AccountStatus `json:"account_status"`
	Valid         bool          `json:"valid"` // Valid is true if AccountStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccountStatus) Scan(value interface{}) error {
	if value == nil {
		ns.AccountStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccountStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccountStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AccountStatus), nil
}

type Account struct {
	ID        int64         `json:"id"`
	Owner     string        `json:"owner"`
	Balance   int64         `json:"balance"`
	Currency  string        `json:"currency"`
	Status    AccountStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
