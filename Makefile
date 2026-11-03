# Directories containing Odin packages (auto-discovered: any dir with .odin files)
PACKAGES := $(sort $(dir $(wildcard */*.odin)))

.PHONY: test
test:
	@for pkg in $(PACKAGES); do \
		echo "=== Testing $$pkg ==="; \
		odin test $$pkg || exit 1; \
	done
	@echo "All tests passed."
