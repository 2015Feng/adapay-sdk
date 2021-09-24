package merchant

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	adapayCore "../adapay-core"
)

type Merchant struct {

	MultiMerchSysConfigs map[string]*adapayCore.MerchSysConfig

	DefaultMerchSysConfig *adapayCore.MerchSysConfig
}


func InitDefaultMerchSysConfig(filePath string) (*Merchant, error) {

	config, err := adapayCore.ReadMerchConfig(filePath)
	if err != nil {
		return nil, err
	}

	ada := &Merchant{
		DefaultMerchSysConfig: config,
	}

	return ada, nil
}


func InitMultiMerchSysConfigs(fileDir string) (*Merchant, error) {

	dirs, _ := ioutil.ReadDir(fileDir)

	configs := map[string]*adapayCore.MerchSysConfig{}

	for _, f := range dirs {

		ext := filepath.Ext(f.Name())
		if ext == ".json" {
			config, err := adapayCore.ReadMerchConfig(fileDir + f.Name())
			if err != nil {
				continue
			}

			key := strings.Replace(f.Name(), ".json", "", -1)
			configs[key] = config
		}
	}

	ada := &Merchant{
		MultiMerchSysConfigs: configs,
	}

	return ada, nil
}


func (a *Merchant) HandleConfig(multiMerchConfigId ...string) *adapayCore.MerchSysConfig {
	if multiMerchConfigId == nil {
		return a.DefaultMerchSysConfig
	} else {
		return a.MultiMerchSysConfigs[multiMerchConfigId[0]]
	}
}


func (m *Merchant) BatchInput() *BatchInput {
	return &BatchInput{Merchant: m}
}


func (a *Merchant) Entry() *Entry {
	return &Entry{Merchant: a}
}


func (a *Merchant) MerProfile() *MerProfile {
	return &MerProfile{Merchant: a}
}


func (a *Merchant) Version() string {
	return "1.0.0"
}
