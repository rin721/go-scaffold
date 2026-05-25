// Package plugin provides a protocol-neutral plugin runtime.
//
// The package is intentionally independent from any application layer. It does
// not import internal packages, discover services, or actively register plugin
// services from configuration. Plugin services or the host composition layer
// create plugin instances and register them with Manager.
package plugin
