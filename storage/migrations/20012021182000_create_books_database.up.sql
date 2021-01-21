

-- Database books


SELECT 'CREATE DATABASE books_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'books_db')\gexec
