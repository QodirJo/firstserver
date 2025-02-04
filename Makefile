CURRENT_DIR=$(shell pwd)
DBURL='postgres://postgres:1111@localhost:5432/auth?sslmode=disable'

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path migrations -database $(DBURL) -verbose up

mig-down:
	migrate -path migrations -database $(DBURL) -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table

swag-gen:
	~/go/bin/swag init -g ./api/api.go -o api/docs force 1

