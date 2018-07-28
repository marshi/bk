.PHONY: build
build:
	go build -o ${CURDIR}/build/bk

.PHONY: install
install:
	ln -s ${CURDIR}/build/bk /usr/local/bin/bk
	ln -s ${CURDIR}/sh/jm.sh /usr/local/bin/jm

.PHONY: clean
clean:
	rm /usr/local/bin/bk
	rm /usr/local/bin/jm
