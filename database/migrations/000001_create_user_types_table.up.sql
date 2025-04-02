CREATE TABLE user_types (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO user_types (type) VALUES ('client'), ('merchant');
