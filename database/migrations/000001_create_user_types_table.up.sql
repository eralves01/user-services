CREATE TABLE user_types (
    id INTEGER UNIQUE NOT NULL,
    value VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO user_types (id, value) VALUES (1, 'client'), (2,'merchant');
