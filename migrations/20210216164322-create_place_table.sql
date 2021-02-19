-- +migrate Up
CREATE TABLE teago.`place` (
                              uid INT auto_increment NOT NULL,
                              external_id INT NOT NULL,
                              name VARCHAR(150) NOT NULL,
                              description VARCHAR(255) NOT NULL,
                              image_path VARCHAR(150) NOT NULL DEFAULT '',
                              is_active TINYINT(4) NOT NULL,
                              create_date TIMESTAMP NOT NULL,
                              CONSTRAINT place_PK PRIMARY KEY (uid)
);
-- +migrate Down
DROP TABLE teago.`place`;
