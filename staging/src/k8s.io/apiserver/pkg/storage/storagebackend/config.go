/*
Copyright 2018-2020, Arm Limited and affiliates.
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storagebackend

import (
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/storage/value"
)

const (
	StorageTypeUnset = ""
	StorageTypeETCD3 = "etcd3"

	DefaultCompactInterval = 12 * time.Hour
)

// Config is configuration for creating a storage backend.
type Config struct {
	// Type defines the type of storage backend. Default ("") is "etcd3".
	Type string
	// Prefix is the prefix to all keys passed to storage.Interface methods.
	Prefix string
	// ServerList is the list of storage servers to connect with.
	ServerList []string
	// TLS credentials
	KeyFile  string
	CertFile string
	CAFile   string
	// Paging indicates whether the server implementation should allow paging (if it is
	// supported). This is generally configured by feature gating, or by a specific
	// resource type not wishing to allow paging, and is not intended for end users to
	// set.
	Paging bool

	Codec runtime.Codec
	// Transformer allows the value to be transformed prior to persisting into etcd.
	Transformer value.Transformer

	// CompactionInterval is an interval of requesting compaction from apiserver.
	// If the value is 0, no compaction will be issued.
	CompactionInterval time.Duration

	// CountMetricPollPeriod specifies how often should count metric be updated
	CountMetricPollPeriod time.Duration
}

func NewDefaultConfig(prefix string, codec runtime.Codec) *Config {
	return &Config{
		Prefix:             prefix,
		Codec:              codec,
		CompactionInterval: DefaultCompactInterval,
	}
}
