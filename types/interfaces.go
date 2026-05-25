package types

import (
	"github.com/rei0721/go-scaffold/pkg/cache"
	"github.com/rei0721/go-scaffold/pkg/crypto"
)

// Crypto 密码加密器类型别名
// 使用 pkg/crypto.Crypto 接口
type Crypto = crypto.Crypto

// CacheInjectable 定义可注入缓存的组件接口
// 所有需要使用缓存的组件应实现此接口
//
// 设计原则:
// - 使用 atomic.Value 存储 cache.Cache
// - 参数为nil时禁用缓存功能
// - 支持运行时动态替换
//
// 实现模式:
//
//	type Component struct {
//	    cache atomic.Value // 存储 cache.Cache
//	}
//
//	func (c *Component) SetCache(cache cache.Cache) {
//	    c.cache.Store(cache)
//	}
//
//	func (c *Component) getCache() cache.Cache {
//	    if cache := c.cache.Load(); cache != nil {
//	        return cache.(cache.Cache)
//	    }
//	    return nil
//	}
//
// 使用场景:
// - 解除组件初始化顺序依赖
// - 支持运行时动态替换缓存
// - 测试时可选不注入缓存
type CacheInjectable interface {
	// SetCache 设置缓存实例
	// 使用原子操作，并发安全
	// 参数:
	//   c: 缓存实例，为nil时禁用缓存功能
	SetCache(c cache.Cache)
}
