# RideFlow — Implementation TODO

**Project:** Ride-hailing backend API  
**Stack:** Go, Chi, MongoDB, Redis, Docker Compose  
**Architecture:** Modular monolith  
**Goal:** Build a working ride-hailing backend while learning MongoDB from fundamentals to advanced querying and optimization.

## Phase 1 — Project Setup

- [x] Create the GitHub repository `rideflow`.
- [x] Initialize the Go module.
- [x] Install Chi and the official MongoDB Go Driver.
- [x] Set up MongoDB using Docker Compose.
- [x] Configure environment variables for the MongoDB URI, database name, server port, and secrets.
- [x] Create a shared MongoDB client with connection pooling and graceful shutdown.
- [x] Add health and readiness endpoints.
- [x] Establish a clean package structure.

**MongoDB practice:** Connect, select a database, create collections, and perform basic CRUD operations using `mongosh`.

## Phase 2 — Data Modeling

- [ ] Define passenger documents.
- [ ] Define driver documents.
- [ ] Define ride request documents.
- [ ] Define trip history documents.
- [ ] Define fare breakdown documents.
- [ ] Choose which fields to embed and which entities to reference.
- [ ] Use BSON dates for timestamps.
- [ ] Model geographic locations using GeoJSON.
- [ ] Add collection schema validation.
- [ ] Document the reasoning behind the data model.

**MongoDB practice:** Embedding versus referencing, BSON types, schema validation, and document size limits.

## Phase 3 — Authentication and Users

- [ ] Implement passenger registration and login.
- [ ] Implement driver registration and login.
- [ ] Hash passwords securely.
- [ ] Add JWT authentication middleware.
- [ ] Implement the current-user endpoint.
- [ ] Add role-based authorization.
- [ ] Prevent duplicate accounts using unique indexes.
- [ ] Validate requests and handle duplicate-key errors.

**MongoDB practice:** CRUD, unique indexes, projections, and validation.

## Phase 4 — Driver Management

- [ ] Allow drivers to switch between available and unavailable.
- [ ] Add an endpoint to update the driver's current location.
- [ ] Store location as a GeoJSON Point.
- [ ] Create a `2dsphere` index.
- [ ] Implement nearby-driver discovery.
- [ ] Filter nearby drivers by availability and supported ride category.
- [ ] Add distance limits and result limits.
- [ ] Test invalid coordinates and boundary cases.

**MongoDB practice:** Geospatial queries, compound indexes, and query execution plans.

## Phase 5 — Ride Requests

- [ ] Allow passengers to request a ride.
- [ ] Validate pickup, destination, and ride category.
- [ ] Calculate an estimated fare.
- [ ] Create a ride document with a defined initial state.
- [ ] Implement ride-request retrieval.
- [ ] Allow passengers to cancel eligible requests.
- [ ] Prevent invalid state transitions.
- [ ] Add idempotency keys to prevent duplicate ride requests.

**MongoDB practice:** Conditional updates, atomic writes, nested documents, and unique indexes.

## Phase 6 — Driver Matching

- [ ] Find nearby eligible drivers.
- [ ] Rank candidates using distance and availability.
- [ ] Offer the ride to a driver.
- [ ] Implement driver acceptance and rejection.
- [ ] Handle offer expiration and retry with another driver.
- [ ] Prevent multiple passengers from claiming the same driver.
- [ ] Handle simultaneous acceptance attempts safely.
- [ ] Add integration tests for concurrent requests.

**MongoDB practice:** Atomic conditional updates, concurrency control, TTL indexes for temporary records where appropriate, and transaction trade-offs.

## Phase 7 — Trip Lifecycle

- [ ] Implement driver arrival.
- [ ] Implement trip start.
- [ ] Implement trip completion.
- [ ] Implement eligible cancellation flows.
- [ ] Validate every state transition on the server.
- [ ] Record trip status history.
- [ ] Calculate the final fare.
- [ ] Make trip completion idempotent.
- [ ] Save pickup and destination snapshots for historical accuracy.

**MongoDB practice:** Atomic updates, embedded versus separate history records, transactions where necessary, and schema validation.

## Phase 8 — Trip History and Ratings

- [ ] List a passenger's previous rides.
- [ ] List a driver's completed trips.
- [ ] Add date and status filters.
- [ ] Implement pagination.
- [ ] Add passenger and driver ratings.
- [ ] Prevent duplicate ratings for the same trip and reviewer.
- [ ] Create appropriate compound indexes.

**MongoDB practice:** Filtering, sorting, compound indexes, pagination, and unique constraints.

## Phase 9 — Aggregation and Analytics

- [ ] Calculate daily ride counts.
- [ ] Calculate completed and cancelled ride percentages.
- [ ] Calculate average fare by ride category.
- [ ] Calculate driver earnings over a date range.
- [ ] Identify popular pickup areas.
- [ ] Compare `$lookup` with an alternative document design.
- [ ] Build reports using `$match`, `$group`, `$sort`, `$project`, and `$unwind`.
- [ ] Use `$facet` for multi-part analytics where useful.

**MongoDB practice:** Aggregation pipelines, accumulators, joins, grouping, and date expressions.

## Phase 10 — Performance and Reliability

- [ ] Run `explain("executionStats")` on important queries.
- [ ] Identify collection scans and inefficient queries.
- [ ] Add indexes based on measured access patterns.
- [ ] Compare query performance before and after indexing.
- [ ] Test realistic concurrent ride requests.
- [ ] Configure request and database-operation timeouts.
- [ ] Handle database errors without exposing internal details.
- [ ] Add structured logging.
- [ ] Add graceful shutdown and resource cleanup.
- [ ] Back up and restore a test database.

**MongoDB practice:** Index selectivity, query plans, connection pooling, durability, and performance measurement.

## Phase 11 — Testing and Documentation

- [ ] Write unit tests for fare calculation and state transitions.
- [ ] Write integration tests against a real MongoDB instance.
- [ ] Test invalid inputs and missing documents.
- [ ] Test duplicate requests and concurrent driver assignment.
- [ ] Document all API endpoints with OpenAPI.
- [x] Add example environment configuration.
- [ ] Write a README with setup and usage instructions.
- [ ] Document the data model and indexes.
- [ ] Record performance findings and design trade-offs.
- [ ] Ensure no credentials or secrets are committed.

## Phase 12 — Optional Advanced Features

- [ ] Add WebSocket updates for ride status.
- [ ] Explore MongoDB change streams.
- [ ] Add historical driver-location telemetry using a time-series collection if justified.
- [ ] Add a driver earnings dashboard endpoint.
- [ ] Experiment with replica sets and transactions.
- [ ] Document a possible scaling strategy without implementing microservices prematurely.

## Completion Criteria

- [ ] A passenger can register, log in, request a ride, and view trip history.
- [ ] A driver can register, update location, become available, accept a ride, and complete a trip.
- [ ] Driver discovery uses MongoDB geospatial queries.
- [ ] Concurrent requests cannot assign the same driver to conflicting rides.
- [ ] Fare calculations and trip state transitions are validated.
- [ ] Aggregation endpoints return correct analytics.
- [ ] Important queries have appropriate indexes and measured query plans.
- [ ] Tests, API documentation, and setup instructions are complete.
