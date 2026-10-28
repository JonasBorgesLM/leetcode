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
make                              # lista os comandos
make new N=1 S=two-sum D=easy     # cria problems/0001-two-sum/
make test N=1                     # testa só esse problema
make test                         # testa tudo
make check                        # gofmt + go vet + testes, antes do commit
make list                         # problemas já criados
```

## Problemas resolvidos

| # | Problema | Dificuldade | Solução |
|---|----------|-------------|---------|
