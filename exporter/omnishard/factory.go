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

package omnishard

import (
	"time"

	"github.com/open-telemetry/opentelemetry-collector/config/configerror"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/open-telemetry/opentelemetry-collector/exporter"
	"go.uber.org/zap"
)

const (
	// The value of "type" key in configuration.
	typeStr = "omnishard"

	// Default values for config options.
	defSendConcurrency       = 20
	defNumWorkers            = 1
	defMaxRecordSize         = 100000
	defBatchFlushInterval    = 5 * time.Second
	defMaxAllowedSizePerSpan = 900000
)

// Factory is the factory for OmniShard exporter.
type Factory struct {
}

// Type gets the type of the Exporter config created by this factory.
func (f *Factory) Type() string {
	return typeStr
}

// CreateDefaultConfig creates the default configuration for exporter.
func (f *Factory) CreateDefaultConfig() configmodels.Exporter {
	return &Config{
		SendConcurrency:       defSendConcurrency,
		NumWorkers:            defNumWorkers,
		BatchFlushInterval:    defBatchFlushInterval,
		MaxAllowedSizePerSpan: defMaxAllowedSizePerSpan,
		MaxRecordSize:         defMaxRecordSize,
	}
}

// CreateTraceExporter initializes and returns a new trace exporter
func (f *Factory) CreateTraceExporter(logger *zap.Logger, cfg configmodels.Exporter) (exporter.TraceExporter, error) {
	config := cfg.(*Config)
	if config.SendConcurrency < 1 {
		config.SendConcurrency = defSendConcurrency
	}
	if config.NumWorkers < 1 {
		config.NumWorkers = 1
	}

	e, err := NewExporter(config, logger, NewClientUnary(logger))
	if err != nil {
		return nil, err
	}

	return e, nil
}

// CreateMetricsExporter creates a metrics exporter based on this config.
func (f *Factory) CreateMetricsExporter(logger *zap.Logger, cfg configmodels.Exporter) (exporter.MetricsExporter, error) {
	return nil, configerror.ErrDataTypeIsNotSupported
}
