-- Create location table
CREATE TABLE location (
    location_id VARCHAR(255) PRIMARY KEY,
    lat DECIMAL(9, 6),
    lon DECIMAL(9, 6),
    alt DECIMAL(10, 8)
);

-- Create camera table
CREATE TABLE camera (
    camera_id VARCHAR(255) PRIMARY KEY,
    type VARCHAR(255),
    description_cam VARCHAR(255)
);

-- Create model table
CREATE TABLE model (
    model_id VARCHAR(255) PRIMARY KEY,
    description_mdl VARCHAR(255)
);

-- Create message table with foreign keys
CREATE TABLE message (
    message_id VARCHAR(255) PRIMARY KEY,
    timestamp TIMESTAMP,
    location_id VARCHAR(255),
    model_id VARCHAR(255),
    camera_id VARCHAR(255),
    number_of_objects INT,
    number_of_events INT,
    status BOOLEAN DEFAULT NULL,
    image_URL VARCHAR(255),
    video_URL VARCHAR(255),
    FOREIGN KEY (location_id) REFERENCES location (location_id),
    FOREIGN KEY (model_id) REFERENCES model (model_id),
    FOREIGN KEY (camera_id) REFERENCES camera (camera_id)
);

-- Create object table with foreign key
CREATE TABLE object (
    message_id VARCHAR(255),
    object_id VARCHAR(255),
    object_type VARCHAR(255),
    gender VARCHAR(255),
    age VARCHAR(255),
    vehicle_type VARCHAR(255),
    vehicle_brand VARCHAR(255),
    vehicle_color VARCHAR(255),
    vehicle_license VARCHAR(255),
    bbox_topleftx INT,
    bbox_toplefty INT,
    bbox_bottomrightx INT,
    bbox_bottomrighty INT,
    PRIMARY KEY (message_id, object_id),
    FOREIGN KEY (message_id) REFERENCES message (message_id)
);

-- Create event table with foreign keys and unique constraint
CREATE TABLE event (
    message_id VARCHAR(255),
    object_id VARCHAR(255),
    event_type VARCHAR(255),
    action VARCHAR(255),
    FOREIGN KEY (message_id, object_id) REFERENCES object (message_id, object_id),
    CONSTRAINT uc_event_message UNIQUE (object_id, message_id)
);

-- Set all constraints to be deferred
SET CONSTRAINTS ALL DEFERRED;
