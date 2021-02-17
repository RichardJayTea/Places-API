-- +migrate Up
CREATE TABLE teago.`checkin` (
                               uid INT auto_increment NOT NULL,
                               user_id INT NOT NULL,
                               place_id INT NOT NULL,
                               create_date TIMESTAMP NOT NULL,
                               CONSTRAINT checkin_PK PRIMARY KEY (uid),
                               CONSTRAINT checkin_user_id_FK FOREIGN KEY (user_id) REFERENCES teago.user(uid),
                               CONSTRAINT checkin_place_id_FK FOREIGN KEY (place_id) REFERENCES teago.place(uid)
);
-- +migrate Down
DROP TABLE teago.`checkin`;
