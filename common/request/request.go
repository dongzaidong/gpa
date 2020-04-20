package request

//Paramer 参数接口
type Paramer interface {
	Type() string
}

// Request 服务请求
type Request struct {
	ID string
	CTime string
	ServiceMap map[string]Paramer
}

// NewRequest 服务实例
func NewRequest(id string, opts ...OptionF) *Request {
	req := &Request{
		ID: id
		ServiceMap: make(map[string]Paramer),
	}

	for _, opt := range opts {
		opt(req)
	}
	return req
}

// FeaturesA 功能A的处理参数
type FeaturesA struct {
	Name      string
	Condition string
}

// Type 功能名称
func (f *FeaturesA) Type() string {
	return f.Name
}

// FeaturesB 功能C的处理参数
type FeaturesB struct {
	Name      string
	Condition string
}

// Type 功能名称
func (f *FeaturesA) Type() string {
	return f.Name
}

// FeaturesC 功能C的处理参数
type FeaturesC struct {
	Name      string
	Condition string
}

// Type 功能名称
func (f *FeaturesA) Type() string {
	return f.Name
}
