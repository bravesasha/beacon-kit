// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
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
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/berachain/beacon-kit/mod/beacon/blockchain"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	dastore "github.com/berachain/beacon-kit/mod/da/pkg/store"
	datypes "github.com/berachain/beacon-kit/mod/da/pkg/types"
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/components/storage"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core"
	"github.com/berachain/beacon-kit/mod/storage/pkg/beacondb"
	"github.com/berachain/beacon-kit/mod/storage/pkg/beacondb/encoding"
	depositdb "github.com/berachain/beacon-kit/mod/storage/pkg/deposit"
)

type StorageBackendInput struct {
	depinject.In
	ChainSpec         primitives.ChainSpec
	AvailabilityStore *dastore.Store[types.BeaconBlockBody]
	Environment       appmodule.Environment
	DepositStore      *depositdb.KVStore[*types.Deposit]
}

func ProvideStorageBackend(
	in StorageBackendInput,
) blockchain.StorageBackend[
	*dastore.Store[types.BeaconBlockBody],
	types.BeaconBlockBody,
	BeaconState,
	*datypes.BlobSidecars,
	*types.Deposit,
	*depositdb.KVStore[*types.Deposit],
] {
	payloadCodec := &encoding.
		SSZInterfaceCodec[*types.ExecutionPayloadHeader]{}
	return storage.NewBackend[
		*dastore.Store[types.BeaconBlockBody],
		types.BeaconBlockBody,
		core.BeaconState[
			*types.BeaconBlockHeader, *types.ExecutionPayloadHeader, *types.Fork,
			*types.Validator, *engineprimitives.Withdrawal,
		],
	](
		in.ChainSpec,
		in.AvailabilityStore,
		beacondb.New[
			*types.Fork,
			*types.BeaconBlockHeader,
			*types.ExecutionPayloadHeader,
			*types.Eth1Data,
			*types.Validator,
		](in.Environment.KVStoreService, payloadCodec),
		in.DepositStore,
	)
}
