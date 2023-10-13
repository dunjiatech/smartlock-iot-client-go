package client

import (
	"encoding/json"
	"time"
)

func (c *Client) SetOpenMode(productKey string, deviceName string, mode int) bool {
	log := c.log
	log.Trace("SetOpenMode mode : ", mode)

	desired := struct {
		OpenMode int `json:"open_mode"`
	}{
		OpenMode: mode,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}

func (c *Client) GetLockStatus(productKey string, deviceName string) (int, error) {
	log := c.log
	log.Trace("GetLockStatus. ")

	ret, err := c.LivingClient.GetThingStatus(productKey, deviceName)
	if err != nil {
		return 0, err
	}

	status := ret.Status
	secs := time.Now().Unix()*1000 - 24*60*60*1000

	if status == 3 && ret.Time > secs {
		status = 1
	}
	return status, nil
}

type BatteryStatus struct {
	Inplace bool
	Voltage int
}

func (c *Client) GetBattery0Status(productKey string, deviceName string) (*BatteryStatus, error) {
	log := c.log
	log.Trace("GetBattery0Status. ")

	props, err := c.LivingClient.GetThingProperties(productKey, deviceName)
	if err != nil {
		return nil, err
	}

	batteryStatus := BatteryStatus{
		Inplace: false,
		Voltage: 0,
	}
	for _, p := range *props {
		if p.Attribute == "DjBatteryInplace0" {
			var value bool
			json.Unmarshal(p.Value, &value)
			batteryStatus.Inplace = value
		} else if p.Attribute == "DjBatteryVoltage0" {
			var value int
			json.Unmarshal(p.Value, &value)
			batteryStatus.Voltage = value
		}
	}
	log.Trace("GetBattery0Status result : ", batteryStatus)

	return &batteryStatus, nil
}
