# Vinance — Database Schema

> Redesigned per [VNC-37](https://linear.app/vinesse/issue/VNC-37/review-redesign-database-structure).
> Changes from previous version are annotated with `-- CHANGED` / `-- ADDED` / `-- REMOVED`.

---

## Tables

### `users`

| Column          | Type       | Constraints              | Notes                       |
|-----------------|------------|--------------------------|-----------------------------|
| `id`            | varchar    | PK                       | UUID                        |
| `email`         | varchar    | NOT NULL, UNIQUE         |                             |
| `password_hash` | varchar    | NOT NULL                 |                             |
| `username`      | varchar    | NOT NULL, UNIQUE         |                             |
| `phone_number`  | varchar    | NOT NULL, UNIQUE         |                             |
| `dob`           | datetime   | NOT NULL                 |                             |
| `created_at`    | datetime   |                          |                             |
| `updated_at`    | datetime   |                          |                             |

> **Removed:** `gender` — no longer required.

---

### `accounts`

| Column          | Type            | Constraints                      | Notes                                             |
|-----------------|-----------------|----------------------------------|---------------------------------------------------|
| `id`            | varchar         | PK                               | UUID                                              |
| `user_id`       | varchar         | NOT NULL, FK → `users.id`        |                                                   |
| `name`          | varchar         | NOT NULL, UNIQUE                 |                                                   |
| `balance`       | decimal(18,6)   | NOT NULL, default 0              | **Changed** from `decimal(4)` — 6 decimal places |
| `currency`      | currency_type   | NOT NULL, default `IDR`          | Base currency of the account                      |
| `type`          | account_type    | NOT NULL                         |                                                   |
| `color`         | varchar         | NOT NULL                         |                                                   |
| `is_archived`   | bool            | default false                    |                                                   |
| `is_excluded`   | bool            | default false                    |                                                   |
| `mark_for_delete` | bool          | default false                    |                                                   |
| `created_at`    | datetime        |                                  |                                                   |
| `updated_at`    | datetime        |                                  |                                                   |

---

### `records`

| Column            | Type            | Constraints                        | Notes                                                                                                                     |
|-------------------|-----------------|------------------------------------|---------------------------------------------------------------------------------------------------------------------------|
| `id`              | varchar         | PK                                 | UUID                                                                                                                      |
| `user_id`         | varchar         | NOT NULL, FK → `users.id`          |                                                                                                                           |
| `account_id`      | varchar         | NOT NULL, FK → `accounts.id`       |                                                                                                                           |
| `subcategory_id`  | varchar         | NOT NULL, FK → `subcategories.id`  |                                                                                                                           |
| `amount`          | decimal(18,6)   | NOT NULL                           | **Changed** from `decimal(4)`. Amount in `currency` (the transaction currency).                                          |
| `currency`        | currency_type   | NOT NULL                           | **Added.** Currency used for this transaction — may differ from the account's base currency.                             |
| `base_amount`     | decimal(18,6)   | NOT NULL                           | **Added.** Amount converted to the account's base currency. Equals `amount` when `currency == account.currency`.         |
| `type`            | record_type     | NOT NULL                           |                                                                                                                           |
| `name`            | varchar         |                                    | Optional label / memo for the transaction                                                                                 |
| `payee`           | varchar         |                                    |                                                                                                                           |
| `payment_type`    | payment_type    |                                    |                                                                                                                           |
| `payment_status`  | payment_status  | NOT NULL                           |                                                                                                                           |
| `is_excluded`     | bool            | default false                      |                                                                                                                           |
| `recorded_at`     | datetime        |                                    | When the transaction actually occurred (user-supplied)                                                                    |
| `created_at`      | datetime        |                                    |                                                                                                                           |
| `updated_at`      | datetime        |                                    |                                                                                                                           |

> **Removed:** `labels varchar` column — replaced by the `record_labels` junction table (see below).
>
> **Dual-currency design:** when a user records a 100 THB transaction on an IDR-based account, `amount = 100`, `currency = THB`, `base_amount = 50000` (user manually inputs the converted value). The UI should prompt for `base_amount` whenever `currency ≠ account.currency`.

---

### `record_labels` *(junction table)*

| Column       | Type    | Constraints                   | Notes                          |
|--------------|---------|-------------------------------|--------------------------------|
| `record_id`  | varchar | NOT NULL, FK → `records.id`   | Composite PK with `label_id`   |
| `label_id`   | varchar | NOT NULL, FK → `labels.id`    |                                |

> **Added.** Replaces the `labels varchar` column on `records` with a proper many-to-many relation.

---

### `labels`

| Column          | Type     | Constraints               | Notes  |
|-----------------|----------|---------------------------|--------|
| `id`            | varchar  | PK                        | UUID   |
| `user_id`       | varchar  | NOT NULL, FK → `users.id` |        |
| `name`          | varchar  | NOT NULL                  |        |
| `is_archived`   | bool     | default false             |        |
| `mark_for_delete` | bool   | default false             |        |
| `created_at`    | datetime |                           |        |
| `updated_at`    | datetime |                           |        |

---

### `categories`

| Column          | Type     | Constraints   | Notes                                  |
|-----------------|----------|---------------|----------------------------------------|
| `id`            | varchar  | PK            | UUID                                   |
| `name`          | varchar  | NOT NULL      |                                        |
| `is_system`     | bool     | default true  | System categories cannot be deleted    |
| `is_archived`   | bool     | default false |                                        |
| `mark_for_delete` | bool   | default false |                                        |
| `created_at`    | datetime |               |                                        |
| `updated_at`    | datetime |               |                                        |

---

### `subcategories`

| Column          | Type     | Constraints                     | Notes                                                         |
|-----------------|----------|---------------------------------|---------------------------------------------------------------|
| `id`            | varchar  | PK                              | UUID                                                          |
| `category_id`   | varchar  | NOT NULL, FK → `categories.id`  |                                                               |
| `user_id`       | varchar  | FK → `users.id`                 | NULL for system subcategories; set for user-created ones      |
| `name`          | varchar  | NOT NULL                        |                                                               |
| `is_system`     | bool     | default true                    |                                                               |
| `is_archived`   | bool     | default false                   |                                                               |
| `mark_for_delete` | bool   | default false                   |                                                               |
| `created_at`    | datetime |                                 |                                                               |
| `updated_at`    | datetime |                                 |                                                               |

---

## Enums

### `currency_type`

```
IDR · USD · SGD · EUR · GBP · JPY · KRW · THB · BTC · ETH · SOL
```

### `account_type`

```
GENERAL · CASH · CURRENT_ACCOUNT · CREDIT_CARD · SAVING_ACCOUNT
BONUS · INSURANCE · INVESTMENT · LOAN · MORTGAGE
```

### `record_type`

```
EXPENSE · INCOME · TRANSFER
```

### `payment_type`

```
CASH · DEBIT_CARD · CREDIT_CARD · BANK_TRANSFER
VOUCHER · MOBILE_PAYMENT · WEB_PAYMENT
```

### `payment_status`

```
CLEARED · UNCLEARED · RECONCILED
```

---

## Seed Data — Categories & Subcategories

All entries below are system records (`is_system = true`).

### Food & Drinks
- Food & Drinks (General)
- Groceries
- Restaurant
- Bar, cafe

### Shopping
- Shopping (General)
- Clothes & shoes
- Jewel, accessories
- Health and beauty
- Kids
- Home, garden
- Pets, animals
- Electronics, accessories
- Gifts, joy
- Stationery, tools
- Groceries
- Drug-store, chemist

### Housing
- Housing (General)
- Rent
- Mortgage
- Energy, utilities
- Services
- Maintenance, repairs
- Property insurance

### Transportation
- Transportation (General)
- Public transport
- Taxi
- Long distance
- Business trips

### Vehicle
- Vehicle (General)
- Fuel
- Parking
- Vehicle maintenance
- Toll
- Vehicle insurance
- Leasing

### Life & Entertainment
- Life & Entertainment (General)
- Health care, doctor
- Wellness, beauty
- Active sport, fitness
- Culture, sport events
- Life events
- Hobbies
- Education, development
- Books, audio, subscriptions
- TV, Streaming
- Holiday, trips, hotels
- Charity, gifts
- Alcohol, tobacco
- Lottery, gambling

### Communication, PC
- Communication, PC (General)
- Phone, cell phone
- Internet
- Software, apps, games
- Postal services

### Financial expenses
- Financial expenses (General)
- Taxes
- Insurances
- Loan, interests
- Fines
- Advisory
- Charges, Fees
- Family Support

### Investments
- Investments (General)
- Realty
- Crypto
- Financial investments
- Savings
- Stock

### Income
- Income (General)
- Wage, invoices
- Interests, dividends
- Sale
- Rental income
- Dues & grants
- Lending, renting
- Checks, coupons
- Freelance
- Refunds (tax, purchase)
- Child Support
- Gifts

### Others
- Others (General)
- Missing

---

## Change Summary (vs. previous schema)

| Area           | Change                                                                                                       |
|----------------|--------------------------------------------------------------------------------------------------------------|
| `users`        | Removed `gender` column                                                                                      |
| `accounts`     | `balance` precision upgraded to `decimal(18,6)`                                                             |
| `records`      | `amount` upgraded to `decimal(18,6)`; added `currency` and `base_amount` for dual-currency support          |
| `records`      | Removed `labels varchar`; replaced by `record_labels` junction table                                        |
| `record_labels`| New junction table for many-to-many record ↔ label relationship                                             |
| Categories     | Defined 11 system parent categories and all subcategories as seed data                                       |
