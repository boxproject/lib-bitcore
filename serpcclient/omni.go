package serpcclient

import (
	"encoding/json"
	"github.com/boxproject/lib-bitcore/sebtcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

type FutureOmniGetbalance chan *response
type FutureOmniGetTransaction chan *response
type FutureOmniListTransactions chan *response

type OmniBalance struct {
	Balance  string
	Reserved string
}

type OmniGetTransactionResult struct {
	//"txid": "7ff0c467ed54285406469b2e3abf03143efb1a63bd03e070dc07a033371314c9",
	Txid string

	//"fee": "0.00005140",
	Fee string

	//"sendingaddress": "n1GVisU77i1LLq5f4tV9KhrwxHXhzRgPCm",
	Sendingaddress string

	//"referenceaddress": "my111isbAUBSkk2CEBzCwcnRr6Eavjkhgt",
	Referenceaddress string

	//"ismine": true,
	Ismine bool

	//"version": 0,
	Version int

	//"type_int": 0,
	Type_int int `json:"type_int"`

	//"type": "Simple Send",
	Type string

	//"propertyid": 2147483651,
	Propertyid int64

	//"divisible": false,
	Divisible bool

	//"amount": "1000",
	Amount string

	//"valid": true,
	Valid bool

	//"blockhash": "6108f9231e081f28125eb2984a70aa3a2622cb95f84b498ff1b570d1caaef3fe",
	Blockhash string

	//"blocktime": 1531730633,
	Blocktime int64

	//"positioninblock": 1,
	Positioninblock int64

	//"block": 114,
	Block int64

	//"confirmations": 6
	Confirmations int64
}

/***
"txid": "e3785c0201fbb0012b7dd88770563cb8083ba012b5544bba6d3e6e35888cca61",
    "fee": "0.00005140",
    "sendingaddress": "n1GVisU77i1LLq5f4tV9KhrwxHXhzRgPCm",
    "referenceaddress": "mopRxRL9XfxbDy5CEDbUiXY8gfDuuYLyv7",
    "ismine": true,
    "version": 0,
    "type_int": 0,
    "type": "Simple Send",
    "propertyid": 2147483651,
    "divisible": false,
    "amount": "7",
    "valid": true,
    "blockhash": "2ad2ec41da9fe55687a73b0a792aeb56af7d2ce9599a11b357da09e7dec813a4",
    "blocktime": 1531909645,
    "positioninblock": 2,
    "block": 534,
    "confirmations": 12
*/
type OmniListTransactionResult struct {
	Txid             string
	Fee              string
	Sendingaddress   string
	Referenceaddress string
	Ismine           bool
	Version          int
	TypeInt          int `json:"type_int"`
	Type             string
	Propertyid       int64
	Divisible        bool
	Amount           string
	Valid            bool
	Blockhash        string
	Blocktime        int64
	Positioninblock  int
	Block            int
	Confirmations    int
}

// Receive waits for the response promised by the future and returns the
// available balance from the server for the specified account.
func (r FutureOmniGetbalance) Receive() (OmniBalance, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return OmniBalance{}, err
	}

	// Unmarshal result as a OmniBalance.
	var balance = OmniBalance{}
	err = json.Unmarshal(res, &balance)
	if err != nil {
		return OmniBalance{}, err
	}

	return balance, nil
}

// omni_getbalance "address" propertyid
//Returns the token balance for a given address and property.

func (c *Client) OmniGetbalance(address string, propertyid int) (OmniBalance, error) {
	cmd := sebtcjson.NewOmniGetbalanceCmd(address, propertyid)
	return ((FutureOmniGetbalance)(c.sendCmd(cmd))).Receive()
}

func (c *Client) OmniGetTransaction(txHash *chainhash.Hash) (*OmniGetTransactionResult, error) {

	hash := ""
	if txHash != nil {
		hash = txHash.String()
	}

	cmd := sebtcjson.NewOmniGetTransactionCmd(hash)
	return ((FutureOmniGetTransaction)(c.sendCmd(cmd))).Receive()
}

func (r FutureOmniGetTransaction) Receive() (*OmniGetTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a OmniBalance.
	var transactionResult = new(OmniGetTransactionResult)
	err = json.Unmarshal(res, &transactionResult)
	if err != nil {
		return nil, err
	}

	return transactionResult, nil
}

func (c *Client) OmniListTransactions(address string, max, skipCount, startBlock, endBlock int) ([]OmniListTransactionResult, error) {
	cmd := sebtcjson.NewOmniListTransactionsCmd(address, max, skipCount, startBlock, endBlock)
	return ((FutureOmniListTransactions)(c.sendCmd(cmd))).Receive()
}

func (r FutureOmniListTransactions) Receive() ([]OmniListTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var transactionResult []OmniListTransactionResult
	err = json.Unmarshal(res, &transactionResult)
	if err != nil {
		return nil, err
	}

	return transactionResult, nil
}
