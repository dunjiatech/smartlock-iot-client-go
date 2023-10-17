package iot

import (
	"encoding/json"
	"fmt"

	iot20180120 "github.com/alibabacloud-go/iot-20180120/v5/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ShadowMessage struct {
	State struct {
		Desired json.RawMessage `json:"desired"`
	} `json:"state"`
	Metadata struct {
		Desired json.RawMessage `json:"desired"`
	} `json:"metadata"`
	Timestamp int64 `json:"timestamp"`
	Version   int   `json:"version"`
}

func (c *Client) GetDeviceShadow(productKey string, deviceName string) (*ShadowMessage, error) {
	log := c.log
	log.Trace("GetDeviceShadow productKey : ", productKey, " deviceName : ", deviceName)

	request := &iot20180120.GetDeviceShadowRequest{
		ProductKey: tea.String(productKey),
		DeviceName: tea.String(deviceName),
	}
	runtime := &util.RuntimeOptions{}

	aliIotClient := c.aliIotClient
	ret, err := aliIotClient.GetDeviceShadowWithOptions(request, runtime)
	if err != nil {
		log.Error("aliIotClient error. err : ", err)
		return nil, err
	}
	log.Trace("GetDeviceShadow result : ", ret)

	if ret.Body == nil {
		log.Info("GetDeviceShadow ret.body is nil. ret : ", ret)
		return nil, fmt.Errorf("", ret)
	}

	body := ret.Body

	if *body.Success == false {
		log.Info("GetDeviceShadow body.Success is false. body : ", body)
		return nil, fmt.Errorf("", body)
	}

	if body.ShadowMessage == nil {
		var m = ShadowMessage{}
		return &m, nil
	}

	var msg ShadowMessage
	err = json.Unmarshal([]byte(*body.ShadowMessage), &msg)
	if err != nil {
		log.Error("GetDeviceShadow Unmarshal ShadowMessage fail. err : ", err)
		return nil, err
	}
	return &msg, nil
}

func (c *Client) UpdateDeviceShadow(productKey string, deviceName string, desired interface{}, version int, delta bool) (bool, error) {
	log := c.log
	log.Trace("UpdateDeviceShadow productKey : ", productKey, " deviceName : ", deviceName, " desired : ", desired, " delta : ", delta)

	shadowMessage := struct {
		Method string `json:"method"`
		State  struct {
			Desired interface{} `json:"desired"`
		} `json:"state"`
		Version int `json:"version"`
	}{
		Method:  "update",
		Version: version,
	}
	shadowMessage.State.Desired = desired
	strShadowMessage, _ := json.Marshal(shadowMessage)

	request := &iot20180120.UpdateDeviceShadowRequest{
		ProductKey:    tea.String(productKey),
		DeviceName:    tea.String(deviceName),
		ShadowMessage: tea.String(string(strShadowMessage)),
		DeltaUpdate:   tea.Bool(delta),
	}
	runtime := &util.RuntimeOptions{}

	aliIotClient := c.aliIotClient
	ret, err := aliIotClient.UpdateDeviceShadowWithOptions(request, runtime)
	if err != nil {
		log.Error("aliIotClient error. err : ", err)
		return false, err
	}
	log.Trace("UpdateDeviceShadow result : ", ret)

	if ret.Body == nil {
		log.Info("UpdateDeviceShadow return body is nil. ret : ", ret)
		return false, fmt.Errorf("", ret)
	}

	body := ret.Body

	if *body.Success == false {
		log.Info("UpdateDeviceShadow body.Success is false. body : ", body)
		return false, fmt.Errorf("", body)
	}
	return true, nil
}

func (c *Client) UpdateDeviceShadowEx(productKey string, deviceName string, desired interface{}, delta bool) (bool, error) {
	m, err := c.GetDeviceShadow(productKey, deviceName)
	if m == nil {
		return false, err
	}
	version := (m.Version + 11) / 10 * 10
	return c.UpdateDeviceShadow(productKey, deviceName, desired, version, delta)
}
