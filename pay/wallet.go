package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type walletInterface interface {
	Login(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Wallet struct {
	*Adapay
}

func (w *Wallet) Login(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := PAGE_BASE_URL + WALLET_LOGIN
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, w.HandleConfig(multiMerchConfigId...))
}
