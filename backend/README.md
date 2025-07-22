# TwinWorld Backend

A Go-based REST API backend for the TwinWorld platform, connecting large companies with small businesses to transform textile waste into innovation.

## Features

- **Company Management**: Manage large fashion companies and their waste materials
- **Business Management**: Support small businesses and entrepreneurs
- **Material Tracking**: Track waste materials and their availability
- **Partnership Management**: Connect companies with businesses
- **AI Integration**: Endpoints for AI-powered redesign suggestions
- **Impact Metrics**: Track environmental and economic impact

## Tech Stack

- **Framework**: Gin (Go web framework)
- **Database**: SQLite (with GORM ORM)
- **CORS**: Cross-origin resource sharing enabled
- **Validation**: Request validation with Gin binding

## API Endpoints

### Health Check
- `GET /health` - Server health status

### Companies
- `GET /api/v1/companies` - List all companies
- `POST /api/v1/companies` - Create a new company
- `GET /api/v1/companies/:id` - Get company by ID
- `PUT /api/v1/companies/:id` - Update company
- `DELETE /api/v1/companies/:id` - Delete company

### Businesses
- `GET /api/v1/businesses` - List all businesses
- `POST /api/v1/businesses` - Create a new business
- `GET /api/v1/businesses/:id` - Get business by ID
- `PUT /api/v1/businesses/:id` - Update business
- `DELETE /api/v1/businesses/:id` - Delete business

### Materials
- `GET /api/v1/materials` - List all materials
- `POST /api/v1/materials` - Create a new material
- `GET /api/v1/materials/:id` - Get material by ID
- `PUT /api/v1/materials/:id` - Update material
- `DELETE /api/v1/materials/:id` - Delete material

### AI Redesign
- `POST /api/v1/ai/redesign` - Submit for AI redesign
- `GET /api/v1/ai/suggestions` - Get AI suggestions

### Partnerships
- `GET /api/v1/partnerships` - List all partnerships
- `POST /api/v1/partnerships` - Create a new partnership
- `GET /api/v1/partnerships/:id` - Get partnership by ID
- `PUT /api/v1/partnerships/:id` - Update partnership
- `DELETE /api/v1/partnerships/:id` - Delete partnership

### Impact Metrics
- `GET /api/v1/impact` - Get all impact metrics
- `GET /api/v1/impact/waste-reduction` - Get waste reduction stats
- `GET /api/v1/impact/economic-value` - Get economic value created
- `GET /api/v1/impact/businesses-empowered` - Get businesses empowered count

## Setup

1. **Install Go** (if not already installed):
   ```bash
   brew install go
   ```

2. **Navigate to backend directory**:
   ```bash
   cd backend
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Run the server**:
   ```bash
   go run cmd/server/main.go
   ```

The server will start on `http://localhost:8081` by default.

## Environment Variables

- `PORT` - Server port (default: 8081)

## Development

### Project Structure
```
backend/
├── cmd/
│   └── server/
│       └── main.go
├── handlers/
├── models/
│   ├── company.go
│   ├── business.go
│   ├── material.go
│   └── partnership.go
├── database/
├── middleware/
├── config/
├── go.mod
├── go.sum
└── README.md
```

### Adding New Endpoints

1. Create handler functions in `handlers/` directory
2. Add routes in `cmd/server/main.go`
3. Create models in `models/` directory if needed

## Testing

Test the API endpoints using curl or any API client:

```bash
# Health check
curl http://localhost:8081/health

# Get companies
curl http://localhost:8081/api/v1/companies

# Get impact metrics
curl http://localhost:8081/api/v1/impact
```

## CORS Configuration

The backend is configured to allow requests from:
- `http://localhost:8080` (Vite dev server)
- `http://localhost:3000` (React dev server)

## Database

Currently using SQLite for development. The database file will be created automatically when the server starts.

## Future Enhancements

- [ ] PostgreSQL integration
- [ ] Authentication & Authorization
- [ ] File upload for AI redesign
- [ ] Real-time notifications
- [ ] Advanced analytics
- [ ] Rate limiting
- [ ] API documentation with Swagger 