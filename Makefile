all: make install

make:
	go build -o ${CURDIR}/build/bk

install: uninstall
	ln -s ${CURDIR}/build/bk /usr/local/bin/bk
	ln -s ${CURDIR}/sh/jm.sh /usr/local/bin/jm

uninstall:
	rm /usr/local/bin/bk
	rm /usr/local/bin/jm
