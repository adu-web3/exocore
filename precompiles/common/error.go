package common

const (
	ErrContractInputParaOrType = "the contract input parameter type or value error,arg index:%d, type is:%s,value:%v"
	ErrContractCaller          = "the caller doesn't have the permission to call this function"

	ErrInputClientChainAddrLength = "the length of input client chain addr doesn't match,input:%d,need:%d"

	ErrInputOperatorAddrLength = "mismatched length of the input operator address,actual is:%d,expect:%v"
)
