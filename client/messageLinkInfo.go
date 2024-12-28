// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetMessageLinkInfo Returns information about a public or private message link. Can be called for any internal link of the type internalLinkTypeMessage
// @param uRL The message link
func (client *Client) GetMessageLinkInfo(uRL string) (*tdlib.MessageLinkInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getMessageLinkInfo",
		"url":   uRL,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageLinkInfo tdlib.MessageLinkInfo
	err = json.Unmarshal(result.Raw, &messageLinkInfo)
	return &messageLinkInfo, err

}
