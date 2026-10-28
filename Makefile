# Atalhos do dia a dia. Rode `make` (ou `make help`) para ver a lista.
#
# Variáveis:
#   N  número do problema   (ex.: N=1)
#   S  slug do problema     (ex.: S=two-sum)
#   D  dificuldade          (ex.: D=easy)

.DEFAULT_GOAL := help
.PHONY: help new test check fmt list

# Pasta do problema N (problems/0001-*), vazia se N não foi informado.
dir = $(if $(N),$(wildcard problems/$(shell printf "%04d" $(N))-*))

help: ## mostra os comandos disponíveis
	@grep -E '^[a-z]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  make %-8s %s\n", $$1, $$2}'
	@echo
	@echo "  ex.: make new N=1 S=two-sum D=easy"
	@echo "       make test N=1"

new: ## cria um problema (N=, S=, D= opcional)
	@test -n "$(N)" -a -n "$(S)" || { echo "uso: make new N=1 S=two-sum D=easy"; exit 1; }
	@scripts/new.sh $(N) $(S) $(D)

test: ## roda os testes (todos, ou só do problema N=)
ifdef N
	@test -n "$(dir)" || { echo "problema $(N) não encontrado em problems/"; exit 1; }
	go test -v ./$(dir)/
else
	go test ./...
endif

check: fmt ## formata, roda go vet e todos os testes (use antes do commit)
	go vet ./...
	go test ./...

fmt: ## formata o código com gofmt
	gofmt -w .

list: ## lista os problemas já criados
	@ls problems | grep -v '^\.' || echo "nenhum problema ainda"
