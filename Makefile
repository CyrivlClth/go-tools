GO ?= go
MODULE ?= github.com/CyrivlClth/go-tools
PACKAGES = "/" "/container/slice" "/idgen"

.PHONY: test
test:
	for d in $(PACKAGES); do \
  		$(GO) test -v $(MODULE)$$d; \
  	done
