// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package store

import (
	consensusv1 "github.com/itsdevbear/bolaris/types/consensus/v1"
)

type Deposit struct {
	*consensusv1.Deposit
}

// NewDeposit creates a new deposit.
func NewDeposit(pubkey []byte, amount uint64, withdrawalCredentials []byte) *Deposit {
	return &Deposit{
		Deposit: &consensusv1.Deposit{
			Pubkey:                pubkey,
			Amount:                amount,
			WithdrawalCredentials: withdrawalCredentials,
		},
	}
}

// commitDeposits commits the cached deposits to the queue.
func (s *BeaconStore) commitDeposits(depositCache []*Deposit) error {
	err := s.deposits.PushMulti(s.sdkCtx, depositCache)
	if err != nil {
		return err
	}
	return nil
}

// PersistDeposits commits the cached deposits to the queue
// and processes the queued deposits.
func (s *BeaconStore) PersistDeposits(depositCache []*Deposit, n uint64) ([]*Deposit, error) {
	var err error
	if len(depositCache) > 0 {
		if err = s.commitDeposits(depositCache); err != nil {
			return nil, err
		}
	}
	depositsToProcess, err := s.deposits.PopMulti(s.sdkCtx, n)
	if err != nil {
		return nil, err
	}
	for _, deposit := range depositsToProcess {
		// TODO: If an error occurs in the middle of processing deposits,
		// should we continue to process the remaining deposits?
		if err = s.processDeposit(deposit); err != nil {
			return nil, err
		}
	}
	return depositsToProcess, nil
}

// processDeposit processes a deposit with the staking keeper.
func (s *BeaconStore) processDeposit(deposit *Deposit) error {
	_, err := s.stakingKeeper.Delegate(s.sdkCtx, deposit)
	return err
}

// GetStakingNonce returns the latest staking nonce in the previous block.
// That nonce is also the expected staking nonce at the beginning of the current block.
func (s *BeaconStore) GetStakingNonce() (uint64, error) {
	headIdx, err := s.deposits.HeadIndex(s.sdkCtx)
	if err != nil {
		return 0, err
	}
	return headIdx, nil
}
