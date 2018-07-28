all:
	build install

make:
	go build -o build/bk

install:
	ln -s ${CURDIR}/build/bk /usr/local/bin/bk > /dev/null
	ln -s ${CURDIR}/sh/jm.sh /usr/local/bin/jm > /dev/null

uninstall:
	rm /usr/local/bin/bk
	rm /usr/local/bin/jm
