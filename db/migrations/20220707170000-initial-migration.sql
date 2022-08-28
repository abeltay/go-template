-- +migrate Up
CREATE TABLE users (
  id SERIAL NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE,
  full_name TEXT,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE users;
