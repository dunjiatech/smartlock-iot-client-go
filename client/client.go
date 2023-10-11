package client

import (
	"smartlock-iot-client/client/iot"
	"smartlock-iot-client/client/living"

	"github.com/sirupsen/logrus"
)

type Client struct {
	IotClient    *iot.Client
	LivingClient *living.Client
	log          *logrus.Entry
}

func CreateClient(accessKeyId *string, accessKeySecret *string, appKey *string, appSecret *string, projectId *string) *Client {
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
