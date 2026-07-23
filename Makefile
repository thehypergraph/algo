ODIN ?= odin

# Directories that contain at least one *_test.odin file.
TEST_DIRS := $(sort $(dir $(shell find . -name '*_test.odin' -not -path './.git/*')))

.PHONY: test build run clean

# Run every test package in the project.
test:
	@for dir in $(TEST_DIRS); do \
		echo "==> testing $$dir"; \
		$(ODIN) test "$$dir" || exit 1; \
	done

# Build the main executable.
build:
	$(ODIN) build . -out:algo

# Build and run the main program.
run:
	$(ODIN) run .

clean:
	rm -f algo
