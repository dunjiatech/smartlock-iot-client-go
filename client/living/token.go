package living

import (
	"encoding/json"
)

type Token struct {
	CloudToken string `json:"cloudToken"`
	ExpireIn   int64  `json:"expireIn"`
}

func (c *Client) getToken() (*Token, error) {
	log := c.log
	log.Trace("getToken")

	params := map[string]interface{}{
		"grantType": "project",
		"res":       c.projectId,
	}

	data, err := c.doRequest("/cloud/token", "1.0.1", "", params)
	if err != nil {
		return nil, err
	}

	var ret Token
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error("getToken. json.Unmarshal fail. err : ", err)
		return nil, err
	}

	log.Debug("getToken. result : ", ret)
	return &ret, nil
}
