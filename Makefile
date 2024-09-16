.PHONY: run
run:
	$(info Running...)
	go run ./cmd

.PHONY: go-jet
go-jet:
	rm -rf internal/repository/tables
	mkdir internal/repository/tables
	jet -dsn=postgresql://postgres:Usx6YSbXyJ54FnW7w2pueNfmR@localhost:5432/tenders?sslmode=disable -schema=public -path=./internal/repository/tables