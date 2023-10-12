package client

import "time"

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
