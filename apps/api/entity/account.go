package entity

import "time"

type CurrencyType string
type AccountType string

const (
	CurrencyTypeIDR CurrencyType = "IDR"
	CurrencyTypeUSD CurrencyType = "USD"
	CurrencyTypeSGD CurrencyType = "SGD"
	CurrencyTypeEUR CurrencyType = "EUR"
	CurrencyTypeGBP CurrencyType = "GBP"
	CurrencyTypeJPY CurrencyType = "JPY"
	CurrencyTypeKRW CurrencyType = "KRW"
	CurrencyTypeTHB CurrencyType = "THB"
	CurrencyTypeETH CurrencyType = "ETH"
	CurrencyTypeBTC CurrencyType = "BTC"
	CurrencyTypeSOL CurrencyType = "SOL"

	AccountTypeGeneral        AccountType = "GENERAL"
	AccountTypeCash           AccountType = "CASH"
	AccountTypeCurrentAccount AccountType = "CURRENT_ACCOUNT"
	AccountTypeCreditCard     AccountType = "CREDIT_CARD"
	AccountTypeSavingAccount  AccountType = "SAVING_ACCOUNT"
	AccountTypeBonus          AccountType = "BONUS"
	AccountTypeInsurance      AccountType = "INSURANCE"
	AccountTypeInvestment     AccountType = "INVESTMENT"
	AccountTypeLoan           AccountType = "LOAN"
	AccountTypeMortgage       AccountType = "MORTGAGE"
)

type Account struct {
	ID            string       `json:"id" db:"id"`
	UserID        string       `json:"user_id" db:"user_id"`
	Name          string       `json:"name" db:"name"`
	Balance       float64      `json:"balance" db:"balance"`
	Currency      CurrencyType `json:"currency" db:"currency"`
	Type          AccountType  `json:"type" db:"type"`
	Color         string       `json:"color" db:"color"`
	IsArchived    bool         `json:"is_archived" db:"is_archived"`
	IsExcluded    bool         `json:"is_excluded" db:"is_excluded"`
	MarkForDelete bool         `json:"mark_for_delete" db:"mark_for_delete"`
	CreatedAt     time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
}
