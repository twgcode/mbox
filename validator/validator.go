package validator

import (
	"fmt"
	"sync"

	"github.com/twgcode/mbox/exception"

	"github.com/go-playground/locales/en"

	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	v     *validator.Validate
	trans ut.Translator
	once  sync.Once
)

// V 全局校验器
func V() *validator.Validate {
	return v
}

func Init(locale string) (err error) {
	once.Do(func() {
		v, trans, err = NewValidate(locale)
	})
	return err
}

func Validate(target interface{}) (err error) {
	err = v.Struct(target)
	return err
}

// ValidateRequest 校验请求参数
func ValidateRequest(req interface{}) (err error) {
	if err = v.Struct(req); err != nil {
		err = exception.NewInvalidParam(err.Error())
		return
	}
	return
}

// ValidateBadRequest 校验请求参数 使用  BadRequest
func ValidateBadRequest(req interface{}) (err error) {
	if err = v.Struct(req); err != nil {
		err = exception.NewBadRequest(err.Error())
		return
	}
	return
}

// NewValidate 创建一个 校验器
func NewValidate(locale string) (v *validator.Validate, trans ut.Translator, err error) {
	var (
		ok bool
	)
	v = validator.New()
	zhT := zhongwen.New()
	enT := en.New() // 英文翻译器
	uni := ut.New(enT, zhT, enT)
	trans, ok = uni.GetTranslator(locale)
	if !ok {
		return nil, nil, fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}

	// 注册翻译器
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	return
}
