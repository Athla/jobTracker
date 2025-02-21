# 🎯 JobTracker

A modern, streamlined job application tracking system built with Go and React. Track your job search journey with ease and style.

![JobTracker Dashboard](path_to_dashboard_screenshot.png)

## ✨ Features

- 🔐 **Secure Authentication**: JWT-based authentication system
- 📋 **Kanban Board**: Visual tracking of job applications across different stages
- 🔄 **Status Tracking**: Monitor applications from wishlist to offer/rejection
- 📱 **Responsive Design**: Seamless experience across all devices
- 🔍 **Search & Filter**: Quickly find specific applications
- 📊 **Analytics**: Track your application success rate
- 🗃️ **Persistent Storage**: SQLite database with automatic backups

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose
- Go 1.21+
- Node.js 16+
- Make (optional but recommended)

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/yourusername/jobTracker.git
cd jobTracker
```

2. **Set up environment variables**
```bash
cp .env.example .env
# Edit .env with your preferred settings
```

3. **Start the application**
```bash
# Using Make
make docker-run

# Or using Docker Compose directly
docker-compose up --build
```

4. **Access the application**
- Frontend: http://localhost:80
- Backend API: http://localhost:8080

### Default Credentials
```
Username: admin
Password: admin
```

## 🏗️ Architecture

```
jobTracker/
├── cmd/                    # Command line applications
├── frontend/              # React frontend application
├── internal/              # Internal packages
│   ├── auth/             # Authentication logic
│   ├── database/         # Database operations & migrations
│   ├── models/           # Data models
│   ├── server/           # HTTP server & handlers
│   └── utils/            # Utility functions
└── tests/                # Integration & unit tests
```

## 🛠️ Development

### Available Make Commands

```bash
make build          # Build the application
make run            # Run locally
make docker-run     # Run with Docker
make docker-down    # Stop Docker containers
make test           # Run tests
make watch          # Run with live reload
make migrate-up     # Run database migrations
make migrate-down   # Rollback last migration
make db-reset       # Reset database
make setup          # Initial setup
```

### API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST   | /login   | Authenticate user |
| GET    | /api/jobs | Get all jobs |
| POST   | /api/jobs | Create new job |
| PUT    | /api/jobs/:id | Update job |
| DELETE | /api/jobs/:id | Delete job |
| PATCH  | /api/jobs/:id/status | Update job status |

## 📚 Tech Stack

- **Backend**
  - Go
  - Gin (Web Framework)
  - SQLite
  - JWT Authentication

- **Frontend**
  - React
  - TypeScript
  - Tailwind CSS

- **DevOps**
  - Docker
  - Make

## 🔐 Security

- Passwords are hashed using bcrypt
- JWT tokens for authentication
- CORS protection
- SQL injection prevention
- Input validation
- Rate limiting

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [SQLite](https://www.sqlite.org/)
- [React](https://reactjs.org/)
- [Tailwind CSS](https://tailwindcss.com/)

## 📧 Contact

Guilherme "grimm" Rodrigues - [@grimmacez](https://twitter.com/grimmacez) - guilher.c.rodrigues@gmail.com

Project Link: [https://github.com/Athla/jobTracker](https://github.com/Athla/jobTracker)

## Known Issues

- Page Blinking
- Page Reloading
