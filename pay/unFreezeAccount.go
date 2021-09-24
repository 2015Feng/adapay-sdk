package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type UnFreezeAccountInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
	List(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type UnFreezeAccount struct {
	*Adapay
}

func (f *UnFreezeAccount) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + UnFREEZE_ACCOUNT_FREEZE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, f.HandleConfig(multiMerchConfigId...))
}

func (f *UnFreezeAccount) List(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + UnFREEZE_ACCOUNT_FREEZE_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, f.HandleConfig(multiMerchConfigId...))
}
