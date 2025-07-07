
# Go Media App Microservice 📸

This microservice is part of the **Media** domain and is responsible for storing user profile photos in an S3 bucket. The service interacts with Amazon S3 to upload, store, and retrieve profile images, providing a reliable and scalable solution for media storage.

## Repository Link 📁
- [GitHub Repository](https://github.com/Loony213/Media)

## Docker Image 🐳
- **Docker Image:** `kamartinez/mediago`

## Purpose 🎯
The **Go Media App** microservice provides an efficient way to store and manage user profile photos in Amazon S3. The service allows for scalable, secure, and accessible media storage for user images.

## Architecture Style 🏗️
- **Microservice Architecture:** The service is designed to be a standalone microservice that interacts with other components through APIs, enabling scalability and independent deployment.
- **Design Pattern:** This service follows a **Modular Architecture** with a clear separation of concerns, making it maintainable and easily extensible.

## Technologies 💻
- **Programming Language:** Go (Golang)
- **Containerization:** Docker
- **Storage:** Amazon S3 for media file storage
- **API Integration:** REST APIs for interacting with the media storage service

## Project Structure 🧑‍💻
The repository is structured as follows:

```
go-media-app/
├── cmd/
│   └── server/
│       └── main.go            # Main entry point to start the Go Media service.
│
├── internal/                  # Contains the core application logic.
│   └── profile/               # Handles user profile image interactions.
│       ├── handler.go         # Handles incoming requests related to profile images.
│       └── service.go         # Defines the service to interact with the media storage.
│   └── s3utils/               # S3-related utilities to interact with Amazon S3.
│       └── s3.go              # Manages S3 interactions, including upload and download.
│
├── pkg/                       # Shared packages for the application.
│   └── logger/                # Handles logging functionalities.
│       └── logger.go          # Implements custom logging functionality.
│
├── go.mod                     # Go module dependencies.
├── go.sum                     # Go checksum file for modules.
└── Dockerfile                 # Docker configuration for containerization.
```

### Folder Descriptions 📂
- **cmd/server/**: Contains the main entry point for the Go Media service.
- **internal/profile/**: Handles the logic related to user profile photos, including request handling and service functions.
- **internal/s3utils/**: Contains utilities for interacting with Amazon S3 to upload and retrieve media files.
- **pkg/logger/**: Provides custom logging utilities for the service.
- **go.mod/go.sum**: Go module files for managing dependencies.
- **Dockerfile**: Configuration file for building the Docker image of the service.

## How to Deploy ⚙️
1. **Clone the Repository:**
   ```bash
   git clone https://github.com/Loony213/Media.git
   ```

2. **Install Dependencies:**
   Navigate to the project directory and install the necessary Go modules:
   ```bash
   go mod tidy
   ```

3. **Run the Go Media Service:**
   - Once dependencies are set up, run the Go Media service:
     ```bash
     go run cmd/server/main.go
     ```

4. **Docker Deployment:**
   - Build the Docker image:
     ```bash
     docker build -t kamartinez/mediago .
     ```
   - Run the container:
     ```bash
     docker run -p 5000:5000 kamartinez/mediago
     ```

5. **Access the Service:**
   - The Go Media service will be available on `http://localhost:5000` once the container is running.

## Features ✨
- **Profile Photo Storage**: Stores user profile images in Amazon S3.
- **S3 Integration**: Secure and scalable storage of media files using Amazon S3.
- **Modular Architecture**: Easy to scale and extend with additional features.
- **Logging**: Implements custom logging for better debugging and monitoring.

## License 📜
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
