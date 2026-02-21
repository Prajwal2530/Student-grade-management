# System Design Document

## Architecture Overview
The system is divided into two main components:
1. **Backend API**: A RESTful service built with Go (Golang), Gin web framework, and GORM. It provides core business logic, handles authentication (JWT), and interfaces with the database.
2. **Frontend UI** *(To be developed)*: A user interface that will consume the backend REST API.
3. **Database**: PostgreSQL 15, storing all relational data (Users, Courses, Enrollments, Grades).

## Entity-Relationship (ER) Diagram
![ER Diagram](er-diagram.png)

*The database schema enforces data integrity using unique constraints to prevent duplicate enrollments and duplicate grades for the same student in the same course.*

## Component Responsibilities

### Backend Endpoints
- **Public**: Health checks, Registration (Student only), and Login.
- **Admin**: User creation (elevation to Teacher or Admin), Course creation, and global listing.
- **Teacher**: Course management, Student enrollment, and Grading.
- **Student**: View enrolled courses, grades, and GPA calculation.

### Authorization Flow
- Clients authenticate via `/login` and receive a JWT.
- JWTs contain the `user_id` and `role`.
- Subsequent requests must pass the JWT in the `Authorization: Bearer <token>` header.
- Middleware intercepts requests to validate token signature and enforce Role-Based Access Control (RBAC).

## Infrastructure
The entire stack is containerized using Docker Compose. The `docker-compose.yml` orchestrates:
- The `postgres` container with persistent volumes.
- The `grade_api` backend container, built from the `backend/Dockerfile` and dependent on the database. 
