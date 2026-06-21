// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package gobackend

import "github.com/gomlx/compute"

// OptimizationBarrier is an identity on the Go (SimpleGo) backend: it evaluates
// eagerly with no compiler optimization pass to fence, so the operands pass
// through unchanged. The barrier matters only for ahead-of-time compilers (XLA),
// where it prevents CSE of recomputed subgraphs; on the Go backend there is no CSE
// to prevent, so returning the operands as-is is both correct and sufficient.
func (f *Function) OptimizationBarrier(operands ...compute.Value) ([]compute.Value, error) {
	return operands, nil
}
