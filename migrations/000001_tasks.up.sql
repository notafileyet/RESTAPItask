CREATE TABLE tasks
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    status     VARCHAR(100),
    is_done    BOOLEAN               DEFAULT FALSE,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP             DEFAULT NULL
);