Postgre server should be lauched in docker with script:
docker run --name fat-postgres -e POSTGRES_PASSWORD=qwerty -d -p 5432:5432 postgres

Sql sripct to be run in database:
psql -h localhost -U postgres -d postgres -a -f schema/tables.sql


or if postgres server run in local machine:
psql -U postgres -d postgres -a -f schema/tables.sql