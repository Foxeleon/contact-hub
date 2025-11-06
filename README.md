# Contact Hub

Welcome to Contact Hub, a full-stack web application designed to browse and manage contact information. This project serves as a robust example of a modern, containerized application built with a Go backend and a React frontend, featuring a clean architecture and a professional development workflow.

## Key Technologies

- **Backend**:
    - **Go**: For a fast, statically typed, and concurrent backend.
    - **Chi Router**: A lightweight, idiomatic, and composable router for building Go HTTP services.
- **Frontend**:
    - **React & TypeScript**: For building a modern, type-safe user interface.
    - **Vite**: A next-generation frontend tooling for an extremely fast development experience.
    - **Material-UI (MUI)**: For a rich set of pre-built, accessible UI components.
- **DevOps & Architecture**:
    - **Docker & Docker Compose**: For complete containerization, ensuring consistent environments and easy deployment.
    - **Nginx**: Used as a reverse proxy within the frontend container to manage API requests, eliminating CORS issues in production.
    - **Hybrid Environment**: Fully configured to support both Docker-based and local development workflows.

## Project Setup & Running

This application is designed to run in two distinct modes: a production-like Docker environment and a local development environment for active coding and debugging.

---

### **Mode 1: Production-like with Docker (Recommended)**

This is the simplest and most reliable way to run the entire application stack.

#### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

#### Quick Start

1.  **Clone the repository:**
    ```bash
    git clone <your-repo-url>
    cd contact-hub
    ```

2.  **Run the application:**
    From the root directory, execute:
    ```bash
    docker-compose up --build
    ```
    This command will build the images for both frontend and backend services and start the containers.

3.  **Access the application:**
    - The web interface will be available at **[http://localhost:5173](http://localhost:5173)**.

---

### **Mode 2: Local Development (For Active Development)**

This mode is ideal when you are actively working on either the frontend or backend and want to take advantage of features like Hot-Reload.

#### Prerequisites

- **Go** (version 1.25.4 recommended).
- **Node.js** (version 22 recommended) and **npm**.
- **Make** (for using backend helper commands).

#### Environment Configuration (Crucial First Step!)

The frontend requires environment variables to know where to find the API.

1.  Navigate to the `frontend/` directory.
2.  Create a file named `.env.local`. **This file should not be committed to Git.**
3.  Add the following content to `frontend/.env.local`:
    ```
    # Local development API endpoint
    VITE_API_URL_PERSONS=http://localhost:8080/api/persons
    ```
    *(Note: A base `.env` file with the production URL is already included in the repository for the Docker build).*

#### Running the Application

You will need two separate terminal windows.

1.  **Start the Backend:**
    - Navigate to the `backend/` directory.
    - Run the command:
      ```bash
      make run
      ```
    - The Go server will start on `http://localhost:8080` with CORS enabled for the frontend.

2.  **Start the Frontend:**
    - Navigate to the `frontend/` directory.
    - Run the commands:
      ```bash
      npm install
      npm run dev
      ```
    - The Vite development server will start, and the application will be available at **[http://localhost:5173](http://localhost:5173)**.

---

### Backend Helper Commands (`backend/`)

Navigate to the `backend/` directory to use these `make` commands:

- `make run`: Starts the development server with CORS enabled.
- `make test`: Runs all unit tests for the backend logic.
- `make tidy`: Tidies up the `go.mod` and `go.sum` files.