// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetPaymentReceipt Returns information about a successful payment
// @param chatID Chat identifier of the PaymentSuccessful message
// @param messageID Message identifier
func (client *Client) GetPaymentReceipt(chatID int64, messageID int64) (*tdlib.PaymentReceipt, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getPaymentReceipt",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var paymentReceipt tdlib.PaymentReceipt
	err = json.Unmarshal(result.Raw, &paymentReceipt)
	return &paymentReceipt, err

}