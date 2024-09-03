create migrate:
	migrate create -ext sql -dir db/migrations/sso -seq sso
up migrate:
	migrate -path ./db/migrations/sso -database 'postgres://postgres:root@localhost:5432/sso?sslmode=disable' up
down migrate:
	migrate -path ./db/migrations/sso -database 'postgres://postgres:root@localhost:5432/sso?sslmode=disable' down
gen protoc:
	protoc -I proto proto/sso/sso.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
create test migrate:
	migrate create -ext sql -dir tests/migrations -seq apps