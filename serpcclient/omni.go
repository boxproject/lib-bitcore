package serpcclient

import (
	"github.com/boxproject/lib-bitcore/sebtcjson"
	"encoding/json"
)

type FutureOmniGetbalance chan *response

type OmniBalance struct{
	Balance string
	Reserved string
}

// Receive waits for the response promised by the future and returns the
// available balance from the server for the specified account.
func (r FutureOmniGetbalance) Receive() (OmniBalance, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return OmniBalance{}, err
	}

	// Unmarshal result as a OmniBalance.
	var balance  = OmniBalance{}
	err = json.Unmarshal(res, &balance)
	if err != nil {
		return OmniBalance{}, err
	}

	return balance, nil
}


// omni_getbalance "address" propertyid
//Returns the token balance for a given address and property.

func (c *Client) OmniGetbalance(address string ,propertyid int) (OmniBalance, error) {
	cmd := sebtcjson.NewOmniGetbalanceCmd(address,propertyid)
	return ((FutureOmniGetbalance)(c.sendCmd(cmd))).Receive()
}
