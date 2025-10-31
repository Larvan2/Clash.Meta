package outboundgroup

type SelectAble interface {
	Set(string) error
	ForceSet(name string)
	CloseRelatedConns()
}

var (
	_ SelectAble = (*Fallback)(nil)
	_ SelectAble = (*URLTest)(nil)
	_ SelectAble = (*Selector)(nil)
)
