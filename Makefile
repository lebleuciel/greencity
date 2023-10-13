# Build Producer Docker image
build-producer:
	docker build -t greencity/producer:latest .

# Build Consumer Docker image
build-consumer:
	docker build -t greencity/consumer:latest .

# Run Docker Compose
run-compose:
	docker-compose up -d

# Stop and remove Docker Compose containers
down-compose:
	docker-compose down

# Clean up images and volumes
clean:
	docker system prune -a --volumes --force
