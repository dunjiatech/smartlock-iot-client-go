package client

import "fmt"

func (c *Client) SetAdminPassword(productKey string, deviceName string, password string) (bool, error) {
	log := c.log
	log.Trace("SetAdminPassword. ")

	if !isPwdLegal(password) {
		log.Warnf("SetAdminPassword password illegality. password : %s", password)
		return false, fmt.Errorf("illegality password")
	}

	desired := struct {
		Pwd string `json:"prh_admin_pwd_1"`
	}{
		Pwd: password,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}
