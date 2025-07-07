
# Media Repository 📸

This repository is part of the **Media** domain and contains microservices designed for managing and exporting media files. One of the primary services is **Go Media**, which facilitates exporting user profile photos. The repository is designed to handle media storage, management, and integration with other systems for processing and exporting images.

## Repository Link 📁
- [GitHub Repository](https://github.com/Loony213/Media)

## Purpose 🎯
The **Media Repository** provides the core infrastructure for handling user media files, focusing on the **Go Media** service that allows exporting user profile photos. This service makes it easy to manage media and integrate with other services that require image processing and storage.

## Architecture Style 🏗️
- **Microservice Architecture:** The repository follows a microservices architecture, enabling scalability and independent development and deployment of each service.
- **Design Pattern:** The system follows the **Service-Oriented Architecture (SOA)** pattern, where each microservice is responsible for a specific task, and communication between services is achieved through APIs.

## Technologies 💻
- **Programming Language:** Go (Golang)
- **Containerization:** Docker (optional)
- **Storage:** External media storage (S3, local server, etc.)
- **API Integration:** REST APIs for exporting and managing media

## Project Structure 🧑‍💻
The repository is structured as follows:

```
Media/
├── go-media-app/              # Contains the Go Media service for exporting user profile photos.
│                 # Main code that handles the media export process.
│
├── README.md                  # This file.
└── requirements.txt           # Python dependencies, if any.
```

### Folder Descriptions 📂
- **go-media-app/**: Contains the Go application for exporting user profile photos and managing media files.
- **requirements.txt**: Lists the dependencies for the repository, if needed.

## How to Deploy ⚙️
1. **Clone the Repository:**
   ```bash
   git clone https://github.com/Loony213/Media.git
   ```

2. **Install Dependencies:**
   Navigate to the project directory and install the necessary dependencies (if any). For Go-based projects, you may need to set up the Go environment.
   ```bash
   go mod tidy
   ```

3. **Run the Go Media Service:**
   - Once dependencies are set up, run the Go Media service to start the export functionality:
     ```bash
     go run main.go
     ```

4. **Docker Deployment:**
   - Build the Docker image:
     ```bash
     docker build -t kamartinez/go-media .
     ```
   - Run the container:
     ```bash
     docker run -p 5000:5000 kamartinez/go-media
     ```

5. **Access the Service:**
   - The Go Media service will be available on `http://localhost:5000` once the container is running.

## Features ✨
- **Export User Profile Photos**: Facilitates exporting profile photos for users from different systems.
- **Modular and Scalable**: Easily scalable for different media management use cases.
- **Integration Ready**: Can be integrated with other systems that require media processing or exporting.

## License 📜
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
