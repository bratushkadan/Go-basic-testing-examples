# config
repository := github.com/bratushkadan/testing_example
tested_dirs := util math db user

# functions
substitute_testing_dirs := $(foreach dir, $(tested_dirs), $(repository)/$(dir))
test_specified_dirs := go test $(substitute_testing_dirs)

# artifacts
coverage_filepath := coverage.out

PHONY: .test
test:
	$(test_specified_dirs)

PHONY: .test-concat-only
test-concat-only:
	go test $(repository)/util -run ^TestConcat

PHONY: .coverage
coverage:
	$(test_specified_dirs) -cover

PHONY: .coverprofile
coverprofile:
	$(test_specified_dirs) -coverprofile=$(coverage_filepath)

PHONY: .coverprofile-analyse
coverprofile-analyse: coverprofile
	go tool cover -html=$(coverage_filepath)