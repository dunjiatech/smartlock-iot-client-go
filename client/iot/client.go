package iot

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	iot20180120 "github.com/alibabacloud-go/iot-20180120/v5/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/sirupsen/logrus"
)

type Client struct {
	aliIotClient *iot20180120.Client
	log          *logrus.Entry
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createAliIotClient(accessKeyId *string, accessKeySecret *string) (_result *iot20180120.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Iot
	config.Endpoint = tea.String("iot.cn-shanghai.aliyuncs.com")
	// _result = &iot20180120.Client{}
	_result, _err = iot20180120.NewClient(config)
	return _result, _err
}

func CreateClient(accessKeyId string, accessKeySecret string) (*Client, error) {
	logrus.Trace("DJIOT CreateClient. accessKeyId : ", accessKeyId, " accessKeySecret : ", accessKeySecret)

	aliIotClient, err := createAliIotClient(&accessKeyId, &accessKeySecret)
	if err != nil {
		logrus.Error("DJIOT CreateClient fail. err : ", err, " accessKeyId : ", accessKeyId, " accessKeySecret : ", accessKeySecret)
		return nil, err
	}

	log := logrus.WithFields(logrus.Fields{
		"Module":          "DJIOT",
		"AccessKeyId":     accessKeyId,
		"AccessKeySecret": accessKeySecret,
	})

	client := Client{
		aliIotClient: aliIotClient,
		log:          log,
	}
	return &client, nil
}
