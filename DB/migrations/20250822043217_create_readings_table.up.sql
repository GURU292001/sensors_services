-- Up Migration: Create the readings_table
CREATE TABLE IF NOT EXISTS readings_table (
    sensor_value FLOAT NOT NULL,               
    sensor_type  VARCHAR(50) NOT NULL,        
    id1 CHAR(10) NOT NULL,                 
    id2 INT AUTO_INCREMENT NOT NULL,                        
    time_stamp TIMESTAMP NOT NULL,
    PRIMARY KEY (id2)   -- Auto_increment must be primary key or unique
);
