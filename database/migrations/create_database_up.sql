
CREATE TABLE IF NOT EXISTS users (
    id_user SERIAL CONSTRAINT pk_id_user PRIMARY KEY,
    user_name VARCHAR(128) NOT NULL,
    email VARCHAR(128) NOT NULL,
    password_key VARCHAR(256) NOT NULL
);
CREATE TABLE IF NOT EXISTS route_history (
    id_history INTEGER NOT NULL CONSTRAINT pk_id_history PRIMARY KEY,
    id_user INTEGER NOT NULL,
    latitude_from REAL NOT NULL,
    longitude_from REAL NOT NULL,
    latitude_where REAL NOT NULL,
    longitude_where REAL NOT NULL,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
    
);
