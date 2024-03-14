-- Create database
CREATE DATABASE test_snippetbox WITH ENCODING='UTF8';

-- Create user
CREATE USER test_web WITH PASSWORD 'pass';

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE test_snippetbox TO test_web;
GRANT ALL ON SCHEMA public TO test_web;