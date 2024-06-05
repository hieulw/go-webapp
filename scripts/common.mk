ifndef __COMMON_MK
__COMMON_MK = True

# Load dotenv file
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

# Check command exists
cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)

# Make default target show help
.DEFAULT_GOAL:=help
.PHONY: help
help: ## Display this help message
	@awk '/^[a-zA-Z\-\\_0-9]+:/ {                         \
		nb = sub( /^## /, "", helpMsg );                    \
		if(nb == 0) {                                       \
			helpMsg = $$0;                                    \
			nb = sub( /^[^:]*:.* ## /, "", helpMsg );         \
		}                                                   \
		if (nb)                                             \
			printf "\033[32m%s\033[0m%s\n", $$1, helpMsg;     \
	}                                                     \
	{ helpMsg = $$0 }'                                    \
	$(MAKEFILE_LIST) | column -ts ":"

endif
