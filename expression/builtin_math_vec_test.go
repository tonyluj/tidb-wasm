// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package expression

import (
	"testing"

	. "github.com/pingcap/check"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/tidb/types"
)

var vecBuiltinMathCases = map[string][]vecExprBenchCase{
	ast.Log: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Log10: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Log2: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Sqrt: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Acos: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Asin: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Atan: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Atan2: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal, types.ETReal}},
	},
	ast.Cos: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Exp: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Degrees: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Cot: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Radians: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Sin: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Tan: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
	},
	ast.Abs: {
		{retEvalType: types.ETDecimal, childrenTypes: []types.EvalType{types.ETDecimal}},
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETInt}},
	},
	ast.Round: {
		{retEvalType: types.ETDecimal, childrenTypes: []types.EvalType{types.ETDecimal}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETInt}},
	},
	ast.Pow: {
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal, types.ETReal}, geners: []dataGenerator{&rangeRealGener{0, 10, 0.5}, &rangeRealGener{0, 100, 0.5}}},
	},
}

func (s *testEvaluatorSuite) TestVectorizedBuiltinMathEvalOneVec(c *C) {
	testVectorizedEvalOneVec(c, vecBuiltinMathCases)
}

func (s *testEvaluatorSuite) TestVectorizedBuiltinMathFunc(c *C) {
	testVectorizedBuiltinFunc(c, vecBuiltinMathCases)
}

func BenchmarkVectorizedBuiltinMathEvalOneVec(b *testing.B) {
	benchmarkVectorizedEvalOneVec(b, vecBuiltinMathCases)
}

func BenchmarkVectorizedBuiltinMathFunc(b *testing.B) {
	benchmarkVectorizedBuiltinFunc(b, vecBuiltinMathCases)
}
