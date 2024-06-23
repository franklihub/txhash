package txhash

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func hashDomain(data apitypes.TypedData) (string, error) {
	domain, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return "", err
	}

	return domain.String(), err
}

func HashDomain(jsonData string) string {
	var typedData apitypes.TypedData
	if err := json.Unmarshal([]byte(jsonData), &typedData); err != nil {
		panic(err)
	}

	data, err := hashDomain(typedData)
	if err != nil {
		fmt.Println(err)
	}

	return data
}
