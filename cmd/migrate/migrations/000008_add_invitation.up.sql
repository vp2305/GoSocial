CREATE TABLE IF NOT EXISTS user_invitations (
    token BYTEA PRIMARY KEY,
    user_id BIGINT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);