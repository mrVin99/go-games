rds-container:
	sudo docker run -d --name redis-container -p 6379:6379 redis:latest

pg-container:
	sudo docker run -d --name postgres-container -p 5432:5432 \
      -e POSTGRES_DB=postgres \
      -e POSTGRES_USER=postgres \
      -e POSTGRES_PASSWORD=postgres \
      -v /path/to/postgres_data:/var/lib/postgresql/data \
      postgres:latest

run-game:
	go build -o main && ./main