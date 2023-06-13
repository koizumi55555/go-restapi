CREATE ROLE user1 LOGIN SUPERUSER PASSWORD 'password';
CREATE DATABASE koizumi;
GRANT all privileges ON DATABASE koizumi TO user1;