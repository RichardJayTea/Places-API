
-- +migrate Up
CREATE TABLE teago.`user` (
  uid INT auto_increment NOT NULL,
  username VARCHAR(150) NOT NULL,
  email VARCHAR(150) NOT NULL,
  password VARCHAR(255) NOT NULL,
  first_name VARCHAR(150) NOT NULL,
  last_name VARCHAR(150) NOT NULL,
  is_active TINYINT(4) NOT NULL,
  image_path VARCHAR(150) NOT NULL DEFAULT '',
  create_date TIMESTAMP NOT NULL,
  CONSTRAINT user_PK PRIMARY KEY (uid)
);
-- +migrate Down
DROP TABLE teago.`user`;
