// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package components

import (
	"github.com/berachain/beacon-kit/mod/async/pkg/broker"
)

// ProvideBlobBroker provides a blob feed for the depinject framework.
func ProvideBlobBroker() *SidecarsBroker {
	return broker.New[*SidecarEvent](
		"blob-broker",
	)
}

// ProvideBlockBroker provides a block feed for the depinject framework.
func ProvideBlockBroker() *BlockBroker {
	return broker.New[*BlockEvent](
		"blk-broker",
	)
}

// ProvideGenesisBroker provides a genesis feed for the depinject framework.
func ProvideGenesisBroker() *GenesisBroker {
	return broker.New[*GenesisEvent](
		"genesis-broker",
	)
}

// ProvideSlotBroker provides a slot feed for the depinject framework.
func ProvideSlotBroker() *SlotBroker {
	return broker.New[*SlotEvent](
		"slot-broker",
	)
}

// ProvideStateRootBroker provides a state root feed for the depinject
// framework.
func ProvideStateRootBroker() *StateRootBroker {
	return broker.New[*StateRootEvent](
		"state-root-broker",
	)
}

// ProvideStatusBroker provides a status feed.
func ProvideStatusBroker() *StatusBroker {
	return broker.New[*StatusEvent](
		"status-broker",
	)
}

// ProvideValidatorUpdateBroker provides a validator updates feed.
func ProvideValidatorUpdateBroker() *ValidatorUpdateBroker {
	return broker.New[*ValidatorUpdateEvent](
		"validator-updates-broker",
	)
}

// DefaultBrokerProviders returns a slice of the default broker providers.
func DefaultBrokerProviders() []interface{} {
	return []interface{}{
		ProvideBlobBroker,
		ProvideBlockBroker,
		ProvideGenesisBroker,
		ProvideSlotBroker,
		ProvideStateRootBroker,
		ProvideStatusBroker,
		ProvideValidatorUpdateBroker,
	}
}
