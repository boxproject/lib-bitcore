# lib-bitcore
a lib for bitcore

modify list:

    1. chain.go, we added some RescanBlock methods below, you can use these methods rescan block convinently:

        func (c *Client) RescanBlockChainAsync() FutureRescanBlockChainResult

        func (c *Client) RescanBlockChainStartAsync(startHeight int) FutureRescanBlockChainResult

        func (c *Client) RescanBlockChainStartStopAsync(startHeight, stopHeight int) FutureRescanBlockChainResult

        func (c *Client) RescanBlockChain() (*sebtcjson.RescanBlockChanResult,error)

        func (c *Client) RescanBlockChainStart(startHeight int) (*sebtcjson.RescanBlockChanResult,error)

        func (c *Client) RescanBlockChainStartStop(startHeight, stopHeight int) (*sebtcjson.RescanBlockChanResult,error)

     2. wallet.go, we added some ImportAddress methods,you do not need rescan all blocks yet.

        func (c *Client) ImportAddressRescanAsync(address ,lable string, rescan bool) FutureImportAddressResult

        func (c *Client) ImportAddressRescan(address,lable string, rescan bool) error