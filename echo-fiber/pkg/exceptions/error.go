package exceptions

// export interface response {
// 	success: boolean; // if request is success
// 	data?: any; // response data
// 	errorCode?: string; // code for errorType
// 	errorMessage?: string; // message display to user
// 	showType?: number; // error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
// 	traceId?: string; // Convenient for back-end Troubleshooting: unique request ID
// 	host?: string; // onvenient for backend Troubleshooting: host of current access server
//   }
type Error struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ShowType     int    `json:"showType"`
	TraceId      string `json:"traceId"`
	Host         string `json:"host"`
}

func (e *Error) Error() string {
	return e.ErrorMessage
}

func DefaultError(c int, m string) *Error {
	return &Error{
		Success:      false,
		ErrorCode:    c,
		ErrorMessage: m,
		ShowType:     9,
		TraceId:      "",
		Host:         "",
	}
}

// 404
func NotFound() *Error {
	return &Error{
		Success:      false,
		ErrorCode:    404,
		ErrorMessage: "404 NOT FOUND",
		ShowType:     9,
		TraceId:      "",
		Host:         "",
	}
}

// 400
func BadRequest() *Error {
	return &Error{
		Success:      false,
		ErrorCode:    400,
		ErrorMessage: "400 BAD REQUEST",
		ShowType:     4,
		TraceId:      "",
		Host:         "",
	}
}

// 403
func Forbidden() *Error {
	return &Error{
		Success:      false,
		ErrorCode:    403,
		ErrorMessage: "403 Forbidden",
		ShowType:     9,
		TraceId:      "",
		Host:         "",
	}
}

// 500
func InternalServerError() *Error {
	return &Error{
		Success:      false,
		ErrorCode:    500,
		ErrorMessage: "500 Internal Server Error",
		ShowType:     9,
		TraceId:      "",
		Host:         "",
	}

}
