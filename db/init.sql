CREATE USER smart_scale_user WITH PASSWORD '******';

CREATE DATABASE smart_scale;

\c smart_scale;

CREATE TABLE weights(
   id BIGSERIAL PRIMARY KEY,
   empid INT NOT NULL,
   weight DECIMAL NOT NULL,
   recorded_at timestamp
);

CREATE USER smart_scale_user WITH PASSWORD '******';

REVOKE CONNECT ON DATABASE smart_scale FROM PUBLIC;

GRANT CONNECT ON DATABASE smart_scale TO smart_scale_user;

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO smart_scale_user;