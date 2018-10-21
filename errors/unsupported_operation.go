package errors

import "fmt"

// UnsupportedOperation describes an error when the requested operation is
// unsupported
type UnsupportedOperation struct {
	OperationFriendlyName string
	Reason                string
}

// Error formats the unsupported operation into a string and fulfils the error
// interface
func (uo *UnsupportedOperation) Error() string {
	str := fmt.Sprintf("unsupported operation '%v', %v",
		uo.OperationFriendlyName, uo.Reason)
	return str
}
