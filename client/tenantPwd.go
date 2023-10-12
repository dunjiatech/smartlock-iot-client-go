package client

type TenantPwd1 struct {
	Pwd string `json:"prh_tenant_pwd_1"`
}

type TenantPwd2 struct {
	Pwd string `json:"prh_tenant_pwd_2"`
}

type TenantPwd3 struct {
	Pwd string `json:"prh_tenant_pwd_3"`
}

func (c *Client) SetTenantPassword(productKey *string, deviceName *string, index int, passsword string) bool {
	log := c.log
	log.Trace("SetTenantPassword. ")

	switch index {
	case 1:
		return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, TenantPwd1{Pwd: passsword}, true)
	case 2:
		return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, TenantPwd2{Pwd: passsword}, true)
	case 3:
		return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, TenantPwd3{Pwd: passsword}, true)
	default:
		log.Warn("Unknow Tenant index : ", index)
	}

	return false
}
