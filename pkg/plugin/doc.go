// Package plugin provides a protocol-neutral plugin runtime with hooks.
//
// The package is intentionally independent from any application layer. It does
// not import internal packages, IAM packages, logging, or configuration. It
// does not discover services or actively register plugin services from
// configuration. Plugin services or the host composition layer create plugin
// instances, register hooks, and register plugins with Manager.
package plugin
