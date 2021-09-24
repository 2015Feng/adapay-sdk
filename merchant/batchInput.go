package merchant

import (
	adapayCore "adapay-sdk/adapay-core"
)

type batchInputInterface interface {
	MerConf(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryMerConf(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	MerResidentModify(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type BatchInput struct {
	*Merchant
}

func (b *BatchInput) MerConf(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + MER_CONF
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, b.HandleConfig(multiMerchConfigId...))
}

func (b *BatchInput) QueryMerConf(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + QUERY_MER_CONF
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, b.HandleConfig(multiMerchConfigId...))
}

func (b *BatchInput) MerResidentModify(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + MER_RESIDENT_MODIFY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, b.HandleConfig(multiMerchConfigId...))
}
