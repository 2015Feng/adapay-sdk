package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type FreezeAccountInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
	List(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type FreezeAccount struct {
	*Adapay
}

func (f *FreezeAccount) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + FREEZE_ACCOUNT_FREEZE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, f.HandleConfig(multiMerchConfigId...))
}

func (f *FreezeAccount) List(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + FREEZE_ACCOUNT_FREEZE_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, f.HandleConfig(multiMerchConfigId...))
}
