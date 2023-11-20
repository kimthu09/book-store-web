# Book store management

## I. How to run project

**Book Store Management:**

- Front-end: `localhost:3000`
- Back-end: `localhost:8080`
- MySQL: `localhost:39062`

### 1. Move to docker-compose folder

```bash
cd docker-compose
```

### 2. Run docker-compose

```bash
docker-compose up
```

### 3. Create table (if db empty)

#### a. Connect to database

Host: `127.0.0.1:39062`

User: `root`

Password: `123456`

DB Name: `bookstoremanagement`

#### b. Execute script

Run script on `/book-store-management-backend/sql/create-table.sql`
