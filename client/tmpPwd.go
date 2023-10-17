package client

import "fmt"

type TmpPwdTimeLimit struct {
	Date    [2][3]int `json:"date"`
	Minutes [2]int    `json:"minutes"`
	WeekMap int       `json:"week_map"`
}

func (c *Client) SetTmpPassword(productKey string, deviceName string, index int, password string, cntLimit int, tmLimit *TmpPwdTimeLimit) (bool, error) {
	log := c.log
	log.Trace("SetTmpPassword. ")

	switch index {
	case 1:
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
	case 2:
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

	return false, fmt.Errorf("unsupported index, index : ", index)
}
