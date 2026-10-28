#!/usr/bin/env bash
# Cria a pasta de um novo problema com solution.go e solution_test.go.
#
# Uso: scripts/new.sh <numero> <slug> [dificuldade]
# Ex.: scripts/new.sh 1 two-sum easy
set -euo pipefail

if [ $# -lt 2 ]; then
  echo "uso: $0 <numero> <slug> [easy|medium|hard]" >&2
  exit 1
fi

num=$(printf "%04d" "$1")
slug=$2
difficulty=${3:-}
dir="problems/${num}-${slug}"
# Nome de pacote Go não aceita hífen nem começar com dígito: "two-sum" -> "twosum".
pkg=$(echo "$slug" | tr -d '-' | tr '[:upper:]' '[:lower:]')

cd "$(dirname "$0")/.."

if [ -e "$dir" ]; then
  echo "já existe: $dir" >&2
  exit 1
fi

mkdir -p "$dir"

cat > "$dir/solution.go" <<EOF
// Package $pkg resolve o problema $1 do LeetCode.
// https://leetcode.com/problems/$slug/
package $pkg

// TODO: cole aqui a assinatura da função do LeetCode e implemente.
EOF

cat > "$dir/solution_test.go" <<EOF
package $pkg

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		// TODO: entradas e saída esperada
	}{
		// {name: "exemplo 1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: chamar a solução e comparar com o esperado
		})
	}
}
EOF

cat > "$dir/README.md" <<EOF
# $1. $slug

- Link: https://leetcode.com/problems/$slug/
- Dificuldade: ${difficulty:-?}

## Abordagem

<!-- Explique a ideia da solução. -->

## Complexidade

- Tempo: O(?)
- Espaço: O(?)
EOF

echo "criado: $dir"
echo "não esqueça de adicionar a linha na tabela do README.md principal"
