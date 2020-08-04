package models

const DistributionCanceled = 1
const FraudReportGoodsDispatchedAndInTransit = 2
const FraudReportServiceNowCancelled = 3
const FraudReportServiceRedeemed = 4
const FraudReportAcknowledged = 5
const FraudReportTransactionReversed = 6
const FraudReportError = 7

const DistributionCanceledMessage = "goods-distribution-cancelled"
const FraudReportGoodsDispatchedAndInTransitMessage = "goods-dispatched-and-in-transit"
const FraudReportServiceNowCancelledMessage = "service-now-cancelled"
const FraudReportServiceRedeemedMessage = "service-redeemed"
const FraudReportAcknowledgedMessage = "acknowledged"
const FraudReportTransactionReversedMessage = "transaction-reversed"
const FraudReportErrorMessageMessage = "error"

type ResponseAsVmpiClient struct {
	ResponseData `json:"responseData"`
}

type ResponseData struct {
	FraudReportNotificationResponse int    `json:"fraudReportNoticationResponse,omitempty"`
	SystemFraudReport               string `json:"systemFraudReport,omitempty"`
	PurchaseInformationResponseData `json:"purchaseInformationResponseData"`
}

type PurchaseInformationResponseData struct {
	MerchantReferenceNumber      string `json:"merchantReferenceNumber,omitempty"`
	AssuredCredit                bool   `json:"assuredCredit,omitempty"`
	AuthenticationConducted      string `json:"authenticationConducted,omitempty"`
	CustomerName                 string `json:"customerName,omitempty"`
	SpecialInstructions          string `json:"specialInstructions,omitempty"`
	EmailAccount                 string `json:"emailAccount,omitempty"`
	CancellationPolicy           string `json:"cancellationPolicy,omitempty"`
	WebsiteLink                  string `json:"websiteLink,omitempty"`
	CardCVV2ValidationAtPurchase bool   `json:"cardCVV2ValidationAtPurchase,omitempty"`
	DeviceName                   string `json:"deviceName"`
	DeviceId                     string `json:"deviceId"`
	DeviceIpAddress              string `json:"deviceIpAddress"`
	Communications               string `json:"communications"`
	MonthsSinceFirstPurchase     int    `json:"monthsSinceFirstPurchase"`
	LinkToItemSold               string `json:"linkToItemSold,omitempty"`

	CreditInquiryResponse `json:"creditInquiryResponse,omitempty"`
	DigitalReceipt        `json:"digitalReceipt,omitempty"`
}

type DigitalReceipt struct {
	OrderNumber               interface{}        `json:"orderNumber,omitempty"`
	OrderTotal                Amount             `json:"orderTotal,omitempty"`
	Tax                       Amount             `json:"tax,omitempty"`
	PaymentInformation        PaymentInformation `json:"paymentInformation,omitempty"`
	OrderDateTime             interface{}        `json:"orderDateTime,omitempty"`
	ShippingAndHandlingCharge Amount             `json:"shippingAndHandlingCharge,omitempty"`
	TitlesIncludedInOrder     string             `json:"titlesIncludedInOrder,omitempty"`
	OrderDetails              OrderDetails       `json:"orderDetails,omitempty"`
	InvoiceNumber             string             `json:"invoiceNumber,omitempty"`
	SubTotal                  Amount             `json:"subTotal,omitempty"`
}

type PaymentInformation struct {
	ItemsSubTotal  Amount        `json:"itemsSubTotal,omitempty"`
	Tax            Amount        `json:"tax,omitempty"`
	BillingAddress *AddressModel `json:"billingAddress,omitempty"`
	Total          Amount        `json:"total,omitempty"`
	PaymentMethod  string        `json:"paymentMethod,omitempty"`
}

type OrderDetails struct {
	Item []Item `json:"item,omitempty"`
}

