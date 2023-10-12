package client

type TmpPwdTimeLimit struct {
	Date    [2][3]int `json:"date"`
	Minutes [2]int    `json:"minutes"`
	WeekMap int       `json:"week_map"`
}

func (c *Client) SetTmpPwd1(productKey string, deviceName string, pwd string, cntLimit int, tmLimit *TmpPwdTimeLimit) bool {
	log := c.log
	log.Trace("SetTmpPwd1. ")

	desired := struct {
		Pwd      string      `json:"prh_tmp_pwd_1"`
		TmLimit  interface{} `json:"prh_tmp_pwd_tm_limit_1"`
		CntLimit int         `json:"prh_tmp_pwd_cnt_limit_1"`
	}{
		Pwd:      pwd,
		TmLimit:  tmLimit,
		CntLimit: cntLimit,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}

func (c *Client) SetTmpPwd2(productKey string, deviceName string, pwd string, cntLimit int, tmLimit *TmpPwdTimeLimit) bool {
	log := c.log
	log.Trace("SetTmpPwd2. ")

	desired := struct {
		Pwd      string      `json:"prh_tmp_pwd_2"`
		TmLimit  interface{} `json:"prh_tmp_pwd_tm_limit_2"`
		CntLimit int         `json:"prh_tmp_pwd_cnt_limit_2"`
	}{
		Pwd:      pwd,
		TmLimit:  tmLimit,
		CntLimit: cntLimit,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}
