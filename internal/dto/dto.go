package dto

import (
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// Copy copier的封装,包含字段映射
func Copy(toValue, fromValue interface{}) (err error) {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: false,
		DeepCopy:    false,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: &timestamppb.Timestamp{},
				Fn: func(src interface{}) (interface{}, error) {
					t := src.(time.Time)
					return timestamppb.New(t), nil
				},
			},
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					t := src.(time.Time)
					return t.Format("2006-01-02 15:04:05"), nil
				},
			},
			{
				SrcType: gorm.DeletedAt{},
				DstType: &timestamppb.Timestamp{},
				Fn: func(src interface{}) (interface{}, error) {
					t := src.(gorm.DeletedAt)
					return timestamppb.New(t.Time), nil
				},
			},
		},
	})
}
