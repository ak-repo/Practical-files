

-- create table posts (
--  id serial primary key,
--  content text,
--  author varchar(255)
-- );


-- Drop tables if they exist (order matters because of foreign keys)
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS posts CASCADE;

-- Create posts table
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    content TEXT,
    author VARCHAR(255)
);

-- Create comments table with foreign key
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content TEXT,
    author VARCHAR(255),
    post_id INTEGER REFERENCES posts(id) ON DELETE CASCADE
);



