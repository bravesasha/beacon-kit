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
	"cosmossdk.io/depinject"
	sdklog "cosmossdk.io/log"
	"github.com/berachain/beacon-kit/mod/execution/pkg/deposit"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/components/metrics"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/messages"
)

// DepositServiceIn is the input for the deposit service.
type DepositServiceIn struct {
	depinject.In
	BeaconDepositContract *DepositContract
	ChainSpec             common.ChainSpec
	DepositStore          *DepositStore
	Dispatcher            *Dispatcher
	EngineClient          *EngineClient
	Logger                log.AdvancedLogger[any, sdklog.Logger]
	TelemetrySink         *metrics.TelemetrySink
}

// ProvideDepositService provides the deposit service to the depinject
// framework.
func ProvideDepositService(in DepositServiceIn) (*DepositService, error) {
	var finalizedBlkCh = make(chan *FinalizedBlockEvent)
	if err := in.Dispatcher.Subscribe(
		messages.BeaconBlockFinalizedEvent, finalizedBlkCh,
	); err != nil {
		in.Logger.Error("failed to subscribe to event", "event",
			messages.BeaconBlockFinalizedEvent, "err", err)
		return nil, err
	}

	// Build the deposit service.
	return deposit.NewService[
		*BeaconBlockBody,
		*BeaconBlock,
		*FinalizedBlockEvent,
		*DepositStore,
		*ExecutionPayload,
	](
		in.Logger.With("service", "deposit"),
		math.U64(in.ChainSpec.Eth1FollowDistance()),
		in.TelemetrySink,
		in.DepositStore,
		in.BeaconDepositContract,
		finalizedBlkCh,
	), nil
}
