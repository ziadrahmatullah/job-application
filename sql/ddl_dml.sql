CREATE DATABASE job_db;

CREATE TABLE IF NOT EXISTS users(
    id bigserial,
    name varchar NOT NULL,
    current_job varchar NOT NULL,
    age int NOT NULL CHECK (age > 0),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS jobs(
    id bigserial,
    name varchar NOT NUll,
    company varchar NOT NULL,
    quota int NOT NULL CHECK (quota > 0),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS job_applies(
    id bigserial,
    user_id bigint NOT NULL,
    job_id bigint NOT NULL,
    aplied_at timestamp NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (job_id) REFERENCES jobs(id)
);

INSERT INTO users (name, current_job, age, created_at, updated_at)
VALUES
('Alice', 'Job Researcher', 18, NOW(), NOW()),
('Bob', 'Job Researcher', 18, NOW(), NOW());

INSERT INTO jobs(name, company, quota, created_at, updated_at)
VALUES
('Backend Developer', 'Shopee', 2, NOW(), NOW()),
('Frontend Developer', 'Shopee', 2, NOW(), NOW());