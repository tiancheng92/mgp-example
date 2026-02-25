package validator

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorhill/cronexpr"
	"github.com/tiancheng92/mgp"
)

func dateTimeFormatValidatorFunc(fieldLevel validator.FieldLevel) bool {
	date := fieldLevel.Field().String()
	if date == "" {
		return true
	}
	layout := fieldLevel.Param()
	if _, err := time.ParseInLocation(layout, date, time.Local); err != nil {
		return false
	}
	return true
}

func cronValidatorFunc(fieldLevel validator.FieldLevel) bool {
	cron := fieldLevel.Field().String()
	if cron == "" {
		return true
	}
	arrayCount := len(strings.Split(cron, " "))

	if (fieldLevel.Param() == "minute" && (arrayCount == 6 || arrayCount == 5)) || (fieldLevel.Param() == "second" && arrayCount == 7) {
		_, err := cronexpr.Parse(cron)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func init() {
	mgp.RegisterValidate(
		"date_time_format",
		"string",
		"{0}日期/时间格式错误(格式: {1})",
		dateTimeFormatValidatorFunc,
		nil,
	)
	mgp.RegisterValidate(
		"cron",
		"string",
		"{0} cron表达式错误 ({1})",
		cronValidatorFunc,
		nil,
	)
}
