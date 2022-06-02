migrate:
	migrate -source file://platform/database/postgres/migrations \
	-database postgres://root:password@127.0.0.1:5432/shopping-mono?sslmode=disable up;

migrate-force:
	@read -p "Enter the version of the migration: " version; \
	migrate -source file://platform/database/postgres/migrations \
	-database postgres://root:password@127.0.0.1:5432/shopping-mono?sslmode=disable force $$version; 

rollback:
	migrate -source file://platform/database/postgres/migrations \
    	-database postgres://root:password@127.0.0.1:5432/shopping-mono?sslmode=disable down;

drop:
	migrate -source file://platform/database/postgres/migrations \
        	-database postgres://root:password@127.0.0.1:5432/shopping-mono?sslmode=disable drop;

migration:
	@read -p "Enter the name of the migration: " name; \
		migrate create -ext sql -dir platform/database/postgres/migrations $$name;

sqlc:
	sqlc generate