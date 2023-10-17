package client

import "fmt"

func (c *Client) SetTenantPassword(productKey string, deviceName string, index int, passsword string) (bool, error) {
	log := c.log
	log.Trace("SetTenantPassword. index : ", index, " password: ", passsword)

	switch index {
	case 1:
		return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, struct {
			Pwd string `json:"prh_tenant_pwd_1"`
		}{Pwd: passsword}, true)
	case 2:
		return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, struct {
			Pwd string `json:"prh_tenant_pwd_2"`
		}{Pwd: passsword}, true)
	}

	return false, fmt.Errorf("unsupported index, index : ", index)
}

func (c *Client) ClearAllTenantPassword(productKey string, deviceName string) (bool, error) {
	log := c.log
	log.Trace("ClearAllTenantPwd. ")

	desired := struct {
		Pwd1 string `json:"prh_tenant_pwd_1"`
		Pwd2 string `json:"prh_tenant_pwd_2"`
	}{
		Pwd1: "-",
		Pwd2: "-",
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}

func (c *Client) TenantCheckIn(productKey string, deviceName string, tenantId string, passsword string) (bool, error) {
	log := c.log
	log.Trace("TenantCheckIn. ")

	desired := struct {
		TenantId string `json:"prh_tenant_id"`
		Pwd1     string `json:"prh_tenant_pwd_1"`
	}{
		TenantId: tenantId,
		Pwd1:     passsword,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}

func (c *Client) TenantCheckOut(productKey string, deviceName string) (bool, error) {
	log := c.log
	log.Trace("TenantCheckOut. ")

	desired := struct {
		TenantId string `json:"prh_tenant_id"`
	}{
		TenantId: "-",
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}
