package client

func (c *Client) SetTenantPwd1(productKey string, deviceName string, passsword string) bool {
	log := c.log
	log.Trace("SetTenantPassword1. ")

	desired := struct {
		Password string `json:"prh_tenant_pwd_1"`
	}{
		Password: passsword,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}

func (c *Client) SetTenantPwd2(productKey string, deviceName string, passsword string) bool {
	log := c.log
	log.Trace("SetTenantPassword2. ")

	desired := struct {
		Password string `json:"prh_tenant_pwd_2"`
	}{
		Password: passsword,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}

func (c *Client) ClearAllTenantPwd(productKey string, deviceName string) bool {
	log := c.log
	log.Trace("ClearAllTenantPwd. ")

	desired := struct {
		Pwd1 string `json:"prh_tenant_pwd_1"`
		Pwd2 string `json:"prh_tenant_pwd_2"`
	}{
		Pwd1: "",
		Pwd2: "",
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}
