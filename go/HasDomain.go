package txhash

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func hashDomain(data apitypes.TypedData) (string, error) {
	domain, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return "", err
	}

	return "0x" + common.Bytes2Hex(domain), err
}

func HashDomain(jsonData string) string {
	var typedData apitypes.TypedData
	jData := JsonDataProcess(jsonData)
	if err := json.Unmarshal([]byte(jData), &typedData); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}

	domain, err := hashDomain(typedData)
	if err != nil {
		fmt.Println(err)
	}

	return domain
}
