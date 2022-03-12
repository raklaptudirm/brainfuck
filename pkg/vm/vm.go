// Copyright © 2021 Rak Laptudirm <raklaptudirm@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vm

import (
	"errors"
	"fmt"

	"github.com/raklaptudirm/brainfuck/pkg/ast"
	"github.com/raklaptudirm/brainfuck/pkg/token"
)

type Machine struct {
	Length  int
	Memory  []byte
	Pointer int
}

func New(l int) *Machine {
	return &Machine{
		Length: l,
		Memory: make([]byte, l),
	}
}

var ErrAst = errors.New("invalid ast structure")

func (m *Machine) Execute(n ast.Node) error {
	switch v := n.(type) {
	case *ast.Program:
		return m.executeOperations(v.Operations)
	case *ast.Comment:
		return nil // ignore comments
	case *ast.Loop:
		if m.Memory[m.Pointer] == 0 {
			return nil
		}

		for {
			err := m.executeOperations(v.Operations)
			if err != nil {
				return err
			}

			if m.Memory[m.Pointer] == 0 {
				return nil
			}
		}
	case *ast.Operator:
		return m.executeOperator(v)
	default:
		return ErrAst
	}
}

func (m *Machine) executeOperations(operations []ast.Operation) error {
	for _, operation := range operations {
		err := m.Execute(operation)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Machine) executeOperator(o *ast.Operator) error {
	switch o.Token {
	case token.INC_VAL:
		m.Memory[m.Pointer]++
	case token.DEC_VAL:
		m.Memory[m.Pointer]--
	case token.INC_PTR:
		m.Pointer++
	case token.DEC_PTR:
		m.Pointer--
	case token.INPUT:
		_, err := fmt.Scanf("%c", &m.Memory[m.Pointer])
		return err
	case token.PRINT:
		fmt.Print(string(m.Memory[m.Pointer]))
	default:
		return ErrAst
	}

	return nil
}
