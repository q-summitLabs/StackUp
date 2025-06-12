# Backend Structure
This document provides an overview of the backend structure for the application, detailing the organization of directories and their purposes.


##  Directory Structure

cmd:
  - Contains the main application entry points.
  - Each subdirectory corresponds to a specific service or functionality.

    eg. server, worker, etc.


internal:
  - Contains the core business logic and application code.
  - Organized into packages for modularity.
  - Unexported packages are used internally within the application.

    eg. Database query logic, authentication, handlers, etc.

pkg:
  - Contains reusable packages that can be imported by other applications.
  - These packages are designed to be shared and used across different projects.

    eg. Utility functions, common data structures, etc.
