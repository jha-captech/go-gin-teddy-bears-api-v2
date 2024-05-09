GO := @go
GIN := @gin
PACKAGE := teddy_bears_api_v2

.PHONY: db_setup 
db_setup:
	sqlite3 -init database/teddy_bears_database_setup.sql database/teddy_bear.db " "

.PHONY: goinstall
goinstall:
	${GO} get .

.PHONY: http_dev
http_dev:
	swag init --parseDependency --dir "cmd\http\routes" --output "cmd\http\docs" 
	${GO} run ./cmd/http

.PHONY: http_test
http_test:
	${GO} test ${PACKAGE} -v

.PHONY: cli_dev_location_list
cli_dev_location_list:
	${GO} run ./cmd/cli location list

.PHONY: cli_dev_teddy_bear_list
cli_dev_location_list:
	${GO} run ./cmd/cli teddy_bear list