package types

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// This is introduced because of the Tendermint IAVL Merkle Proof verification exploitation.
var NanoBlackList = []common.Address{
	common.HexToAddress("0x489A8756C18C0b8B24EC2a2b9FF3D4d447F79BEc"),
	common.HexToAddress("0xFd6042Df3D74ce9959922FeC559d7995F3933c55"),
	// Test Account
	common.HexToAddress("0xdb789Eb5BDb4E559beD199B8b82dED94e1d056C9"),
}

// DynamicBlackList holds addresses loaded from external config file.
var DynamicBlackList []common.Address

type blacklistFile struct {
	Addresses []string `json:"addresses"`
}

// LoadBlacklistFile reads a JSON file and populates DynamicBlackList.
func LoadBlacklistFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var file blacklistFile
	if err := json.Unmarshal(data, &file); err != nil {
		return err
	}
	for _, addrStr := range file.Addresses {
		addr := common.HexToAddress(strings.TrimSpace(addrStr))
		DynamicBlackList = append(DynamicBlackList, addr)
	}
	log.Info("Loaded dynamic blacklist", "file", path, "count", len(DynamicBlackList))
	return nil
}

// IsBlacklisted checks if an address is in NanoBlackList or DynamicBlackList.
func IsBlacklisted(addr common.Address) bool {
	for _, blackAddr := range NanoBlackList {
		if addr == blackAddr {
			return true
		}
	}
	for _, blackAddr := range DynamicBlackList {
		if addr == blackAddr {
			return true
		}
	}
	return false
}
