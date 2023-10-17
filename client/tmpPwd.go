package client

import (
	"fmt"
	"strconv"
)

type TmpPwdTimeLimit struct {
	Date    [2][3]int `json:"date"`
	Minutes [2]int    `json:"minutes"`
	WeekMap int       `json:"week_map"`
}

func (c *Client) SetTmpPassword(productKey string, deviceName string, index int, password string, cntLimit int, tmLimit *TmpPwdTimeLimit) (bool, error) {
	log := c.log
	log.Trace("SetTmpPassword. ")

	if !isPwdLegal(password) {
		log.Warnf("SetTmpPassword password illegality. password : %s", password)
		return false, fmt.Errorf("illegality password")
	}

	if index >= 1 && index <= 10 {
		desired := make(map[string]interface{})

		strIndex := strconv.Itoa(index)
		desired["prh_tmp_pwd_"+strIndex] = password
		desired["prh_tmp_pwd_tm_limit_"+strIndex] = tmLimit
		desired["prh_tmp_pwd_cnt_limit_"+strIndex] = cntLimit

		return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
	}

	log.Warnf("SetTmpPassword. unsupported index, [1, 10] index : ", index)
	return false, fmt.Errorf("unsupported index, [1, 10] index : ", index)
}
