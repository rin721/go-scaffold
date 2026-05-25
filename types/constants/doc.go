// Package constants 定义跨层共享常量。
//
// 边界说明:
// - 应用命令名、默认配置路径、关闭超时和 executor pool 名称属于公共运行契约。
// - 常量值可能被 cmd、internal 和 pkg 文档引用,修改时需要同步测试和状态文档。
// - 本包不放业务枚举、HTTP 响应结构或可变配置。
package constants
