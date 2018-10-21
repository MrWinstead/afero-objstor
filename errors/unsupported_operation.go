package errors

import "fmt"

type UnsupportedOperation struct{
	OperationFriendlyName string
	Reason string
}

func (uo *UnsupportedOperation) Error() string {
	str := fmt.Sprintf("unsupported operation '%v', %v",
		uo.OperationFriendlyName, uo.Reason)
	return str
}


