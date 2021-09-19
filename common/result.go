package common

type Response struct {
	Code int `json:"code"`
	Detail string `json:"detail"`
	Data interface{} `json:"data"`
}

var (
	OK = response(200, "ok")
	Err = response(500, "")
)

// 构造函数
func response(code int, detail string) *Response {
	return &Response{
		Code: code,
		Detail: detail,
		Data: nil,
	}
}

// WithData 添加响应数据
func (res *Response) WithData(data interface{}) *Response {
	return &Response{
		Code: res.Code,
		Detail: res.Detail,
		Data: data,
	}
}

// WithDetail 添加详细信息
func (res *Response) WithDetail(detail string) *Response {
	return &Response{
		Code: res.Code,
		Detail: detail,
		Data: res.Data,
	}
}




