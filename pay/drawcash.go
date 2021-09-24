package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type drawcashInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Drawcash struct {
	*Adapay
}

func (c *Drawcash) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + BILL_DOWNLOAD
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, c.HandleConfig(multiMerchConfigId...))
}

func (c *Drawcash) Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + QUERY_CASHS_STAT
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, c.HandleConfig(multiMerchConfigId...))
}
