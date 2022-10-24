package mixpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// get by mixpay result traceId

type Data struct {
	Status           string `json:"status"`
	QuoteAmount      string `json:"quoteAmount"`
	QuoteSymbol      string `json:"quoteSymbol"`
	PaymentAmount    string `json:"paymentAmount"`
	PaymentSymbol    string `json:"paymentSymbol"`
	Payee            string `json:"payee"`
	PayeeMixinNumber string `json:"payeeMixinNumber"`
	PayeeAvatarURL   string `json:"payeeAvatarUrl"`
	Txid             string `json:"txid"`
	Date             int64  `json:"date"`
	SurplusAmount    string `json:"surplusAmount"`
	SurplusStatus    string `json:"surplusStatus"`
	Confirmations    int64  `json:"confirmations"`
	PayableAmount    string `json:"payableAmount"`
	FailureCode      string `json:"failureCode"`
	FailureReason    string `json:"failureReason"`
	ReturnTo         string `json:"returnTo"`
	TraceID          string `json:"traceId"`
}

type MixpayResult struct {
	Code        int64  `json:"code"`
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	Data        Data   `json:"data"`
	TimestampMS int64  `json:"timestampMs"`
}

// type MixpayRequest struct {
// 	TraceId string `json:"traceId"`
// }

func GetMixpayResult(traceId string) (MixpayResult, error) {
	payload := "{\"traceId\":\"" + traceId + "\"}"

	req, err := http.NewRequest("GET", "https://api.mixpay.me/v1/payments_result", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return MixpayResult{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return MixpayResult{}, err
	}
	defer res.Body.Close()

	var mixpayResult MixpayResult
	if err = json.NewDecoder(res.Body).Decode(&mixpayResult); err != nil {
		return MixpayResult{}, err
	}

	if mixpayResult.Data.Status != "success" {
		return MixpayResult{}, err
	}

	fmt.Println(mixpayResult)

	return mixpayResult, nil
}
