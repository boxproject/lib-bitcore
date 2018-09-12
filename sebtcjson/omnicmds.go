package sebtcjson

type OmniGetbalanceCmd struct {
	Address    *string
	Propertyid *int
}

func NewOmniGetbalanceCmd(address string, propertyid int) *OmniGetbalanceCmd {
	return &OmniGetbalanceCmd{
		Address:    &address,
		Propertyid: &propertyid,
	}
}

type OmniGetTransactionCmd struct {
	Txid string
}

func NewOmniGetTransactionCmd(address string) *OmniGetTransactionCmd {
	return &OmniGetTransactionCmd{
		Txid: address,
	}
}

type OmniListTransactionsCmd struct {
	Address    *string `json:"address"`
	Count      *int `jsonrpcdefault:"10"`
	Skip       *int	`json:"skip"`
	Startblock *int `jsonrpcdefault:"0"`
	Endblock   *int	`jsonrpcdefault:"999999"`
}

func NewOmniListTransactionsCmd(address string,max,skip,startBlock,endBlock int,) *OmniListTransactionsCmd {
	return &OmniListTransactionsCmd{
		Address:&address,
		Count:&max,
		Skip:&skip,
		Startblock:&startBlock,
		Endblock:&endBlock,
	}
}