CREATE DATABASE job_db;

CREATE TABLE IF NOT EXISTS users(
    id bigserial,
    name varchar NOT NULL,
    current_job varchar NOT NULL,
    age int NOT NULL CHECK (age > 0),
    email varchar NOT NULL,
    password varchar NOT NULL,
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
    expired_at timestamp NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS apply_jobs(
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

INSERT INTO users (name, current_job, age,email, password, created_at, updated_at)
VALUES
('Alice', 'Job Researcher', 18,'alice@gmail.com' 'no hash', NOW(), NOW()),
('Bob', 'Job Researcher', 18,'bob@gmail.com', 'no hash', NOW(), NOW());

INSERT INTO jobs(name, company, quota, expired_at, created_at, updated_at)
VALUES
('Backend Developer', 'Shopee', 2,'2024-01-01', NOW(), NOW()),
('Frontend Developer', 'Shopee', 2,'2024-01-01' , NOW(), NOW());