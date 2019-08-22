package models

type StatusSolidGate struct {
	PayForm struct {
		Token string `json:"token, omitempty"`
	} `json:"pay_form,omitempty"`

	Order struct {
		Order_id           string `json:"order_id,omitempty"`
		Status             string `json:"status,omitempty"`
		Amount             int    `json:"amount,omitempty"`
		Refunded_amount    string `json:"refunded_amount,omitempty"`
		Currency           string `json:"currency,omitempty"`
		Marketing_amount   string `json:"marketing_amount,omitempty"`
		Marketing_currency string `json:"marketing_currency,omitempty"`
		Fraudulent         bool   `json:"fraudulent,omitempty"`
		Total_fee_amount   int    `json:"total_fee_amount,omitempty"`
	} `json:"order,omitempty"`

	Transactions [] struct {
		Id         string `json:"id,omitempty"`
		Operation  string `json:"operation,omitempty"`
		Status     string `json:"status,omitempty"`
		Descriptor string `json:"descriptor,omitempty"`
		Amount     int    `json:"amount,omitempty"`
		Currency   string `json:"currency,omitempty"`
		Fee        struct {
			Amount   int    `json:"amount,omitempty"`
			Currency string `json:"currency,omitempty"`
		} `json:"fee,omitempty"`
	} `json:"transactions,omitempty"`

	Card struct {
		Bank           string `json:"bank,omitempty"`
		Bin            string `json:"bin,omitempty"`
		Brand          string `json:"brand,omitempty"`
		Country        string `json:"country,omitempty"`
		Number         string `json:"number,omitempty"`
		Card_exp_month string `json:"card_exp_month,omitempty"`
		Card_exp_year  string `json:"card_exp_year,omitempty"`
		Card_type      string `json:"card_type,omitempty"`
		Card_token     struct {
			Token string `json:"token,omitempty"`
		} `json:"card_token,omitempty"`
	} `json:"card,omitempty"`

	Verify_url string `json:"verify_url,omitempty,omitempty"`
}
