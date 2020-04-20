package request

type OptionF func(*Request)

//WithFeatures 设置服务参数
func WithFeatures(param Paramer) OptionF {
	return func(r *Request) {
		r.ServiceMap[param.name] = param
	}
}
