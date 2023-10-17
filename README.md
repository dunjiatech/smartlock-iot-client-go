

# 遁甲门锁服务API参考

## type Client

```
type Client struct {
	IotClient    *iot.Client
	LivingClient *living.Client
	log          *logrus.Entry
}
```

### func CreateClient

```
func CreateClient(accessKeyId string, accessKeySecret string, appKey string, appSecret string, projectId string) (*Client, error) 
```

创建门锁客户端

**请求参数**

| 名称            | 类型   | 描述             |
| --------------- | ------ | :--------------- |
| accessKeyId     | string | 联系遁甲人员获取 |
| accessKeySecret | string | 联系遁甲人员获取 |
| appKey          | string | 联系遁甲人员获取 |
| appSecret       | string | 联系遁甲人员获取 |
| projectId       | string | 联系遁甲人员获取 |

**返回数据**

| 名称       | 类型    | 描述 |
| ---------- | ------- | :--- |
| 门锁客户端 | *Client |      |
| 失败信息   | error   |      |



### func (*Client) SetOpenMode

```
func (c *Client) SetOpenMode(productKey string, deviceName string, mode int) (bool, error)
```

调用该接口设置门锁的开门模式

**请求参数**

| 名称       | 类型   | 描述                                               |
| ---------- | ------ | :------------------------------------------------- |
| productKey | string | 产品id                                             |
| deviceName | string | 设备id                                             |
| mode       | int    | 开门模式      <br />1：普通模式；<br />2：常开模式 |

**返回数据**

| 名称     | 类型  | 描述                                  |
| -------- | ----- | :------------------------------------ |
| 是否成功 | bool  | true : 设置成功<br />false : 设置失败 |
| 失败信息 | error |                                       |



### func (*Client) GetLockStatus

```
func (c *Client) GetLockStatus(productKey string, deviceName string) (int, error)
```

调用该接口获取门锁状态

**请求参数**

| 名称       | 类型   | 描述   |
| ---------- | ------ | :----- |
| productKey | string | 产品id |
| deviceName | string | 设备id |

**返回数据**

| 名称     | 类型  | 描述                                                         |
| -------- | ----- | :----------------------------------------------------------- |
| 门锁状态 | int   | 0 : 未激活<br />1 : 表示在线<br />3 : 表示离线<br />8 : 表示禁用 |
| 失败信息 | error |                                                              |



### func (*Client) GetBattery0Status

```
func (c *Client) GetBattery0Status(productKey string, deviceName string) (*BatteryStatus, error) 
```

调用该接口获取0号电池状态

**请求参数**

| 名称       | 类型   | 描述   |
| ---------- | ------ | :----- |
| productKey | string | 产品id |
| deviceName | string | 设备id |

**返回数据**

| 名称     | 类型           | 描述             |
| -------- | -------------- | :--------------- |
| 电池状态 | *BatteryStatus | 电池状态描述如下 |
| 失败信息 | error          |                  |



### func (*Client) SetAdminPassword

```
func (c *Client) SetAdminPassword(productKey string, deviceName string, passsword string) (bool, error) 
```

设置管理员密码

**请求参数**

| 名称       | 类型   | 描述                                          |
| ---------- | ------ | :-------------------------------------------- |
| productKey | string | 产品id                                        |
| deviceName | string | 设备id                                        |
| passsword  | string | 密码只能传入6～12位的数字字符串，如：“123456” |

**返回数据**

| 名称         | 类型  | 描述                          |
| ------------ | ----- | :---------------------------- |
| 是否设置成功 | bool  | true : 成功<br />false : 失败 |
| 失败信息     | error |                               |



### func (*Client) TenantCheckIn

```
func (c *Client) TenantCheckIn(productKey string, deviceName string, tenantId string, passsword string) (bool, error)
```

租户入住

**请求参数**

| 名称       | 类型   | 描述                                          |
| ---------- | ------ | :-------------------------------------------- |
| productKey | string | 产品id                                        |
| deviceName | string | 设备id                                        |
| tenantId   | string | 租户id                                        |
| passsword  | string | 密码只能传入6～12位的数字字符串，如：“123456” |

**返回数据**

| 名称         | 类型  | 描述                          |
| ------------ | ----- | :---------------------------- |
| 是否设置成功 | bool  | true : 成功<br />false : 失败 |
| 失败信息     | error |                               |



### func (*Client) TenantCheckOut

