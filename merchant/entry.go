package merchant

import (
	adapayCore "../adapay-core"
)

type entryInterface interface {
	BatchEntrys(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryEntry(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Entry struct {
	*Merchant
}

func (e *Entry) BatchEntrys(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + BATCH_ENTRYS
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, e.HandleConfig(multiMerchConfigId...))
}

func (e *Entry) QueryEntry(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + QUERY_ENTRY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, e.HandleConfig(multiMerchConfigId...))
}
