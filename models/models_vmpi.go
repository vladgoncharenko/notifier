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
	MerchantReferenceNumber string      `json:"merchantReferenceNumber"`
	AssuredCredit           bool        `json:"assuredCredit"`
	AuthenticationConducted string      `json:"authenticationConducted"`
	CustomerName            string      `json:"customerName"`
	SpecialInstructions     interface{} `json:"specialInstructions"`
	CreditInquiryResponse   `json:"creditInquiryResponse"`
}

type CreditInquiryResponse struct {
	CreditAmount `json:"creditAmount"`
}

type CreditAmount struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
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
}

type TransactionAmount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

func (r *ResponseAsVmpiClient) AddVisaRequest(data interface{}) {
	r.FraudReportNotificationResponse = 6
	r.SystemFraudReport = "transaction-reversed"
	r.PurchaseInformationResponseData.SpecialInstructions = data
}

func (r *VmpiRequest) GetResponseForVmpi() ResponseAsVmpiClient {
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
	}
	return ResponseAsVmpiClient{ResponseData{
		FraudReportNotificationResponse: FraudReportError,
		SystemFraudReport:               FraudReportErrorMessageMessage,
	}}
}
