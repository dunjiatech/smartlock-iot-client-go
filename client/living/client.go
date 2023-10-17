package living

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	aliIot "github.com/alibabacloud-go/iot-api-gateway/client"
	aliUtil "github.com/alibabacloud-go/tea-utils/service"
	aliTea "github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	projectId string
	aliClient *aliIot.Client
	token     struct {
		CloudToken string
		ExpiresIn  int64
		BeginTime  int64
	}
	log *logrus.Entry
}

func CreateClient(appKey string, appSecret string, projectId string) (*Client, error) {
	logrus.Trace("DJLIVING CreateClient. appKey : ", appKey, " appSecret : ", appSecret, " projectId : ", projectId)

	aliConfig := new(aliIot.Config).
		SetAppKey(appKey).
		SetAppSecret(appSecret).
		SetDomain("api.link.aliyun.com")

	aliClient, err := aliIot.NewClient(aliConfig)
	if err != nil {
		logrus.Error("DJLIVING aliIot.NewClient fail. err : ", err)
		return nil, err
	}

	log := logrus.WithFields(logrus.Fields{
		"Module":    "DJLIVING",
		"AppKey":    appKey,
		"AppSecret": appSecret,
		"ProjectId": projectId,
	})

	client := Client{
		projectId: projectId,
		aliClient: aliClient,
		log:       log,
	}
	return &client, nil
}

func (c *Client) checkToken() error {
	log := c.log
	log.Trace("checkToken")

	token := &c.token
	curMilli := time.Now().UnixMilli()
	if (curMilli - token.BeginTime) > token.ExpiresIn/2 {
		ret, err := c.getToken()
		if ret == nil {
			return err
		}

		token.CloudToken = ret.CloudToken
		token.ExpiresIn = ret.ExpireIn
		token.BeginTime = curMilli
	}
	return nil
}

func (c *Client) doCheckRequest(pathname string, apiVer string, params map[string]interface{}) ([]byte, error) {
	log := c.log
	log.Trace("doCheckRequest pathname : ", pathname, " apiVer : ", apiVer, " params : ", params)

	err := c.checkToken()
	if err != nil {
		return nil, err
	}
	return c.doRequest(pathname, apiVer, c.token.CloudToken, params)
}

type Response struct {
	Code         int             `json:"code"`
	Message      string          `json:"message"`
	LocalizedMsg string          `json:"localizedMsg"`
	Data         json.RawMessage `json:"data"`
	Id           string          `json:"id"`
}

func (c *Client) doRequest(pathname string, apiVer string, cloudToken string, params map[string]interface{}) ([]byte, error) {
	aliClient := c.aliClient
	log := c.log

	log.Trace("doRequest pathname : ", pathname, " apiVer : ", apiVer, " cloudToken : ", cloudToken, " params : ", params)

	req := new(aliIot.CommonParams).SetApiVer(apiVer)
	if len(cloudToken) != 0 {
		req.SetCloudToken(cloudToken)
	}
	body := new(aliIot.IoTApiRequest).SetParams(params).SetRequest(req)
	runtime := new(aliUtil.RuntimeOptions)

	resp, err := aliClient.DoRequest(aliTea.String(pathname), aliTea.String("HTTPS"), aliTea.String("POST"), nil, body, runtime)
	if err != nil {
		log.Error("doRequest aliClient.DoRequest fail. err : ", err)
		return nil, err
	}

	var respBody []byte
	respBody, err = resp.ReadBody()
	if err != nil {
		log.Error("doRequest resp.ReadBody fail. err : ", err)
		return nil, err
	}

	var ret Response
	err = json.Unmarshal(respBody, &ret)
	if err != nil {
		log.Error("doRequest json.Unmarshal fail. err : ", err)
		return nil, err
	}

	if ret.Code != 200 {
		log.Error("doRequest rerv fail. ret : ", ret)
		return nil, fmt.Errorf("", ret)
	}

	return ret.Data, nil
}
