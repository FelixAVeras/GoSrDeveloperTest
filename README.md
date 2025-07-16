🧠 Golang Backend Developer Training - Exercises + Technical Assessment
This content is designed to help you prepare for technical interviews or real-world roles as a Senior Golang Developer, specifically in areas related to microservices, AWS, clean code architecture, CI/CD, testing, and event-driven systems.

🔍 Overview
This project consists of:

7 standalone backend exercises, each focusing on key concepts required for a senior backend role in Go.

1 technical assessment simulation, modeled after real company evaluations.

Each part is crafted to reflect scenarios you might encounter in a company that builds scalable backend services using Go and cloud infrastructure like AWS.

✅ List of Exercises
1. Basic CRUD Microservice
Build a basic REST API for product management using net/http or a lightweight framework (e.g., Gin).

Store data in memory using Go maps or slices.

Apply input validation and clear HTTP responses.

2. Hexagonal Architecture Refactor
Refactor the CRUD logic into domain, service, handler, and infrastructure layers.

Focus on decoupling logic from delivery and storage layers.

Follow clean code and dependency inversion principles.

3. Simulated AWS S3 Upload
Use the AWS Go SDK or MinIO to simulate uploading files to S3.

Use environment variables for credentials.

Log upload success or failure.

4. Canary Release & API Gateway Simulation
Create two versions of a service (v1 and v2).

Route a portion of requests to v2 to simulate canary deployment.

Use a simple Go gateway or reverse proxy (e.g., NGINX or Go’s http.ServeMux) to simulate traffic splitting.

5. Event-Driven Architecture with Channels
Simulate an event-driven system using Go channels.

Create a producer that emits events like OrderCreated.

Implement a consumer that listens to and processes events.

6. CI/CD Pipeline with GitHub Actions
Create a Makefile to simplify build, test, and run commands.

Add a .github/workflows/go.yml file to run tests automatically on push.

(Optional) Add Go linting and coverage reporting.

7. Unit Testing
Write unit tests for business logic using Go’s testing package and/or testify.

Cover edge cases and expected failures.

Aim for at least 80% code coverage.
