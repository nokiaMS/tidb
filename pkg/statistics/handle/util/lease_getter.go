// Copyright 2023 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"time"

	"go.uber.org/atomic"
)

/**
stats lease定义了自动更新任务的执行周期。
*/
// LeaseGetter is used to get the stats lease.
type LeaseGetter interface {
	//获得stats lease。
	// Lease returns the stats lease.
	Lease() time.Duration

	//设置stats lease。
	// SetLease sets the stats lease. Only used for test.
	SetLease(lease time.Duration)
}

var _ LeaseGetter = (*leaseGetter)(nil)

/*
lease getter对象结构体。
*/
type leaseGetter struct {
	lease *atomic.Duration
}

// 创建一个新的leaseGetter对象。
// NewLeaseGetter creates a new LeaseGetter.
func NewLeaseGetter(lease time.Duration) LeaseGetter {
	return &leaseGetter{
		lease: atomic.NewDuration(lease),
	}
}

// 获得stats lease值。
// Lease implements LeaseGetter.
func (g *leaseGetter) Lease() time.Duration {
	return g.lease.Load()
}

// 设置stats lease值。
// SetLease sets the stats lease. Only used for test.
func (g *leaseGetter) SetLease(lease time.Duration) {
	g.lease.Store(lease)
}
