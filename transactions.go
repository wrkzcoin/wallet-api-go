package walletapi

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Transfer - represents a transfer object
type Transfer struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

// Transaction - represents a transaction object
type Transaction struct {
	BlockHeight           int64      `json:"blockHeight"`
	Fee                   int64      `json:"fee"`
	Hash                  string     `json:"hash"`
	IsCoinbaseTransaction bool       `json:"isCoinbaseTransaction"`
	PaymentID             string     `json:"paymentID"`
	Timestamp             int64      `json:"timestamp"`
	UnlockTime            int64      `json:"unlockTime"`
	Transfers             []Transfer `json:"transfers"`
}

// Transactions - represents a transactions object
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
	Transaction  Transaction   `json:"transaction"`
}

// GetAllTransactions - gets all the tranactions in the wallet container
func (wAPI WalletAPI) GetAllTransactions() (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions",
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetTransactionByHash - gets the transaction with the given hash in the wallet container
func (wAPI WalletAPI) GetTransactionByHash(hash string) (tx *Transaction, err error) {
	var txs Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[404])
		}
	}()

	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/hash/"+hash,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &txs)
		if err != nil {
			panic(err)
		}
	}

	return &txs.Transaction, err
}

// GetUnconfirmedTransactions - gets all unconfirmed outgoing transactions
func (wAPI WalletAPI) GetUnconfirmedTransactions() (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/unconfirmed",
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetUnconfirmedTransactionsByAddress - gets all unconfirmed outgoing transactions for a given address
func (wAPI WalletAPI) GetUnconfirmedTransactionsByAddress(address string) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/unconfirmed/"+address,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetTransactionsByStartHeight - gets 1000 transactions for the wallet starting at startHeight
func (wAPI WalletAPI) GetTransactionsByStartHeight(startHeight int64) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	start := strconv.FormatInt(startHeight, 10)
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/"+start,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetTransactionsInRange - gets transactions for the wallet given a range of block heights
func (wAPI WalletAPI) GetTransactionsInRange(start, end int64) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	low := strconv.FormatInt(start, 10)
	high := strconv.FormatInt(end, 10)
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/"+low+"/"+high,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetAddressTransactionsByStartHeight - gets 1000 transactions for the address starting at startHeight
func (wAPI WalletAPI) GetAddressTransactionsByStartHeight(address string, startHeight int64) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	start := strconv.FormatInt(startHeight, 10)
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/address/"+address+"/"+start,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetAddressTransactionsInRange - gets transactions for the address given a range of block heights
func (wAPI WalletAPI) GetAddressTransactionsInRange(address string, start, end int64) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	low := strconv.FormatInt(start, 10)
	high := strconv.FormatInt(end, 10)
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/address/"+address+"/"+low+"/"+high,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// SendTransactionBasic - sends a transaction
func (wAPI WalletAPI) SendTransactionBasic(destination, paymentID string, amount uint64) (string, error) {
	var txHash string

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/send/basic",
		makeJSONString(map[string]interface{}{
			"destination": destination,
			"amount":      amount,
			"paymentId":   paymentID,
		}),
	)

	if err == nil {
		txHash = (*resp)["transactionHash"].(string)
	}

	return txHash, err
}

// SendTransactionAdvanced - sends a transaction
func (wAPI WalletAPI) SendTransactionAdvanced(
	destinations []map[string]interface{},
	mixin, fee, sourceAddresses, paymentID, changeAddress, unlockTime interface{}) (string, error) {
	var txHash string

	body := map[string]interface{}{
		"destinations": destinations,
	}

	if mixin != nil {
		body["mixin"] = mixin
	}
	if fee != nil {
		body["fee"] = fee
	}
	if sourceAddresses != nil {
		body["sourceAddresses"] = sourceAddresses
	}
	if paymentID != nil {
		body["paymentID"] = paymentID
	}
	if unlockTime != nil {
		body["unlockTime"] = unlockTime
	}
	if changeAddress != nil {
		body["changeAddress"] = changeAddress
	}

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/send/advanced",
		makeJSONString(body),
	)

	if err == nil {
		txHash = (*resp)["transactionHash"].(string)
	}

	return txHash, err
}

// SendFusionBasic - sends a fusion transaction
func (wAPI WalletAPI) SendFusionBasic() (string, error) {
	var txHash string

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/send/fusion/basic",
		"",
	)

	if err == nil {
		txHash = (*resp)["transactionHash"].(string)
	}

	return txHash, err
}

// SendFusionAdvanced - sends a fusion transaction
func (wAPI WalletAPI) SendFusionAdvanced(sourceAddresses []string, destination string) (string, error) {
	var txHash string

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/send/fusion/advanced",
		makeJSONString(map[string]interface{}{
			"mixin":           MIXIN,
			"sourceAddresses": sourceAddresses,
			"destination":     destination,
		}),
	)

	if err == nil {
		txHash = (*resp)["transactionHash"].(string)
	}

	return txHash, err
}

// GetTransactionPrivateKey - gets the private key of a transaction with the given hash
func (wAPI WalletAPI) GetTransactionPrivateKey(hash string) (string, error) {
	var privKey string

	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/privatekey/"+hash,
		"",
	)
	if err == nil {
		privKey = (*resp)["transactionPrivateKey"].(string)
	}

	return privKey, err
}