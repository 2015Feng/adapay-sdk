package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type TransferInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
	List(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Transfer struct {
	*Adapay
}

func (t *Transfer) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + TRANSFER_CREATE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, t.HandleConfig(multiMerchConfigId...))
}

func (t *Transfer) List(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + TRANSFER_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, t.HandleConfig(multiMerchConfigId...))
}
