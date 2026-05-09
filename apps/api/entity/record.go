package entity

import "time"

type RecordType string
type PaymentType string
type PaymentStatus string

const (
	RecordTypeIncome   RecordType = "INCOME"
	RecordTypeExpense  RecordType = "EXPENSE"
	RecordTypeTransfer RecordType = "TRANSFER"

	PaymentTypeCash          PaymentType = "CASH"
	PaymentTypeDebitCard     PaymentType = "DEBIT_CARD"
	PaymentTypeCreditCard    PaymentType = "CREDIT_CARD"
	PaymentTypeBankTransfer  PaymentType = "BANK_TRANSFER"
	PaymentTypeVoucher       PaymentType = "VOUCHER"
	PaymentTypeMobilePayment PaymentType = "MOBILE_PAYMENT"
	PaymentTypeWebPayment    PaymentType = "WEB_PAYMENT"

	PaymentStatusCleared    PaymentStatus = "CLEARED"
	PaymentStatusUncleared  PaymentStatus = "UNCLEARED"
	PaymentStatusReconciled PaymentStatus = "RECONCILED"
)

// RecordCursor is the keyset position for cursor-based pagination.
// It encodes the last seen (RecordedAt, ID) pair, matching ORDER BY recorded_at DESC, id DESC.
type RecordCursor struct {
	RecordedAt time.Time
	ID         string
}

// Record maps to the `records` table.
// Labels are stored separately in the `record_labels` junction table.
type Record struct {
	ID            string        `json:"id" db:"id"`
	UserID        string        `json:"user_id" db:"user_id"`
	AccountID     string        `json:"account_id" db:"account_id"`
	SubCategoryID string        `json:"subcategory_id" db:"subcategory_id"`
	Amount        float64       `json:"amount" db:"amount"`
	Currency      CurrencyType  `json:"currency" db:"currency"`
	BaseAmount    float64       `json:"base_amount" db:"base_amount"`
	Type          RecordType    `json:"type" db:"type"`
	Name          string        `json:"name" db:"name"`
	Payee         string        `json:"payee" db:"payee"`
	PaymentType   PaymentType   `json:"payment_type" db:"payment_type"`
	PaymentStatus PaymentStatus `json:"payment_status" db:"payment_status"`
	IsExcluded    bool          `json:"is_excluded" db:"is_excluded"`
	RecordedAt    time.Time     `json:"recorded_at" db:"recorded_at"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at"`
}

// RecordLabel maps to the `record_labels` junction table.
type RecordLabel struct {
	RecordID string `json:"record_id" db:"record_id"`
	LabelID  string `json:"label_id" db:"label_id"`
}
