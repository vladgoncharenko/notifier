package models

type StatusSolidGate struct {
	Order struct {
		OrderId           string `json:"order_id,omitempty"`
		Status            string `json:"status,omitempty"`
		Amount            int    `json:"amount,omitempty"`
		RefundedAmount    string `json:"refunded_amount,omitempty"`
		Currency          string `json:"currency,omitempty"`
		MarketingAmount   string `json:"marketing_amount,omitempty"`
		MarketingCurrency string `json:"marketing_currency,omitempty"`
		Fraudulent        bool   `json:"fraudulent,omitempty"`
		TotalFeeAmount    int    `json:"total_fee_amount,omitempty"`
	} `json:"order,omitempty"`

	Transactions []struct {
		Type       string `json:"type,omitempty"`
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

	CardToken struct {
		Token        string `json:"token,omitempty"`
		Bin          string `json:"bin,omitempty"`
		Brand        string `json:"brand,omitempty"`
		Country      string `json:"country,omitempty"`
		Number       string `json:"number,omitempty"`
		CardExpMonth string `json:"card_exp_month,omitempty"`
		CardExpYear  string `json:"card_exp_year,omitempty"`
		CardType     string `json:"card_type,omitempty"`
	} `json:"card_token,omitempty"`

	VerifyUrl string `json:"verify_url,omitempty,omitempty"`
}