type Amount struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type Item struct {
	ItemType        string `json:"itemType,omitempty"`
	Price           Amount `json:"price,omitempty"`
	ItemDescription string `json:"itemDescription,omitempty"`
	ShippingInfo    string `json:"shippingInfo,omitempty"`
	ArtistOrSeller  string `json:"artistOrSeller,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
}

type AddressModel struct {
	City          string `json:"city,omitempty"`
	SubEntityCode string `json:"subEntityCode,omitempty"`
	Address1      string `json:"address1,omitempty"`
	Address2      string `json:"address2,omitempty"`
	State         string `json:"state,omitempty"`
	Country       string `json:"country,omitempty"`
	PostalCode    string `json:"postalCode,omitempty"`
}

type CreditInquiryResponse struct {
	Amount `json:"creditAmount"`
}

type VmpiRequest struct {
	AlertId       string `json:"alertId"`
	RequestHeader `json:"requestHeader"`
	RequestData   `json:"requestData"`
}

type RequestHeader struct {
	User     string   `json:"user"`
	CallType []string `json:"callType"`
}

type RequestData struct {
	TransactionAmount `json:"transactionAmount"`
	TransactionId     string `json:"transactionId"`
}

type TransactionAmount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

func (r *ResponseAsVmpiClient) AddVisaRequest(data string) {
	r.FraudReportNotificationResponse = 6
	r.SystemFraudReport = "transaction-reversed"
	r.PurchaseInformationResponseData.SpecialInstructions = data
}

func (r *VmpiRequest) GetResponseForVmpiByAmount() ResponseAsVmpiClient {
	switch r.RequestData.TransactionAmount.Value {
	case "1.00":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: DistributionCanceled,
			SystemFraudReport:               DistributionCanceledMessage,
		}}
	case "2.00":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: FraudReportGoodsDispatchedAndInTransit,
			SystemFraudReport:               FraudReportGoodsDispatchedAndInTransitMessage,
		}}
	case "3.00":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: FraudReportServiceNowCancelled,
			SystemFraudReport:               FraudReportServiceNowCancelledMessage,
		}}
	case "4.00":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: FraudReportServiceRedeemed,
			SystemFraudReport:               FraudReportServiceRedeemedMessage,
		}}
	case "5.00":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: FraudReportAcknowledged,
			SystemFraudReport:               FraudReportAcknowledgedMessage,
		}}
	case "6.00":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: FraudReportTransactionReversed,
			SystemFraudReport:               FraudReportTransactionReversedMessage,
		}}
	case "7.00":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: FraudReportTransactionReversed,
			SystemFraudReport:               "",
		}}
	}
	return ResponseAsVmpiClient{ResponseData{
		FraudReportNotificationResponse: FraudReportError,
		SystemFraudReport:               FraudReportErrorMessageMessage,
	}}
}

func (r *VmpiRequest) GetResponseForVmpiByTransactionId() ResponseAsVmpiClient {
	switch r.RequestData.TransactionId {
	case "20200617171013602256":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: FraudReportTransactionReversed,
			SystemFraudReport:               FraudReportTransactionReversedMessage,
			PurchaseInformationResponseData: PurchaseInformationResponseData{
				EmailAccount:                 "byron.wallis@test.com",
				MerchantReferenceNumber:      "19532887",
				CancellationPolicy:           "https://payproglobal.com/",
				WebsiteLink:                  "https://payproglobal.com/",
				CardCVV2ValidationAtPurchase: true,
				DeviceName:                   "iphone x",
				DeviceId:                     "g35hy4h467",
				DeviceIpAddress:              "37.172.47.213",
				LinkToItemSold:               "https://payproglobal.com/",
				CustomerName:                 "byron wallis",
				DigitalReceipt: DigitalReceipt{
					OrderTotal: Amount{
						Value:    45.77,
						Currency: "840",
					},
					Tax: Amount{
						Value:    7.63,
						Currency: "840",
					},
					OrderDateTime:         "2020-05-28T11:34:10.63",
					TitlesIncludedInOrder: "PC Privacy Shield",
					PaymentInformation: PaymentInformation{
						PaymentMethod: "Card 4111",
					},
					OrderDetails: OrderDetails{
						Item: []Item{
							{
								Quantity: 1,
								Price: Amount{
									Value:    38.14,
									Currency: "840",
								},
							},
						},
					},
				},
			},
		}}
	case "11111111111111111111":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: FraudReportTransactionReversed,
			SystemFraudReport:               FraudReportTransactionReversedMessage,
			PurchaseInformationResponseData: PurchaseInformationResponseData{
				EmailAccount:                 "cristian.mustea10@test.com",
				MerchantReferenceNumber:      "53349599329239481",
				CancellationPolicy:           "https://bananame.club/",
				WebsiteLink:                  "https://bananame.club/",
				CardCVV2ValidationAtPurchase: true,
				DeviceName:                   "redmi 8pro",
				DeviceId:                     "k8l798kim",
				DeviceIpAddress:              "77.100.203.146",
				MonthsSinceFirstPurchase:     1,
				LinkToItemSold:               "https://bananame.club/",
				CustomerName:                 "Cristian Mustea",
				DigitalReceipt: DigitalReceipt{
					OrderTotal: Amount{
						Value:    44.99,
						Currency: "840",
					},
					OrderDateTime:         "2020-05-25T00:01:11",
					TitlesIncludedInOrder: "Premium package 53349599329239481",
					PaymentInformation: PaymentInformation{
						PaymentMethod: "Card 4141",
					},
					OrderDetails: OrderDetails{
						Item: []Item{
							{
								Quantity: 1,
								Price: Amount{
									Value:    44.99,
									Currency: "840",
								},
							},
						},
					},
				},
			},
		}}
	case "00000000000000000001":
		return ResponseAsVmpiClient{ResponseData{
			FraudReportNotificationResponse: FraudReportTransactionReversed,
			SystemFraudReport:               FraudReportTransactionReversedMessage,
			PurchaseInformationResponseData: PurchaseInformationResponseData{
				EmailAccount:                 "betterman@gmail.com",
				MerchantReferenceNumber:      "wer3424r",
				CancellationPolicy:           "https://betterme-apps.com/en",
				WebsiteLink:                  "https://betterme-apps.com/en",
				CardCVV2ValidationAtPurchase: true,
				DeviceName:                   "iphone se alex",
				DeviceId:                     "f4134f1234234f3",
				DeviceIpAddress:              "8.8.8.8",
				Communications:               "Support gave a feedback and refund provided",
				MonthsSinceFirstPurchase:     5,
				LinkToItemSold:               "https://betterme-apps.com/en",
				CustomerName:                 "Alex Surovskii",
				DigitalReceipt: DigitalReceipt{
					OrderTotal: Amount{
						Value:    20.20,
						Currency: "840",
					},
					OrderDateTime:         "2020-06-23T13:33:29",
					TitlesIncludedInOrder: "Meal plan for 7 weeks",
					PaymentInformation: PaymentInformation{
						PaymentMethod: "Card 5432",
					},
					OrderDetails: OrderDetails{
						Item: []Item{
							{
								Quantity: 1,
								Price: Amount{
									Value:    20.20,
									Currency: "840",
								},
							},
						},
					},
				},
			},
		}}
	}

	return ResponseAsVmpiClient{ResponseData{
		FraudReportNotificationResponse: FraudReportError,
		SystemFraudReport:               FraudReportErrorMessageMessage,
	}}
}
