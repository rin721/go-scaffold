// Package plugin provides a protocol-neutral plugin runtime.
//
// The package is intentionally independent from any application layer. It does
// not import internal packages and does not know where plugin implementations
// live. A host project can keep local implementations under plugins/* and
// register their factories with Manager.
package plugin
