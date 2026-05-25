// Package types 提供项目级跨层契约入口。
//
// 边界说明:
// - Crypto 是 pkg/crypto.Crypto 的兼容别名,用于减少上层对具体包路径的感知。
// - CacheInjectable 是组件可注入缓存的跨层接口,参数使用 pkg/cache.Cache。
// - 本包不是业务领域模型集合,不承载 HTTP 响应 helper、数据库模型或认证实现。
// - 新增聚合类型前应先确认它是否真的需要跨层共享。
package types
