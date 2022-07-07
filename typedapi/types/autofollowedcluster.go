// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// AutoFollowedCluster type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/ccr/stats/types.ts.ts#L27-L31
type AutoFollowedCluster struct {
	ClusterName              Name                    `json:"cluster_name"`
	LastSeenMetadataVersion  VersionNumber           `json:"last_seen_metadata_version"`
	TimeSinceLastCheckMillis DurationValueUnitMillis `json:"time_since_last_check_millis"`
}

// AutoFollowedClusterBuilder holds AutoFollowedCluster struct and provides a builder API.
type AutoFollowedClusterBuilder struct {
	v *AutoFollowedCluster
}

// NewAutoFollowedCluster provides a builder for the AutoFollowedCluster struct.
func NewAutoFollowedClusterBuilder() *AutoFollowedClusterBuilder {
	r := AutoFollowedClusterBuilder{
		&AutoFollowedCluster{},
	}

	return &r
}

// Build finalize the chain and returns the AutoFollowedCluster struct
func (rb *AutoFollowedClusterBuilder) Build() AutoFollowedCluster {
	return *rb.v
}

func (rb *AutoFollowedClusterBuilder) ClusterName(clustername Name) *AutoFollowedClusterBuilder {
	rb.v.ClusterName = clustername
	return rb
}

func (rb *AutoFollowedClusterBuilder) LastSeenMetadataVersion(lastseenmetadataversion VersionNumber) *AutoFollowedClusterBuilder {
	rb.v.LastSeenMetadataVersion = lastseenmetadataversion
	return rb
}

func (rb *AutoFollowedClusterBuilder) TimeSinceLastCheckMillis(timesincelastcheckmillis *DurationValueUnitMillisBuilder) *AutoFollowedClusterBuilder {
	v := timesincelastcheckmillis.Build()
	rb.v.TimeSinceLastCheckMillis = v
	return rb
}