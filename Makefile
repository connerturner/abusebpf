CC=clang
SOURCE=src
UNAME=$(shell uname -m)

.PHONY: generate
generate: clean
	$(CC) -g -O2 -I /usr/include/$(UNAME)-linux-gnu -target bpf -c $(SOURCE)/egress.bpf.c -o $(SOURCE)/egress.bpf.o

.PHONY: clean
clean:
	rm -rf $(SOURCE)/*.o

.PHONY: show-filters
show-filters:
	tc filter show dev $(DEVICE)

# teardown existing qdiscs and filters (use in development only)
teardown:
	tc filter del dev $(DEVICE)
	tc qdisc del dev $(DEVICE) clsact

# qdisc is needed to attach bpf filter onto and it must take clsact
.PHONY: load-filter create-qdisc
create-qdisc:
	tc qdisc add dev $(DEVICE) clsact 

load-filter: clean generate create-qdisc
	tc filter add dev $(DEVICE) egress bpf object-file src/egress.bpf.o section out_block_c2 da
