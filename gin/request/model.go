/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2022/10/20 18:55
@Description:
*/

package request

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/twgcode/mbox/pb/request"
)

// PageRequest 分页时 必须参数
type PageRequest struct {
	PageNum  uint64 `form:"page_num" json:"page_num" binding:"required,min=1"`   // 页码
	PageSize uint64 `form:"page_size" json:"page_size" binding:"required,min=1"` // 每页数据量
}

// ToPageRequest 转化为pb对应的 PageRequest
func (p *PageRequest) ToPageRequest() *request.PageRequest {
	return &request.PageRequest{PageNum: p.PageNum, PageSize: p.PageSize}
}

type TimeRangeRequest struct {
	StartTime time.Time `form:"start_time" json:"start_time" binding:""`
	EndTime   time.Time `form:"end_time" json:"end_time" binding:""`
}

// ToTimeRangeRequest 转化为pb对应的 TimeRangeRequest
func (t *TimeRangeRequest) ToTimeRangeRequest() *request.TimeRangeRequest {
	return &request.TimeRangeRequest{StartTime: timestamppb.New(t.StartTime), EndTime: timestamppb.New(t.EndTime)}
}

type KeywordRequest struct {
	Keyword string `form:"keyword" json:"keyword"`
}

// ToKeywordRequest 转化为pb对应的 KeywordRequest
func (k *KeywordRequest) ToKeywordRequest() *request.KeywordRequest {
	return &request.KeywordRequest{Keyword: k.Keyword}
}

type KeywordRequiredRequest struct {
	Keyword string `form:"keyword" json:"keyword" binding:"required"`
}

// ToKeywordRequest 转化为pb对应的 KeywordRequest
func (k *KeywordRequiredRequest) ToKeywordRequest() *request.KeywordRequest {
	return &request.KeywordRequest{Keyword: k.Keyword}
}
