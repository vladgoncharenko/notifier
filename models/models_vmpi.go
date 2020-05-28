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
}

type VmpiRequest struct {
	AlertId        string `json:"alertId"`
	*RequestHeader `json:"requestHeader"`
}

type RequestHeader struct {
	User     string   `json:"user"`
	CallType []string `json:"callType"`
}

func (r *VmpiRequest) GetResponseForVmpi() ResponseAsVmpiClient {
	for _, v := range r.CallType {
		switch v {
		case "1":
			return ResponseAsVmpiClient{ResponseData{
				FraudReportNotificationResponse: DistributionCanceled,
				SystemFraudReport:               DistributionCanceledMessage,
			}}
		case "2":
			return ResponseAsVmpiClient{ResponseData{
				FraudReportNotificationResponse: FraudReportGoodsDispatchedAndInTransit,
				SystemFraudReport:               FraudReportGoodsDispatchedAndInTransitMessage,
			}}
		case "3":
			return ResponseAsVmpiClient{ResponseData{
				FraudReportNotificationResponse: FraudReportServiceNowCancelled,
				SystemFraudReport:               FraudReportServiceNowCancelledMessage,
			}}
		case "4":
			return ResponseAsVmpiClient{ResponseData{
				FraudReportNotificationResponse: FraudReportServiceRedeemed,
				SystemFraudReport:               FraudReportServiceRedeemedMessage,
			}}
		case "5":
			return ResponseAsVmpiClient{ResponseData{
				FraudReportNotificationResponse: FraudReportAcknowledged,
				SystemFraudReport:               FraudReportAcknowledgedMessage,
			}}
		case "6":
			return ResponseAsVmpiClient{ResponseData{
				FraudReportNotificationResponse: FraudReportTransactionReversed,
				SystemFraudReport:               FraudReportTransactionReversedMessage,
			}}
		}
	}
	return ResponseAsVmpiClient{ResponseData{
		FraudReportNotificationResponse: FraudReportError,
		SystemFraudReport:               FraudReportErrorMessageMessage,
	}}
}
