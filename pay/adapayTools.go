package adapay

import (
	adapayCore "adapay-sdk/adapay-core"
)

type adapayToolsInterface interface {
	DownloadBill(bill_date string, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	UserIdentity(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	VerifySign(signData string, originalData string, multiMerchConfigId ...string) error
}

type AdapayTools struct {
	*Adapay
}

func (b *AdapayTools) DownloadBill(bill_date string, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + BILL_DOWNLOAD

	reqParam := make(map[string]interface{})
	reqParam["bill_date"] = bill_date

	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, b.HandleConfig(multiMerchConfigId...))
}

func (u *AdapayTools) UserIdentity(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + USER_IDENTITY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, u.HandleConfig(multiMerchConfigId...))
}

func (v *AdapayTools) VerifySign(signData string, originalData string, multiMerchConfigId ...string) error {

	return adapayCore.RsaSignVerify(signData, originalData, v.HandleConfig(multiMerchConfigId...))
}
