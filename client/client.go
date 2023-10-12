package client

import (
	"github.com/dunjiatech/smartlock-iot-client-go/client/iot"
	"github.com/dunjiatech/smartlock-iot-client-go/client/living"

	"github.com/sirupsen/logrus"
)

type Client struct {
	IotClient    *iot.Client
	LivingClient *living.Client
	log          *logrus.Entry
}

func CreateClient(accessKeyId string, accessKeySecret string, appKey string, appSecret string, projectId string) *Client {
	logrus.Trace("DJSDK CreateClient. ")

	iotClient := iot.CreateClient(accessKeyId, accessKeySecret)
	if iotClient == nil {
		return nil
	}
	livingClient := living.CreateClient(appKey, appSecret, projectId)
	if livingClient == nil {
		return nil
	}

	log := logrus.WithFields(logrus.Fields{
		"Module": "DJSDK",
	})

	client := Client{
		IotClient:    iotClient,
		LivingClient: livingClient,
		log:          log,
	}

	return &client
}

func (c *Client) SetOpenMode(productKey string, deviceName string, mode int) bool {
	log := c.log
	log.Trace("SetOpenMode mode : ", mode)

	desired := struct {
		OpenMode int `json:"open_mode"`
	}{
		OpenMode: mode,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}

func (c *Client) SetAdminPassword(productKey string, deviceName string, passsword string) bool {
	log := c.log
	log.Trace("SetAdminPassword. ")

	desired := struct {
		Password string `json:"prh_admin_pwd_1"`
	}{
		Password: passsword,
	}

	return c.IotClient.UpdateDeviceShadowEx(productKey, deviceName, desired, true)
}
