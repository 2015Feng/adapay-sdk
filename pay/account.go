package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type accountInterface interface {
	Payment(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Account struct {
	*Adapay
}

func (w *Account) Payment(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := PAGE_BASE_URL + WALLET_PAY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, w.HandleConfig(multiMerchConfigId...))
}
