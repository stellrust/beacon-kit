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

package beacondb

// GetLatestExecutionPayloadHeader retrieves the latest execution payload
// header from the BeaconStore.
func (kv *KVStore[
	ForkT, BeaconBlockHeaderT, ExecutionPayloadHeaderT, Eth1DataT, ValidatorT,
]) GetLatestExecutionPayloadHeader() (
	ExecutionPayloadHeaderT, error,
) {
	return kv.latestExecutionPayloadHeader.Get(kv.ctx)
}

// SetLatestExecutionPayloadHeader sets the latest execution payload header in
// the BeaconStore.
func (kv *KVStore[
	ForkT, BeaconBlockHeaderT, ExecutionPayloadHeaderT, Eth1DataT, ValidatorT,
]) SetLatestExecutionPayloadHeader(
	payloadHeader ExecutionPayloadHeaderT,
) error {
	return kv.latestExecutionPayloadHeader.Set(kv.ctx, payloadHeader)
}

// GetEth1DepositIndex retrieves the eth1 deposit index from the beacon state.
func (kv *KVStore[
	ForkT, BeaconBlockHeaderT, ExecutionPayloadT, Eth1DataT, ValidatorT,
]) GetEth1DepositIndex() (uint64, error) {
	return kv.eth1DepositIndex.Get(kv.ctx)
}

// SetEth1DepositIndex sets the eth1 deposit index in the beacon state.
func (kv *KVStore[
	ForkT, BeaconBlockHeaderT, ExecutionPayloadT, Eth1DataT, ValidatorT,
]) SetEth1DepositIndex(
	index uint64,
) error {
	return kv.eth1DepositIndex.Set(kv.ctx, index)
}

// GetEth1Data retrieves the eth1 data from the beacon state.
func (kv *KVStore[
	ForkT, BeaconBlockHeaderT, ExecutionPayloadT, Eth1DataT, ValidatorT,
]) GetEth1Data() (Eth1DataT, error) {
	return kv.eth1Data.Get(kv.ctx)
}

// SetEth1Data sets the eth1 data in the beacon state.
func (kv *KVStore[
	ForkT, BeaconBlockHeaderT, ExecutionPayloadT, Eth1DataT, ValidatorT,
]) SetEth1Data(
	data Eth1DataT,
) error {
	return kv.eth1Data.Set(kv.ctx, data)
}