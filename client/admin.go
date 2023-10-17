package client

func (c *Client) SetAdminPassword(productKey string, deviceName string, passsword string) (bool, error) {
	log := c.log
	log.Trace("SetAdminPassword. ")

	desired := struct {
		Pwd string `json:"prh_admin_pwd_1"`
	}{
		Pwd: passsword,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}
