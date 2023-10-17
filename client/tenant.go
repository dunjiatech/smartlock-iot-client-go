package client

import (
	"fmt"
	"strconv"
)

const maxTenantPwdIndex int = 20

func (c *Client) SetTenantPassword(productKey string, deviceName string, index int, password string) (bool, error) {
	log := c.log
	log.Trace("SetTenantPassword. index : ", index, " password: ", password)

	if index >= 1 && index <= maxTenantPwdIndex {
		desired := make(map[string]interface{})

		strIndex := strconv.Itoa(index)
		desired["prh_tenant_pwd_"+strIndex] = password

		return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
	}

	log.Warnf("SetTenantPassword. unsupported index, [1, %d] index : %d", maxTenantPwdIndex, index)
	return false, fmt.Errorf("unsupported index, [1, %d] index : %d", maxTenantPwdIndex, index)
}

func (c *Client) ClearAllTenantPassword(productKey string, deviceName string) (bool, error) {
	log := c.log
	log.Trace("ClearAllTenantPwd. ")

	desired := make(map[string]interface{})
	for i := 1; i <= maxTenantPwdIndex; i++ {
		s := strconv.Itoa(i)
		desired["prh_tenant_pwd_"+s] = "-"
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
