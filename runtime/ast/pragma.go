/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import "github.com/onflow/cadence/runtime/common"

// Pragma

type PragmaDeclaration struct {
	Expression Expression
	Range
}

func (*PragmaDeclaration) isDeclaration() {}

func (*PragmaDeclaration) isStatement() {}

func (p *PragmaDeclaration) Accept(visitor Visitor) Repr {
	return visitor.VisitPragmaDeclaration(p)
}

func (p *PragmaDeclaration) DeclarationIdentifier() *Identifier {
	return nil
}

func (p *PragmaDeclaration) DeclarationKind() common.DeclarationKind {
	return common.DeclarationKindPragma
}

func (p *PragmaDeclaration) DeclarationAccess() Access {
	return AccessNotSpecified
}