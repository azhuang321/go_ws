package utils

import (
	"context"
	"time"

	"user_api/global"
)

var Store = CaptchaRedisStore{}

type CaptchaRedisStore struct {}

func (c CaptchaRedisStore)Set(id string, value string) error {
	_,err := global.GClient.Set(context.Background(),id,value,time.Duration(global.Config.CaptchaExpireSec) * time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c CaptchaRedisStore)Get(id string, clear bool) string {
	ctx := context.Background()
	val,err := global.GClient.Get(ctx,id).Result()
	if clear {
		if _,err := global.GClient.Del(ctx,id).Result(); err != nil {
			return ""
		}
	}
	if err != nil {
		return ""
	}
	return val
}

func (c CaptchaRedisStore)Verify(id, answer string, clear bool) bool {
	ctx := context.Background()
	val,err := global.GClient.Get(ctx,id).Result()
	if clear {
		if _,err := global.GClient.Del(ctx,id).Result(); err != nil {
			return false
		}
	}
	if err != nil {
		return false
	}
	if val == answer {
		return true
	}
	return false
}