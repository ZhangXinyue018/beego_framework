package node

import (
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"errors"
	"time"
)

type NodeRpc struct {
	EosNodeUrl              string
	EosWalletUrl            string
	UotcEosAccount          string
	EosCryptoName           string
	UotcEosAccountActiveKey string
	UotcWallet              string
	UotcWalletPassword      string
}

type AbiActionType int

const (
	SYSTEM_NEW_ACCOUNT = AbiActionType(0)
	SYSTEM_BUY_RAM     = AbiActionType(1)
	SYSTEM_DELEGATE_BW = AbiActionType(2)
)

type EosRequestType string

const (
	CHAIN_GET_INFO          = EosRequestType("/v1/chain/get_info")
	CHAIN_GET_BLOCK         = EosRequestType("/v1/chain/get_block")
	CHAIN_GET_TABLE_ROWS    = EosRequestType("/v1/chain/get_table_rows")
	CHAIN_GET_ABI           = EosRequestType("/v1/chain/get_abi")
	CHAIN_ABI_JSON_TO_BIN   = EosRequestType("/v1/chain/abi_json_to_bin")
	CHAIN_PUSH_TRANSACTION  = EosRequestType("/v1/chain/push_transaction")
	CHAIN_GET_ACCOUNT       = EosRequestType("/v1/chain/get_account")
	WALLET_UNLOCK           = EosRequestType("/v1/wallet/unlock")
	WALLET_SIGN_TRANSACTION = EosRequestType("/v1/wallet/sign_transaction")
)

const EOS_SYSTEM_ACCOUNT = "eosio"

type GetInfoResp struct {
	ServerVersion            string `json:"server_version"`
	ChainId                  string `json:"chain_id"`
	HeadBlockNum             int64  `json:"head_block_num"`
	LastIrreversibleBlockNum int64  `json:"last_irreversible_block_num"`
	LastIrreversibleBlockId  string `json:"last_irreversible_block_id"`
	HeadBlockId              string `json:"head_block_id"`
}

type GetBlockResp struct {
	Timestamp      string `json:"timestamp"`
	Id             string `json:"id"`
	BlockNum       int64  `json:"block_num"`
	RefBlockPrefix int64  `json:"ref_block_prefix"`
}

func (rpc *NodeRpc) requestEosReturnError(requestUrl string, requestType EosRequestType,
	params interface{}) (string, error) {
	url := requestUrl + string(requestType)
	e, err := json.Marshal(params)
	var payload *strings.Reader
	if err != nil {
		payload = strings.NewReader("")
	} else {
		payload = strings.NewReader(string(e))
	}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 500 {
		r, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		return string(r), nil
	} else {
		return "", errors.New("cannot request node")
	}
}

func (rpc *NodeRpc) requestNodeReturnError(requestType EosRequestType, params interface{}) (string, error) {
	return rpc.requestEosReturnError(rpc.EosNodeUrl, requestType, params)
}

func (rpc *NodeRpc) requestWalletReturnError(requestType EosRequestType, params interface{}) (string, error) {
	return rpc.requestEosReturnError(rpc.EosWalletUrl, requestType, params)
}

func (rpc *NodeRpc) requestNode(requestType EosRequestType, params interface{}) (string) {
	result, err := rpc.requestNodeReturnError(requestType, params)
	if err != nil {
		panic(err)
	}
	return result
}

func (rpc *NodeRpc) requestWallet(requestType EosRequestType, params interface{}) (string) {
	result, err := rpc.requestWalletReturnError(requestType, params)
	if err != nil {
		panic(err)
	}
	return result
}

type AbiJsonToBinResp struct {
	Binargs string `json:"binargs"`
}

type WalletSignedResp struct {
	Expiration     string      `json:"expiration"`
	RefBlockNum    int64       `json:"ref_block_num"`
	RefBlockPrefix int64       `json:"ref_block_prefix"`
	Actions        interface{} `json:"actions"`
	Signatures     interface{} `json:"signatures"`
}

