// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package gobackend_test

import (
	"testing"

	"github.com/gomlx/compute/internal/gobackend"
	"github.com/gomlx/compute/support/testutil"
)

// TestOptimizationBarrier verifies the Go backend returns its operands unchanged:
// the barrier is a compile-time constraint only, a no-op on values.
func TestOptimizationBarrier(t *testing.T) {
	builder := backend.Builder("test_optimization_barrier")
	mainFn := builder.Main()

	a, err := mainFn.Constant([]float32{1, 2, 3}, 3)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	b, err := mainFn.Constant([]float32{4, 5}, 2)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	results, err := mainFn.OptimizationBarrier(a, b)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if err = mainFn.Return(results, nil); err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	exec, err := builder.Compile()
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	outputs, err := exec.Execute(nil, nil, 0)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	if len(outputs) != 2 {
		t.Fatalf("expected 2 outputs, got %d", len(outputs))
	}
	if ok, diff := testutil.IsEqual([]float32{1, 2, 3}, outputs[0].(*gobackend.Buffer).Flat); !ok {
		t.Errorf("operand 0 changed by barrier:\n%s", diff)
	}
	if ok, diff := testutil.IsEqual([]float32{4, 5}, outputs[1].(*gobackend.Buffer).Flat); !ok {
		t.Errorf("operand 1 changed by barrier:\n%s", diff)
	}
}
