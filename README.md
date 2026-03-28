# Go CI/CD & Unit Testing Project

![Go Version](https://img.shields.io/badge/go-v1.24.x-blue?style=flat-square&logo=go)
![Build Status](https://img.shields.io/badge/CI%2FCD-Success-brightgreen?style=flat-square&logo=github-actions)
![Platform](https://img.shields.io/badge/platform-EC2-orange?style=flat-square&logo=amazon-aws)

This project is a Go-based web application designed as a learning ground for **CI/CD pipelines**, **Unit Testing**, and **Automated Deployment** using GitHub Actions. It features a simple authentication system (Signup, Login, Logout) built with the Gin framework and GORM.

## 🚀 Features

- **User Authentication**: Secure Signup, Login, and Logout functionality.
- **RESTful API**: Built using the [Gin Gonic](https://gin-gonic.com/) web framework.
- **Database Integration**: Uses [GORM](https://gorm.io/) with SQLite for data persistence.
- **Unit Testing**: Comprehensive test suite ensuring code reliability.
- **CI/CD Pipeline**: Fully automated build and deployment process.

---

## 🛠️ Tech Stack

- **Backend**: Go (Golang)
- **Framework**: Gin Gonic
- **ORM**: GORM (Go Object Relational Mapper)
- **Database**: SQLite (In-memory for testing, file-based for production)
- **Testing**: standard `testing` package with [Testify](https://github.com/stretchr/testify) assertions.
- **Infrastructure**: AWS EC2
- **Automation**: GitHub Actions

---

## 🏗️ Project Structure

```text
├── .github/workflows/   # CI/CD pipeline configurations
├── config/              # Database and environment configurations
├── controllers/         # API business logic handlers
├── models/              # GORM database models
├── routes/              # API route definitions
├── test/             # Unit and Integration tests
├── main.go              # Application entry point
└── go.mod               # Dependency management
```

---

## 🧪 Unit Testing

The project emphasizes robust testing practices. Tests are located in the `test/` directory and use a mock/in-memory SQLite database to ensure tests are isolated and fast.

### Running Tests Locally

To execute the test suite, run:

```bash
go test -v ./test/...
```

**What's tested?**
- `userSignup_test.go`: Validates user registration, duplicate checks, and field validation.
- `userLogin_test.go`: Verifies authentication logic and session handling.
- `userLogout_test.go`: Ensures safe session termination.

---

## ⚙️ CI/CD Workflow (GitHub Actions)

The repository includes a sophisticated CI/CD pipeline defined in `.github/workflows/deploy.yml`. 

### Pipeline Stages:

1.  **Build Phase**:
    *   Sets up the Go environment (v1.24.x).
    *   Runs **Unit Tests** to ensure quality.
    *   Compiles the code into a **Linux binary** (`GOOS=linux GOARCH=amd64`).
    *   Uploads the binary as a build artifact.

2.  **Deployment Phase (CD)**:
    *   Triggered automatically upon a successful build on the `main` branch.
    *   Securely connects to an **AWS EC2** instance via SSH.
    *   Deploys the new binary to the server.
    *   Restarts the application service using `systemctl`.

### Required GitHub Secrets:

To enable the deployment, the following secrets must be configured in your GitHub repository settings:
- `SSH_PRIVATE_KEY`: Your EC2 instance private key.
- `SSH_HOST`: Public IP or DNS of the EC2 instance.
- `SSH_USER`: The username for SSH (e.g., `ubuntu`).
- `REMOTE_DIR`: Directory where the application resides on the server.
- `SERVICE_NAME`: The name of the `systemd` service for your app.

---

## 🚦 Getting Started

### Prerequisites

- Go (version 1.24 or higher)

### Setup

1.  **Clone the repository**:
    ```bash
    git clone <repository-url>
    cd Unit-Testing
    ```

2.  **Install dependencies**:
    ```bash
    go mod tidy
    ```

3.  **Run the application**:
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

---

## 📄 License

This project is for educational purposes. Feel free to use and modify it for your learning journey!