```
func (c *Client) TenantCheckOut(productKey string, deviceName string) (bool, error)
```

租户退租

**请求参数**

| 名称       | 类型   | 描述   |
| ---------- | ------ | :----- |
| productKey | string | 产品id |
| deviceName | string | 设备id |

**返回数据**

| 名称         | 类型  | 描述                          |
| ------------ | ----- | :---------------------------- |
| 是否设置成功 | bool  | true : 成功<br />false : 失败 |
| 失败信息     | error |                               |



### func (*Client) SetTenantPassword

```
func (c *Client) SetTenantPassword(productKey string, deviceName string, index int, passsword string) (bool, error)
```

设置租户密码

**请求参数**

| 名称       | 类型   | 描述                                          |
| ---------- | ------ | :-------------------------------------------- |
| productKey | string | 产品id                                        |
| deviceName | string | 设备id                                        |
| index      | int    | 密码序号。目前支持传入：1、2                  |
| passsword  | string | 密码只能传入6～12位的数字字符串，如：“123456” |

**返回数据**

| 名称         | 类型  | 描述                          |
| ------------ | ----- | :---------------------------- |
| 是否设置成功 | bool  | true : 成功<br />false : 失败 |
| 失败信息     | error |                               |



### func (*Client) ClearAllTenantPassword

```
func (c *Client) ClearAllTenantPassword(productKey string, deviceName string) (bool, error)
```

清除所有的租户密码

**请求参数**

| 名称       | 类型   | 描述   |
| ---------- | ------ | :----- |
| productKey | string | 产品id |
| deviceName | string | 设备id |

**返回数据**

| 名称         | 类型  | 描述                          |
| ------------ | ----- | :---------------------------- |
| 是否设置成功 | bool  | true : 成功<br />false : 失败 |
| 失败信息     | error |                               |



### func (*Client) SetTmpPassword1

```
func (c *Client) SetTmpPassword1(productKey string, deviceName string, password string, cntLimit int, tmLimit *TmpPwdTimeLimit) (bool, error) 
```

设置1号临时密码

**请求参数**

| 名称       | 类型             | 描述                                             |
| ---------- | ---------------- | :----------------------------------------------- |
| productKey | string           | 产品id                                           |
| deviceName | string           | 设备id                                           |
| password   | string           | 密码只能传入6～12位的数字字符串，如：“123456”    |
| cntLimit   | int              | 密码使用次数                                     |
| tmLimit    | *TmpPwdTimeLimit | 见《type TmpPwdTimeLimit》章节，密码使用时间限制 |

**返回数据**

| 名称         | 类型  | 描述                          |
| ------------ | ----- | :---------------------------- |
| 是否设置成功 | bool  | true : 成功<br />false : 失败 |
| 失败信息     | error |                               |



### func (*Client) SetTmpPassword2

```
func (c *Client) SetTmpPassword2(productKey string, deviceName string, password string, cntLimit int, tmLimit *TmpPwdTimeLimit) (bool, error) 
```

设置2号临时密码

**请求参数**

| 名称       | 类型             | 描述                                             |
| ---------- | ---------------- | :----------------------------------------------- |
| productKey | string           | 产品id                                           |
| deviceName | string           | 设备id                                           |
| password   | string           | 密码只能传入6～12位的数字字符串，如：“123456”    |
| cntLimit   | int              | 密码使用次数                                     |
| tmLimit    | *TmpPwdTimeLimit | 见《type TmpPwdTimeLimit》章节，密码使用时间限制 |

**返回数据**

| 名称         | 类型  | 描述                          |
| ------------ | ----- | :---------------------------- |
| 是否设置成功 | bool  | true : 成功<br />false : 失败 |
| 失败信息     | error |                               |



## type BatteryStatus

电池状态

```
type BatteryStatus struct {
	Inplace bool		// 是否在位。 true:在位；false:不在位
	Voltage int     // 电量百分比
}
```



## type TmpPwdTimeLimit

临时密码时间限制

```
type TmpPwdTimeLimit struct {
	Date    [2][3]int `json:"date"`					// 开始结束年月日，如：{{2023, 10, 10}, {2023, 10, 11}}
	Minutes [2]int    `json:"minutes"`      // 一天的分钟数，{0, 1439} ： 全天有效
	WeekMap int       `json:"week_map"`     // 星期位图，127 ： 7天有效
}
```

