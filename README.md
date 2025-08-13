# Internship Hub

A full-stack web application for tracking internship applications, built with Go (backend) and SvelteKit (frontend).

## 🚀 Features

- **User Authentication**: Secure registration and login with JWT tokens
- **Application Management**: Create, read, update, and delete internship applications
- **File Upload**: Upload and manage resume/cover letter files
- **Responsive UI**: Modern, mobile-friendly interface
- **Real-time Updates**: Dynamic application status tracking

## 🛠️ Tech Stack

- **Backend**: Go with Gin framework
- **Frontend**: SvelteKit with TypeScript
- **Database**: PostgreSQL
- **Authentication**: JWT tokens
- **File Storage**: Local file system
- **Containerization**: Docker & Docker Compose

## 📋 Prerequisites

- **Go** 1.19+ 
- **Node.js** 18+
- **Docker** and **Docker Compose**
- **Git**

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/JeremiasZimmerman213/internship_hub_GO.git
cd internship_hub_GO
```

### 2. Environment Setup

#### Backend Environment
```bash
cd backend
cp .env.example .env
```

Edit `backend/.env` with your configuration:
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5433
DB_USER=your_db_user
DB_PASSWORD=your_secure_password
DB_NAME=internship_tracker

# JWT Configuration  
JWT_SECRET=your_super_secret_jwt_key_here

# Server Configuration
PORT=8080

# Environment
GIN_MODE=debug
```

#### Frontend Environment
```bash
cd frontend
cp .env.example .env
```

Edit `frontend/.env`:
```env
# For local development
VITE_API_BASE_URL=http://localhost:8080

# For network access (replace with your IP)
# VITE_API_BASE_URL=http://YOUR_IP_ADDRESS:8080
```

#### Docker Environment
```bash
# In project root
cp .env.docker.example .env.docker
```

Edit `.env.docker`:
```env
POSTGRES_USER=your_db_user
POSTGRES_PASSWORD=your_secure_password
POSTGRES_DB=internship_tracker
```

### 3. Start the Database

```bash
docker-compose up -d
```

### 4. Install Dependencies

#### Backend
```bash
cd backend
go mod tidy
```

#### Frontend
```bash
cd frontend
npm install
```

### 5. Create Upload Directory

```bash
mkdir backend/uploads
```

### 6. Run the Application

#### Start Backend (from `backend/` directory)
```bash
go run main.go
```

#### Start Frontend (from `frontend/` directory)
```bash
npm run dev
```

### 7. Access the Application

- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080

## 🌐 Network Access

To access the application from other devices on your network:

### 1. Update Frontend Configuration
```bash
# In frontend/.env
VITE_API_BASE_URL=http://YOUR_IP_ADDRESS:8080
```

### 2. Start Frontend with Network Access
```bash
npm run dev -- --host 0.0.0.0
```

### 3. Configure Windows Firewall (Windows users)
```powershell
# Run as Administrator
netsh advfirewall firewall add rule name="Go Backend" dir=in action=allow protocol=TCP localport=8080
netsh advfirewall firewall add rule name="Vite Frontend" dir=in action=allow protocol=TCP localport=5173
```

### 4. Access from Other Devices
- **Frontend**: http://YOUR_IP_ADDRESS:5173
- **Backend API**: http://YOUR_IP_ADDRESS:8080

## 📝 API Endpoints

### Authentication
- `POST /auth/signup` - Register a new user
- `POST /auth/login` - Login user

### Protected Routes (require JWT token)
- `GET /user/profile` - Get user profile
- `GET /applications` - Get all applications
- `GET /applications/:id` - Get specific application
- `POST /applications` - Create new application
- `PUT /applications/:id` - Update application
- `DELETE /applications/:id` - Delete application
- `GET /uploads/*filepath` - Serve uploaded files

## 🔧 Development

### Project Structure
```
internship_hub_GO/
├── backend/               # Go backend
│   ├── config/           # Database and environment config
│   ├── controllers/      # Route handlers
│   ├── middleware/       # JWT authentication middleware
│   ├── models/          # Database models
│   ├── uploads/         # File upload directory
│   └── main.go          # Entry point
├── frontend/             # SvelteKit frontend
│   ├── src/
│   │   ├── lib/         # Shared components and services
│   │   └── routes/      # Page routes
│   └── package.json
└── docker-compose.yml    # Database container
```

### Running Tests
```bash
# Backend tests
cd backend
go test ./...

# Frontend tests  
cd frontend
npm test
```

### Building for Production
```bash
# Backend
cd backend
go build -o internship-hub main.go

# Frontend
cd frontend
npm run build
```

## 🔒 Security Features

- JWT-based authentication
- CORS protection
- Environment variable configuration
- Secure file upload handling
- Input validation and sanitization

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Ensure PostgreSQL container is running: `docker-compose ps`
   - Check database credentials in `.env` files

2. **CORS Errors**
   - Verify `VITE_API_BASE_URL` in frontend `.env`
   - Check CORS configuration in `backend/main.go`

3. **File Upload Errors**
   - Ensure `backend/uploads/` directory exists
   - Check file permissions

4. **Network Access Issues**
   - Configure firewall rules
   - Use correct IP address in environment variables

## 📄 License

This project is licensed under the MIT License.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## 📞 Support

If you encounter any issues, please create an issue on GitHub or contact the maintainer.
