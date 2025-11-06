# Contact Hub

A simple full-stack web application for managing person contacts. Built with Go (Backend) and React (Frontend).

## Features

*   **Backend**: Go with Chi router, providing a RESTful API.
*   **Frontend**: React with Vite, TypeScript, and Material-UI.
*   **API**: Supports full-text search, date-range filtering, and server-side pagination.
*   **Containerized**: Fully containerized with Docker and Docker Compose for easy setup and deployment.

## Prerequisites

*   [Docker](https://www.docker.com/get-started)
*   [Docker Compose](https://docs.docker.com/compose/install/)

## Quick Start

1.  **Clone the repository:**
    ```bash
    git clone <your-repo-url>
    cd contact-hub
    ```

2.  **Run the application:**
    ```bash
    docker-compose up --build
    ```

3.  **Access the application:**
    *   The **Frontend** will be available at [http://localhost:5173](http://localhost:5173).
    *   The **Backend API** will be available at [http://localhost:8080](http://localhost:8080).

The application is now running. Any changes made to the `backend/data` directory will be reflected in the application after restarting the containers.