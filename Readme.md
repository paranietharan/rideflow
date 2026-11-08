# Ride flow

- A Go-based ride-hailing backend featuring geospatial driver discovery, concurrent ride matching, trip lifecycle management, and MongoDB-powered analytics.

## Configuration Details

### Server config
| Variable | Description | Example Value | Default Value |
|---|---|---|---|
| `SERVER_PORT` | Port on which the HTTP server listens. | `8000` | `8000` |
| `APP_ENV` | Application environment (e.g., `development`, `staging`, `production`). | `development` | `development` |
| `MONGO_URI` | MongoDB connection string used to connect to the database. | `mongodb://localhost:27017` | `mongodb://localhost:27017` |
| `MONGO_DB_NAME` | Name of the MongoDB database used by the application. | `rideflow` | `rideflow` |
| `JWT_SECRET` | Secret key used to sign and verify JSON Web Tokens (JWT). | `your-secure-random-secret` | `rideflow-dev-secret-key` |