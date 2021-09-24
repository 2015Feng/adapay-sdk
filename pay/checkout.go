package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type checkoutInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Checkout struct {
	*Adapay
}

func (w *Checkout) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := PAGE_BASE_URL + WALLET_CHECKOUT
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, w.HandleConfig(multiMerchConfigId...))
}

func (m *Checkout) QueryList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + QUERY_CHECKOUT_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, m.HandleConfig(multiMerchConfigId...))
}
