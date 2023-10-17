package client

type TmpPwdTimeLimit struct {
	Date    [2][3]int `json:"date"`
	Minutes [2]int    `json:"minutes"`
	WeekMap int       `json:"week_map"`
}

func (c *Client) SetTmpPassword1(productKey string, deviceName string, password string, cntLimit int, tmLimit *TmpPwdTimeLimit) (bool, error) {
	log := c.log
	log.Trace("SetTmpPassword1. ")

	desired := struct {
		Pwd      string      `json:"prh_tmp_pwd_1"`
		TmLimit  interface{} `json:"prh_tmp_pwd_tm_limit_1"`
		CntLimit int         `json:"prh_tmp_pwd_cnt_limit_1"`
	}{
		Pwd:      password,
		TmLimit:  tmLimit,
		CntLimit: cntLimit,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}

func (c *Client) SetTmpPassword2(productKey string, deviceName string, password string, cntLimit int, tmLimit *TmpPwdTimeLimit) (bool, error) {
	log := c.log
	log.Trace("SetTmpPassword2. ")

	desired := struct {
		Pwd      string      `json:"prh_tmp_pwd_2"`
		TmLimit  interface{} `json:"prh_tmp_pwd_tm_limit_2"`
		CntLimit int         `json:"prh_tmp_pwd_cnt_limit_2"`
	}{
		Pwd:      password,
		TmLimit:  tmLimit,
		CntLimit: cntLimit,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}
