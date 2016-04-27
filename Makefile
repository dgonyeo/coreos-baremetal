
BIN_DIR=/usr/local/bin
DATA_DIR=/var/lib/bootcfg
ENV_FILE=/etc/bootcfg.env

all: build

build:
	./build

test:
	./test

setup:
	id -u bootcfg > /dev/null 2>&1 || useradd -U bootcfg
	mkdir -p ${DATA_DIR}/{profiles,groups,ignition,cloud,assets}
	chgrp -R bootcfg ${DATA_DIR}
	chmod -R g+w ${DATA_DIR}
	touch ${ENV_FILE}

teardown:
	userdel bootcfg

install: setup
	cp bin/bootcfg $(BIN_DIR)
	cp bin/bootcmd $(BIN_DIR)
	@echo "**************"
	@echo "INSTALL SUCESS"
	@echo "**************"
	@echo "bootcfg was installed to /usr/local/bin/bootcfg"
	@echo "bootcmd was installed to /usr/local/bin/bootcmd"
	@echo "Default data directory is located at /var/lib/bootcfg"

uninstall:
	rm $(BIN_DIR)/bootcfg
	rm $(BIN_DIR)/bootcmd

.PHONY: build
