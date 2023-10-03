package living

import (
	"encoding/json"
)

type ThingStatus struct {
	Status int   `json:"status"`
	Time   int64 `json:"time"`
}

func (c *Client) GetThingStatus(productKey string, deviceName string) *ThingStatus {
	log := c.log
	log.Debug("GetThingStatus. produceKey : ", productKey, " deviceName : ", deviceName)

	params := map[string]interface{}{
		"productKey": productKey,
		"deviceName": deviceName,
	}

	data, err := c.doCheckRequest("/cloud/thing/status/get", "1.0.2", params)
	if err != nil {
		log.Error("GetThingStatus. doCheckRequest fail. err : ", err)
		return nil
	}

	var ret ThingStatus
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error("GetThingStatus. json.Unmarshal fail. err : ", err)
		return nil
	}

	log.Debug("GetThingStatus. result : ", ret)
	return &ret
}
