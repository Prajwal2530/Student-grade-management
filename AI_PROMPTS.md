# AI Prompts Summary: Student Grade Management System API

This document provides a structured overview of the prompts utilized during the generation, refinement, and optimization of the Student Grade Management System API.

## Initial Project Generation Prompt

**Prompt:**
> "Generate a RESTful API in Go for a Student Grade Management System. Include basic CRUD operations for students and their grades. Structure the project using standard REST patterns, including separate routing and controller layers. Use Go modules for dependency management."

**Summary:**
This prompt instructed the AI to bootstrap the foundational architecture of the Go-based backend application. It established the core entity relationships (students and grades), defined the required API endpoints, and set up a standard project directory structure ensuring modularity.

## Role & Database Design Prompt

**Prompt:**
> "Extend the current Student Grade Management System API to include role-based access control (RBAC) with 'Admin' and 'Teacher' roles. Integrate a PostgreSQL database wrapper using GORM. Provide the complete database schema definitions and update the CRUD operations to use the database context."

**Summary:**
This prompt directed the AI to transition the application from in-memory or mock data structures to a persistent relational database model. It also introduced security and authorization by defining specific user roles and outlining the necessary database schema for subjects, roles, and grades.

## Docker Integration Prompt

**Prompt:**
> "Create a `Dockerfile` and `docker-compose.yml` for the Go API. The Docker Compose configuration should spin up the Go application along with a PostgreSQL database instance. Ensure the network allows secure communication between the app container and the database container."

**Summary:**
This prompt tasked the AI with containerizing the application for standardized deployment. It required the configuration of both the application environment and its dependent database service, ensuring seamless orchestration and internal network accessibility through Docker Compose.

## Environment Variable Refactoring Prompt

**Prompt:**
> "Refactor the hardcoded configuration values within the application to use environment variables. Implement a `.env` file loader using a library like `godotenv`. Ensure database credentials, server ports, and any API specific secrets are loaded dynamically at runtime."

**Summary:**
This prompt guided the AI to improve application security and configurability. It moved sensitive credentials and environment-specific settings out of the source code, bringing the application in alignment with standard twelve-factor app methodologies.

## Postman Collection Generation Prompt

**Prompt:**
> "Generate a complete Postman collection in JSON format for the Student Grade Management System API. Detail all available endpoints, including required headers, path variables, query parameters, and example JSON request/response bodies for both success and error states."

**Summary:**
This prompt requested the creation of standard API documentation and testing assets. It ensured that all routes, payload structures, and expected behaviors were comprehensively modeled, enabling streamlined endpoint verification and front-end integration testing.

## Code Simplification / Medium-Level Instruction Prompt

**Prompt:**
> "Review the current routing and database handler functions. Simplify any redundant error-checking logic, consolidate duplicated boilerplate code into reusable middleware or utility functions, and ensure consistent JSON response formatting across all controllers."

**Summary:**
This prompt focused the AI on code quality and maintainability. It initiated a refactoring pass to remove technical debt, improve error handling consistency, and apply DRY (Don't Repeat Yourself) principles across the application's core logic layers.
