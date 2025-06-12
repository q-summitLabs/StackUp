# StackedUp
Central banking application for tracking user's budgets and spending

# Getting Started

# Backend Development
This document provides an overview of the backend structure for the application, detailing the organization of directories and their purposes.

## Setup for Backend Development (local)
1. Install brew if you haven't already:
    ```bash
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    ```
2. Install Go:
    ```bash
    brew install go 1.24   
    ``` 

3. Install sqlc:
    ```bash
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest 
    ```

    This is for generating Go code from SQL queries.

4. Install goose
    ```
    go install github.com/pressly/goose/v3/cmd/goose@latest
    ```
    This is for managing database migrations.

5. Install Air
    ```bash
    go install github.com/air-verse/air@latest
    ```

    Air is a live reloading tool for Go applications, which helps in development by automatically restarting the server when code changes are detected.


6. Install Postgres
    TBD


# Frontend Development
This document provides an overview of the frontend structure for the application, detailing the organization of directories and their purposes.

## Setup for Frontend Development (local)

1. Install Node.js and npm:
    ```bash
    brew install node
    ```

2. Install yarn:
    ```bash
    npm install --global yarn
    ```
3. Install dependencies:
    ```bash
    yarn install
    ```

4. Start the development server:
    ```bash
    yarn dev
    ```