CREATE TABLE IF NOT EXISTS urldetail(
    id SERIAL PRIMARY KEY,
    url VARCHAR (500) UNIQUE NOT NULL,
    path VARCHAR (500) NOT NULL,
    is_custom BOOLEAN DEFAULT false,
    created_at TIMESTAMP
);