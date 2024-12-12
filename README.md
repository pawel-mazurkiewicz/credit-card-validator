# credit-card-validator
Micro-project with goal to create a microservice using Golang and deploy it on AWS

# Build and run

```
# Build the Docker image
docker build -t credit-card-validator:latest .

# Run the Docker container
docker run -p 8080:8080 credit-card-validator:latest
```