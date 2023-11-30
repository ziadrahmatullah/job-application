CREATE DATABASE job_db;

CREATE TABLE IF NOT EXISTS users(
    id bigserial,
    name varchar NOT NULL,
    current_job varchar NOT NULL,
    age int NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS jobs(
    id bigserial,
    name varchar NOT NUll,
    company varchar NOT NULL,
    quota int,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS job_applies(
    id bigserial,
    user_id bigint NOT NULL,
    job_id bigint NOT NULL,
    aplied_at timestamp NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (job_id) REFERENCES jobs(id),
);