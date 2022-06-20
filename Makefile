CC=clang
SOURCE=src
UNAME=$(shell uname -m)

.PHONY: generate
generate: clean
	$(CC) -I /usr/include/$(UNAME)-linux-gnu -O2 -target bpf -c $(SOURCE)/egress.bpf.c -o $(SOURCE)/egress.bpf.o

.PHONY: clean
clean:
	rm -rf $(SOURCE)/*.o
