CREATE TABLE users(
    id uuid primary key not null default gen_random_uuid(),
    email varchar not null unique,
    encrypted_password varchar not null
);