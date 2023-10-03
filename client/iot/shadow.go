package iot

import (
	"encoding/json"

	iot20180120 "github.com/alibabacloud-go/iot-20180120/v5/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ShadowMessage struct {
	State struct {
		Desired json.RawMessage `json:"desired"`
	} `json:"state"`
	Metadat struct {
		Desired json.RawMessage `json:"desired"`
	} `json:"metadat"`
	Timestamp int64 `json:"timestamp"`
	Version   int   `json:"version"`
}

func (c *Client) GetDeviceShadow(productKey *string, deviceName *string) *ShadowMessage {
	log := c.log
	log.Trace("GetDeviceShadow productKey : ", productKey, " deviceName : ", deviceName)

	aliIotClient := c.aliIotClient
	request := &iot20180120.GetDeviceShadowRequest{
		ProductKey: tea.String(*productKey),
		DeviceName: tea.String(*deviceName),
	}
	runtime := &util.RuntimeOptions{}

	ret, err := aliIotClient.GetDeviceShadowWithOptions(request, runtime)
	if err != nil {
		log.Error("GetDeviceShadow fail. err : ", err)
		return nil
	}
	log.Trace("GetDeviceShadow result : ", ret)

	if ret.Body == nil {
		log.Info("GetDeviceShadow fail. ret : ", ret)
		return nil
	}

	body := ret.Body
	if body.ShadowMessage == nil {
		log.Info("GetDeviceShadow fail. body : ", body)
		return nil
	}

	var msg ShadowMessage
	err = json.Unmarshal([]byte(*body.ShadowMessage), &msg)
	if err != nil {
		log.Error("GetDeviceShadow Unmarshal ShadowMessage fail. err : ", err)
		return nil
	}
	return &msg
}
