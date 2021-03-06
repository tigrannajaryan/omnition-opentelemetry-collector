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
	"path"
	"testing"
	"time"

	"github.com/open-telemetry/opentelemetry-collector/config"
	"github.com/open-telemetry/opentelemetry-collector/config/configcheck"
	"github.com/open-telemetry/opentelemetry-collector/config/configgrpc"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultConfig(t *testing.T) {
	factories, err := config.ExampleComponents()
	assert.Nil(t, err)

	factory := &Factory{}
	factories.Exporters[factory.Type()] = factory
	cfg, err := config.LoadConfigFile(
		t, path.Join(".", "testdata", "default.yaml"), factories,
	)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	e := cfg.Exporters["omnishard"]

	assert.Equal(t, e,
		&Config{
			ExporterSettings: configmodels.ExporterSettings{
				TypeVal: "omnishard",
				NameVal: "omnishard",
			},
			SendConcurrency:       defSendConcurrency,
			BatchFlushInterval:    defBatchFlushInterval,
			MaxRecordSize:         defMaxRecordSize,
			MaxAllowedSizePerSpan: defMaxAllowedSizePerSpan,
			NumWorkers:            defNumWorkers,
		},
	)
}

func TestConfig(t *testing.T) {
	factories, err := config.ExampleComponents()
	assert.Nil(t, err)

	factory := &Factory{}
	factories.Exporters[factory.Type()] = factory
	cfg, err := config.LoadConfigFile(
		t, path.Join(".", "testdata", "config.yaml"), factories,
	)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	e := cfg.Exporters["omnishard"]

	assert.Equal(t, e,
		&Config{
			ExporterSettings: configmodels.ExporterSettings{
				TypeVal: "omnishard",
				NameVal: "omnishard",
			},
			GRPCSettings: configgrpc.GRPCSettings{
				Endpoint: "0.0.0.0:1234",
				Headers: map[string]string{
					"some.key":                     "some.value",
					"some_other":                   "some.value",
					"you can have a sentence here": "F0000000-0000-0000-0000-000000000000",
				},
				UseSecure: true,
			},
			SendConcurrency:       4,
			NumWorkers:            4,
			MaxAllowedSizePerSpan: 50000,
			MaxRecordSize:         200000,
			BatchFlushInterval:    2 * time.Second,
		},
	)
}

func TestConfigCheck(t *testing.T) {
	cfg := (&Factory{}).CreateDefaultConfig()
	assert.NoError(t, configcheck.ValidateConfig(cfg))
}
