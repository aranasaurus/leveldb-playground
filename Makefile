BINDIR=bin

leveldb-playground:
	go build -o ${BINDIR}/$@

.PHONY: run
run: leveldb-playground
	${BINDIR}/$^

.PHONY: clean
clean:
	go clean
	rm ${BINDIR}/*
