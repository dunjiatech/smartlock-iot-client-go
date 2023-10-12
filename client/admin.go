package client

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

func (c *Client) SetAdminPassword(productKey string, deviceName string, passsword string) bool {
	log := c.log
	log.Trace("SetAdminPassword. ")

	desired := struct {
		Pwd string `json:"prh_admin_pwd_1"`
	}{
		Pwd: passsword,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}
