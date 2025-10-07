package request

type AddRequest struct {
	Link   string   `json:"link"`   // 要添加的链接
	Names  []string `json:"names"`  // 名称列表
	Remark string   `json:"remark"` // 备注
	Tags   []string `json:"tags"`   // 标签列表
	Name   string   `json:"name"`   // 名称
}
