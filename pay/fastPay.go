package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type fastPayInterface interface {
	CardBind(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
	CardBindConfirm(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
	CardBindlist(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
	Confirm(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
	SmsCode(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type FastPay struct {
	*Adapay
}

func (f *FastPay) CardBind(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := PAGE_BASE_URL + FAST_CARD_APPLY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, f.HandleConfig(multiMerchConfigId...))
}

func (f *FastPay) CardBindConfirm(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := PAGE_BASE_URL + FAST_CARD_CONFIRM
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, f.HandleConfig(multiMerchConfigId...))
}

func (f *FastPay) CardBindlist(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := PAGE_BASE_URL + FAST_CARD_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, f.HandleConfig(multiMerchConfigId...))
}

func (f *FastPay) Confirm(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + FAST_PAY_CONFIRM
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, f.HandleConfig(multiMerchConfigId...))
}

func (f *FastPay) SmsCode(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + FAST_PAY_SMS_CODE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, f.HandleConfig(multiMerchConfigId...))
}
