// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// ResetPassword Removes 2-step verification password without previous password and access to recovery email address. The password can't be reset immediately and the request needs to be repeated after the specified time
func (client *Client) ResetPassword() (tdlib.ResetPasswordResult, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "resetPassword",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.ResetPasswordResultEnum(result.Data["@type"].(string)) {

	case tdlib.ResetPasswordResultOkType:
		var resetPasswordResult tdlib.ResetPasswordResultOk
		err = json.Unmarshal(result.Raw, &resetPasswordResult)
		return &resetPasswordResult, err

	case tdlib.ResetPasswordResultPendingType:
		var resetPasswordResult tdlib.ResetPasswordResultPending
		err = json.Unmarshal(result.Raw, &resetPasswordResult)
		return &resetPasswordResult, err

	case tdlib.ResetPasswordResultDeclinedType:
		var resetPasswordResult tdlib.ResetPasswordResultDeclined
		err = json.Unmarshal(result.Raw, &resetPasswordResult)
		return &resetPasswordResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
