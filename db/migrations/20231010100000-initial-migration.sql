-- +migrate Up
CREATE TABLE users (
  id SERIAL NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  full_name TEXT NOT NULL,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE users;
