package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func (c *Client) SetOpenMode(productKey string, deviceName string, mode int) (bool, error) {
	log := c.log
	log.Trace("SetOpenMode mode : ", mode)

	if mode != 1 && mode != 2 {
		return false, fmt.Errorf("unsupport this open mode. mode:", mode)
	}

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

func (c *Client) GetBatteryStatus(productKey string, deviceName string, index int) (*BatteryStatus, error) {
	log := c.log
	log.Trace("GetBatteryStatus. ")

	props, err := c.LivingClient.GetThingProperties(productKey, deviceName)
	if err != nil {
		return nil, err
	}

	inplaceName := "DjBatteryInplace" + strconv.Itoa(index)
	voltageName := "DjBatteryVoltage" + strconv.Itoa(index)

	batteryStatus := BatteryStatus{
		Inplace: false,
		Voltage: 0,
	}
	for _, p := range *props {
		if p.Attribute == inplaceName {
			var value bool
			json.Unmarshal(p.Value, &value)
			batteryStatus.Inplace = value
		} else if p.Attribute == voltageName {
			var value int
			json.Unmarshal(p.Value, &value)
			batteryStatus.Voltage = value
		}
	}
	log.Trace("GetBatteryStatus result : ", batteryStatus)

	return &batteryStatus, nil
}
