# LeetCode em Go

Minhas soluções para problemas do [LeetCode](https://leetcode.com/u/JonasBorgesLM/), escritas em Go, cada uma com testes.

## Estrutura

```
problems/
  0001-two-sum/
    solution.go        # solução
    solution_test.go   # casos de teste (exemplos do enunciado + edge cases)
    README.md          # abordagem e complexidade
```

## Uso

```bash
# criar um novo problema
scripts/new.sh 1 two-sum easy

# rodar todos os testes
go test ./...

# rodar os testes de um problema
go test ./problems/0001-two-sum/
```

## Problemas resolvidos

| # | Problema | Dificuldade | Solução |
|---|----------|-------------|---------|
