CREATE TABLE todos
(
    id          BIGINT       NOT NULL auto_increment,
    task       VARCHAR(100) NOT NULL,
    description TEXT         NULL,
    completed BOOLEAN DEFAULT false,
    created_at  TIMESTAMP    NOT NULL DEFAULT current_timestamp,
    updated_at  TIMESTAMP    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    deleted_at  TIMESTAMP    NULL,
    PRIMARY KEY (id)
) engine = innodb;