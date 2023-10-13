package living

import (
	"encoding/json"
	"fmt"
)

type ThingStatus struct {
	Status int   `json:"status"`
	Time   int64 `json:"time"`
}

func (c *Client) GetThingStatus(productKey string, deviceName string) (*ThingStatus, error) {
	log := c.log
	log.Trace("GetThingStatus. produceKey : ", productKey, " deviceName : ", deviceName)

	params := map[string]interface{}{
		"productKey": productKey,
		"deviceName": deviceName,
	}

	data, err := c.doCheckRequest("/cloud/thing/status/get", "1.0.2", params)
	if err != nil {
		log.Error("GetThingStatus. doCheckRequest fail. err : ", err)
		return nil, err
	}

	var ret ThingStatus
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error("GetThingStatus. json.Unmarshal fail. err : ", err)
		return nil, fmt.Errorf("json.Unmarshal fail")
	}

	log.Trace("GetThingStatus. result : ", ret)
	return &ret, nil
}

type ThingProperty struct {
	Attribute   string          `json:"attribute"`
	BatchId     string          `json:"batchId"`
	GmtModified int64           `json:"gmtModified"`
	IotId       string          `json:"iotId"`
	Value       json.RawMessage `json:"value"`
}

func (c *Client) GetThingProperties(productKey string, deviceName string) (*[]ThingProperty, error) {
	log := c.log
	log.Trace("GetThingProperties. produceKey : ", productKey, " deviceName : ", deviceName)

	params := map[string]interface{}{
		"productKey": productKey,
		"deviceName": deviceName,
	}

	data, err := c.doCheckRequest("/cloud/thing/properties/get", "1.0.2", params)
	if err != nil {
		log.Error("GetThingProperties. doCheckRequest fail. err : ", err)
		return nil, err
	}

	var ret []ThingProperty
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error("GetThingProperties. json.Unmarshal fail. err : ", err)
		return nil, fmt.Errorf("json.Unmarshal fail")
	}

	log.Trace("GetThingProperties. result : ", ret)
	return &ret, nil
}
