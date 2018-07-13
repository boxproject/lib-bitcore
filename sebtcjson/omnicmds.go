package sebtcjson


type OmniGetbalanceCmd struct{
	Address *string
	Propertyid *int
}
func NewOmniGetbalanceCmd(address string ,propertyid int) *OmniGetbalanceCmd {
	return &OmniGetbalanceCmd{
		Address:&address,
		Propertyid:&propertyid,
	}
}
