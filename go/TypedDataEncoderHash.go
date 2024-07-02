package txhash

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func typedDataEncoderHash(data apitypes.TypedData) (string, error) {
	domain, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return "", fmt.Errorf("domain hash error: %v", err)
	}
	dataHash, err := data.HashStruct(data.PrimaryType, data.Message)
	if err != nil {
		return "", fmt.Errorf("data hash error: %v", err)
	}

	//fmt.Printf("domain: %s\ndataHash: %s\n", common.Bytes2Hex(domain), common.Bytes2Hex(dataHash))

	prefixedData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domain), string(dataHash)))
	prefixedDataHash := crypto.Keccak256(prefixedData)
	return "0x" + common.Bytes2Hex(prefixedDataHash), nil
}

func TypedDataEncoderHash(jsonData string) string {
	var typedData apitypes.TypedData
	jData := JsonDataProcess(jsonData)
	if err := json.Unmarshal([]byte(jData), &typedData); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}

	typeencodehash, err := typedDataEncoderHash(typedData)
	if err != nil {
		log.Fatalf("hash error: %v", err)
	}

	//fmt.Println(data)
	return typeencodehash

}
