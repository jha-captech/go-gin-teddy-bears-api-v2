-- Drop tables if they exist
DROP TABLE IF EXISTS picnic_participants;

DROP TABLE IF EXISTS picnics;

DROP TABLE IF EXISTS picnic_locations;

DROP TABLE IF EXISTS teddy_bears;

-- Create picnics table
CREATE TABLE
    picnics (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        picnic_name TEXT UNIQUE NOT NULL,
        location_id INTEGER,
        start_time DATETIME NOT NULL,
        has_music BOOLEAN NOT NULL DEFAULT 1,
        has_food BOOLEAN NOT NULL DEFAULT 1
    );

-- Create picnic_locations table
CREATE TABLE
    picnic_locations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        location_name TEXT UNIQUE NOT NULL,
        capacity INTEGER NOT NULL DEFAULT 25,
        municipality TEXT NOT NULL
    );

-- Create teddy_bears table
CREATE TABLE
    teddy_bears (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT UNIQUE NOT NULL,
        primary_color TEXT NOT NULL,
        accent_color TEXT,
        is_dressed BOOLEAN NOT NULL DEFAULT 1,
        owner_name TEXT NOT NULL,
        characteristic TEXT
    );

-- Create picnic_participants table
CREATE TABLE
    picnic_participants (
        picnic_id INTEGER NOT NULL,
        teddy_bear_id INTEGER NOT NULL,
        PRIMARY KEY (picnic_id, teddy_bear_id),
        FOREIGN KEY (picnic_id) REFERENCES picnics (id),
        FOREIGN KEY (teddy_bear_id) REFERENCES teddy_bears (id)
    );

-- Insert data into picnics table
INSERT INTO
    picnics (
        picnic_name,
        location_id,
        start_time,
        has_music,
        has_food
    )
VALUES
    (
        'Picnic At Oakwood',
        1,
        '2023-03-04 14:00:00',
        1,
        1
    ),
    (
        '2nd Picnic At Oakwood',
        1,
        '2023-03-11 15:00:00',
        1,
        1
    ),
    (
        '100 Acre Festival',
        2,
        '2023-06-21 14:30:00',
        1,
        1
    ),
    (
        'Mid-Summer Picnic',
        3,
        '2023-07-29 15:00:00',
        1,
        1
    );

-- Insert data into picnic_locations table
INSERT INTO
    picnic_locations (location_name, capacity, municipality)
VALUES
    ('Big Wood', 25, 'Oakwood'),
    ('100 Acre Wood', 30, 'East Sussex'),
    ('The Commons', 20, 'Coppell');

-- Insert data into picnic_participants table
INSERT INTO
    picnic_participants (picnic_id, teddy_bear_id)
VALUES
    (1, 1),
    (1, 2),
    (1, 4),
    (1, 6),
    (2, 1),
    (2, 3),
    (2, 4),
    (2, 5),
    (2, 6);

-- Insert data into teddy_bears table
INSERT INTO
    teddy_bears (
        name,
        primary_color,
        accent_color,
        is_dressed,
        owner_name,
        characteristic
    )
VALUES
    (
        'Teddy',
        'Brown',
        NULL,
        1,
        'Little Billy',
        'The one true Teddy'
    ),
    (
        'Suzie',
        'Light Brown',
        'Black',
        1,
        'Janey',
        'Cuddly'
    ),
    (
        'TouTou',
        'Pink',
        'White',
        1,
        'Sarah',
        'Nylon skin'
    ),
    ('Nounours', 'Brown', 'Red', 0, 'Clair', 'Fluffy'),
    ('Bear', 'Light Blue', 'White', 0, 'Xavier', NULL),
    (
        'Winnie the Pooh',
        'Yellow',
        NULL,
        1,
        'Christopher Robin',
        'Red Shirt'
    );