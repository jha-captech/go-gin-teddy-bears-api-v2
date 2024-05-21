GO := @go
GIN := @gin
PACKAGE := teddy_bears_api_v2

# SETUP

.PHONY: db 
db:
	sqlite3 -init database/teddy_bears_database_setup.sql database/teddy_bear.db " "


.PHONY: goinstall
goinstall:
	${GO} get .


# GIN


.PHONY: gin_dev
gin_dev:
	swag init --parseDependency --generalInfo ".\..\main.go" --dir "cmd\gin\routes" --output "cmd\gin\docs" 
	${GO} run ./cmd/gin


.PHONY: gin_test
gin_test:
	${GO} test ${PACKAGE} -v


# FIBER


.PHONY: fiber_dev
fiber_dev:
	swag init --parseDependency --generalInfo ".\..\main.go" --dir "cmd\fiber\routes" --output "cmd\fiber\docs" 
	${GO} run ./cmd/fiber


# CHI


.PHONY: chi_dev
chi_dev:
	swag init --parseDependency --generalInfo ".\..\main.go" --dir "cmd\chi\routes" --output "cmd\chi\docs" 
	${GO} run ./cmd/chi


# CLI


.PHONY: cli_dev_location_list
cli_dev_location_list:
	${GO} run ./cmd/cli location list


.PHONY: cli_dev_teddy_bear_list
cli_dev_teddy_bear_list:
	${GO} run ./cmd/cli teddy_bear list