all: build

SUBDIRS = $(shell find . -maxdepth 1 -type d|grep service|sort -n)

.PHONY: mod
mod:
	$(call make_exec, mod)
.PHONY: clean
clean:
	$(call make_exec, clean)
.PHONY: build
build:
	$(call make_exec, build)
.PHONY: proto
proto:
	$(call make_exec, proto)
.PHONY: up
up:
	$(call make_exec, up)

define make_exec
	@for dir in $(SUBDIRS) ; do \
		$(MAKE) $1 -C $$dir; \
		done;
endef
