package response

// {
// 	"success": true,
// 	"data": {},
// 	"errorCode": "1001",
// 	"errorMessage": "error message",
// 	"showType": 2,
// 	"traceId": "someid",
// 	"host": "10.1.1.1"
// 	}
type Result struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorCode    int         `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Current  uint64      `json:"current"`
	PageSize uint64      `json:"pageSize"`
	Total    uint64      `json:"total"`
}

func OK(data interface{}) *Result {
	return &Result{
		Success:      true,
		Data:         data,
		ErrorCode:    0,
		ErrorMessage: "OK",
	}
}

func Page(data interface{}, current uint64, pageSize uint64, total uint64) *Result {
	return &Result{
		Success: true,
		Data: &PageResult{
			List:     data,
			Current:  current,
			PageSize: pageSize,
			Total:    total,
		},
		ErrorCode:    0,
		ErrorMessage: "OK",
	}
}
