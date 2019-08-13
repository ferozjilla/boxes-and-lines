.DEFAULT_GOAL = list 

.PHONY: test list

list:    ## list Makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

test:    ## run all tests
	pushd miro && \
	ginkgo . && \
	popd
