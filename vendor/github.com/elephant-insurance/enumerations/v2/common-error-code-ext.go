package enumerations

import "fmt"

// ToString allows us to log uniform error messages for given error codes.
func (cecid *CommonErrorCodeID) ToString() string {
	rtn := `invalid common error code`
	if cecid != nil {
		if cc := CommonErrorCode.ByID(cecid); cc != nil {
			rtn = fmt.Sprintf(`%v: %v`, string(cc.ID), cc.Desc)
		}
	}

	return rtn
}
