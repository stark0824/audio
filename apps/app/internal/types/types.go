// Code generated by goctl. DO NOT EDIT.
package types

type Req struct {
}

type Resp struct {
}

type DramaInfoRequest struct {
	DramaId int `json:"dramaId,default=0"`
}

type DramaInfoResp struct {
	DramaInfo    Info           `json:"dramaInfo"`
	CataLog      map[int]string `json:"catalog"`
	Category     map[int]string `json:"category"`
	Process      map[int]string `json:"process"`
	Status       map[int]string `json:"status"`
	PayType      []PayType      `json:"payType"`
	ProducerList []Producer     `json:"producerList"`
}

type RecommendListReq struct {
	Page  int64 `json:"page,default=1"`
	Limit int64 `json:"limit,default=20"`
}

type RecommendListResp struct {
	List  []RecPosition `json:"list"`
	Count int64         `json:"count"`
}

type RecommendDataReq struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`      // 标题
	NovelIds string `json:"novel_ids"` // 作品id集合: | 分隔
	Sort     int64  `json:"sort"`      // 排序
	TempType int64  `json:"temp_type"` // 模板类型 1:单排,2:横排,3:竖排
}

type RecommendOpsReq struct {
	Id int64 `json:"id"`
}

type ShortRecPositionResp struct {
	List []ShortRecPosition `json:"list"`
}

type ProducerListResp struct {
	List   []ProducerList `json:"list"`
	Status map[int]string `json:"status"`
}

type ProducerEditReq struct {
	Id     int64  `json:"id,default=0"`
	Name   string `json:"name,default=''"`
	Mark   string `json:"mark,default=''"`
	Status int64  `json:"status,default=0"`
	Uid    int64  `json:"uid,default=0"`
}

type ProducerEditResp struct {
	Id int64 `json:"id,default=0"`
}

type CollDramaReq struct {
	Did    int64  `json:"did,default=0"`
	Status int64  `json:"status,default=10"`
	Uid    string `header:"uid,default=''"`
}

type CheckCollDramaReq struct {
	Did int64  `json:"did,default=0"`
	Uid string `header:"uid,default=''"`
}

type CheckCollDramaResp struct {
	Follow int64 `json:"follow"`
}

type ShortInfo struct {
	Id         int64  `json:"id"`
	Cover      string `json:"cover"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	ShortIntro string `json:"short_intro"`
	Intro      string `json:"intro"`
}

type RecPosition struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`      // 标题
	TempType  int64  `json:"temp_type"` // 模板类型 1:单排,2:横排,3:竖排
	TempName  string `json:"temp_name"`
	Sort      int64  `json:"sort"`      // 排序
	NovelIds  string `json:"novel_ids"` // 作品id集合: | 分隔
	NovelName string `json:"novel_name"`
	IsDel     int64  `json:"is_del"`     // 0未删除 1已删除
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间
}

type ShortRecPosition struct {
	Id       int64       `json:"id"`
	Name     string      `json:"name"`      // 标题
	TempType int64       `json:"temp_type"` // 模板类型 1:单排,2:横排,3:竖排
	Sort     int64       `json:"sort"`      // 排序
	List     []ShortInfo `json:"drama_list"`
}

type Info struct {
	Id                  int64  `json:"id"`
	Cover               string `json:"cover"`
	Name                string `json:"name"`
	Author              string `json:"author"`
	Category            int64  `json:"category"`
	CatalogId           int64  `json:"catalog_id"`
	Nid                 int64  `json:"nid"`
	ProducerId          int64  `json:"nid"`
	ProducerName        string `json:"producer_name"`
	Process             int64  `json:"process"`
	ShortIntro          string `json:"short_intro"`
	Intro               string `json:"intro"`
	TotalPrice          int64  `json:"total_price"`
	IsPay               int64  `json:"is_pay"`
	PayIndex            int64  `json:"pay_index"`
	Price               int64  `json:"price"`
	DiscountAll         int64  `json:"discount_all"`
	DiscountAllEtime    string `json:"discount_all_etime"`
	DiscountSingle      int64  `json:"discount_single"`
	DiscountSingleEtime string `json:"discount_single_etime"`
	Status              int64  `json:"status"`
	ServerDivided       int64  `json:"server_divided"`
	ProducerDivided     int64  `json:"server_divided"`
}

type PayType struct {
	Value int64  `json:"value"`
	Label string `json:"label"`
}

type Producer struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ProducerList struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Mark         string `json:"mark"`
	Status       int64  `json:"status"`
	StatusVal    string `json:"status_val,default=''"`
	Uid          int64  `json:"uid"`
	UserNickname string `json:"user_nickname,default=''"`
	CreatedAt    string `json:"created_time"`
}