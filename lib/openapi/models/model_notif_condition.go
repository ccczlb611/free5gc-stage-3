/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NotifCondition struct {
	MonitoredAttributes []string `json:"monitoredAttributes,omitempty" yaml:"monitoredAttributes" bson:"monitoredAttributes" mapstructure:"MonitoredAttributes"`
	UnmonitoredAttributes []string `json:"unmonitoredAttributes,omitempty" yaml:"unmonitoredAttributes" bson:"unmonitoredAttributes" mapstructure:"UnmonitoredAttributes"`
}