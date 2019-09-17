// Copyright 2019 Omnition Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package omnitelk

import (
	"github.com/open-telemetry/opentelemetry-service/config/configmodels"
)

// Config contains the main configuration options for the OmnitelK exporter.
type Config struct {
	configmodels.ExporterSettings `mapstructure:",squash"`

	// Endpoint of the server to connect to.
	Endpoint string `mapstructure:"endpoint"`

	// Number of gRPC streams to use concurrently to send encoded
	// data to the server. Default value is 1. Higher values may be necessary
	// to achieve good throughput.
	// TODO: run perf test and set a recommendation.
	Streams int `mapstructure:"streams"`
}
