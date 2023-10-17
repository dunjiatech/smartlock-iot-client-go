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

func CreateClient(accessKeyId string, accessKeySecret string, appKey string, appSecret string, projectId string) (*Client, error) {
	logrus.Trace("DJSDK CreateClient. ")

	iotClient, err0 := iot.CreateClient(accessKeyId, accessKeySecret)
	if iotClient == nil {
		return nil, err0
	}
	livingClient, err1 := living.CreateClient(appKey, appSecret, projectId)
	if livingClient == nil {
		return nil, err1
	}

	log := logrus.WithFields(logrus.Fields{
		"Module": "DJSDK",
	})

	client := Client{
		IotClient:    iotClient,
		LivingClient: livingClient,
		log:          log,
	}

	return &client, nil
}