func (rpc *NodeRpc) actionAbiJsonToBin(actionType AbiActionType, params map[string]interface{}) (*map[string]interface{}) {
	switch actionType {
	case SYSTEM_NEW_ACCOUNT:
		accountName := params["accountName"].(string)
		ownerPublicKey := params["ownerPublicKey"].(string)
		activePublicKey := params["activePublicKey"].(string)

		resultStr := rpc.requestNode(CHAIN_ABI_JSON_TO_BIN, map[string]interface{}{
			"code":   EOS_SYSTEM_ACCOUNT,
			"action": "newaccount",
			"args": map[string]interface{}{
				"creator": rpc.UotcEosAccount,
				"name":    accountName,
				"owner": map[string]interface{}{
					"threshold": 1,
					"keys": []map[string]interface{}{
						{
							"key":    ownerPublicKey,
							"weight": 1,
						},
					},
					"accounts": []map[string]interface{}{},
					"waits":    []map[string]interface{}{},
				},
				"active": map[string]interface{}{
					"threshold": 1,
					"keys": []map[string]interface{}{
						{
							"key":    activePublicKey,
							"weight": 1,
						},
					},
					"accounts": []map[string]interface{}{},
					"waits":    []map[string]interface{}{},
				},
			},
		})

		abiJsonToBinResp := &AbiJsonToBinResp{}
		err := json.Unmarshal([]byte(resultStr), abiJsonToBinResp)
		if err != nil {
			panic(err)
		}

		return &map[string]interface{}{
			"account": EOS_SYSTEM_ACCOUNT,
			"name":    "newaccount",
			"authorization": []map[string]interface{}{
				{
					"actor":      rpc.UotcEosAccount,
					"permission": "active",
				},
			},
			"data": abiJsonToBinResp.Binargs,
		}
	case SYSTEM_BUY_RAM:
		accountName := params["accountName"].(string)
		newAccountBytes := params["newAccountBytes"].(int)

		resultStr := rpc.requestNode(CHAIN_ABI_JSON_TO_BIN, map[string]interface{}{
			"code":   EOS_SYSTEM_ACCOUNT,
			"action": "buyrambytes",
			"args": map[string]interface{}{
				"payer":    rpc.UotcEosAccount,
				"receiver": accountName,
				"bytes":    newAccountBytes,
			},
		})

		abiJsonToBinResp := &AbiJsonToBinResp{}
		err := json.Unmarshal([]byte(resultStr), abiJsonToBinResp)
		if err != nil {
			panic(err)
		}

		return &map[string]interface{}{
			"account": EOS_SYSTEM_ACCOUNT,
			"name":    "buyrambytes",
			"authorization": []map[string]interface{}{
				{
					"actor":      rpc.UotcEosAccount,
					"permission": "active",
				},
			},
			"data": abiJsonToBinResp.Binargs,
		}
	case SYSTEM_DELEGATE_BW:
		accountName := params["accountName"].(string)
		newAccountNet := params["newAccountNet"].(float64)
		newAccountCpu := params["newAccountCpu"].(float64)

		resultStr := rpc.requestNode(CHAIN_ABI_JSON_TO_BIN, map[string]interface{}{
			"code":   EOS_SYSTEM_ACCOUNT,
			"action": "delegatebw",
			"args": map[string]interface{}{
				"from":               rpc.UotcEosAccount,
				"receiver":           accountName,
				"stake_net_quantity": fmt.Sprintf("%.4f %s", newAccountNet, rpc.EosCryptoName),
				"stake_cpu_quantity": fmt.Sprintf("%.4f %s", newAccountCpu, rpc.EosCryptoName),
				"transfer":           false,
			},
		})

		abiJsonToBinResp := &AbiJsonToBinResp{}
		err := json.Unmarshal([]byte(resultStr), abiJsonToBinResp)
		if err != nil {
			panic(err)
		}

		return &map[string]interface{}{
			"account": EOS_SYSTEM_ACCOUNT,
			"name":    "delegatebw",
			"authorization": []map[string]interface{}{
				{
					"actor":      rpc.UotcEosAccount,
					"permission": "active",
				},
			},
			"data": abiJsonToBinResp.Binargs,
		}
	}
	panic(errors.New("Abi json to bin action not supported"))
}

func (rpc *NodeRpc) getInfo() (*GetInfoResp, *GetBlockResp) {
	chainInfoStr := rpc.requestNode(CHAIN_GET_INFO, map[string]interface{}{})
	getInfoResp := &GetInfoResp{}
	err := json.Unmarshal([]byte(chainInfoStr), getInfoResp)
	if err != nil {
		panic(err)
	}

	blockInfoStr := rpc.requestNode(CHAIN_GET_BLOCK, map[string]interface{}{
		"block_num_or_id": getInfoResp.HeadBlockNum,
	})
	getBlockResp := &GetBlockResp{}
	err = json.Unmarshal([]byte(blockInfoStr), getBlockResp)
	if err != nil {
		panic(err)
	}
	return getInfoResp, getBlockResp
}

func (rpc *NodeRpc) formatExpiration() (string) {
	expireTime := time.Now().UTC().Add(time.Minute)
	return expireTime.Format("2006-01-02T15:04:05.000")
}

func (rpc *NodeRpc) signTransaction(param interface{}) (*WalletSignedResp) {
	rpc.unlockWallet()
	resultStr := rpc.requestWallet(WALLET_SIGN_TRANSACTION, param)
	walletSignedResp := &WalletSignedResp{}
	err := json.Unmarshal([]byte(resultStr), walletSignedResp)
	if err != nil {
		panic(err)
	}
	return walletSignedResp
}

func (rpc *NodeRpc) unlockWallet() () {
	rpc.requestWalletReturnError(WALLET_UNLOCK, []interface{}{
		rpc.UotcWallet, rpc.UotcWalletPassword,
	})
}
