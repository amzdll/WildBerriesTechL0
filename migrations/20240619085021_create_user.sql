-- +goose Up

CREATE USER developer WITH PASSWORD 'developer';
GRANT ALL PRIVILEGES ON DATABASE wb TO developer;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE "delivery", "items", "order", "payment" TO developer;