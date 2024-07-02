package txhash

import (
	"encoding/json"
	"fmt"
)

func JsonDataProcess(jsondata string) string {
	var data map[string]interface{}

	if err := json.Unmarshal([]byte(jsondata), &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}

	types := data["types"].(map[string]interface{})

	_, hasEIP712Domain := types["EIP712Domain"]
	_, hasPrimaryType := data["primaryType"]

	//检查是否存在hasPrimaryType字段
	if !hasPrimaryType {
		var firstType string
		for key := range types {
			if key != "EIP712Domain" {
				firstType = key
				break
			}
		}
		data["primaryType"] = firstType
	}

	// 检查是否存在EIP712Domain字段
	if !hasEIP712Domain {
		eip712Domain := []map[string]string{
			{"name": "name", "type": "string"},
			{"name": "version", "type": "string"},
			{"name": "chainId", "type": "uint256"},
			{"name": "verifyingContract", "type": "address"},
		}
		types["EIP712Domain"] = eip712Domain
	}

	//检查value改名message
	//判断params是否为空，如果为空更改params的值用"0x"表示(golang特性)
	if data["value"] != nil {
		data["message"] = data["value"]
		if message, ok := data["message"].(map[string]interface{}); ok {
			if params, exists := message["params"].([]interface{}); exists && len(params) == 0 {
				message["params"] = "0x"
			}
		}
	}
	delete(data, "value")

	// 输出修改后的JSON数据
	modifiedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return ""
	}

	//fmt.Println(string(modifiedJSON))
	return string(modifiedJSON)
}
