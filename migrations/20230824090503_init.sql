CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE address_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    user_id INT,
    recipient_name VARCHAR(255),
    -- e.g., the person or company's name
    address_line1 VARCHAR(255) NOT NULL,
    -- e.g., "123 Main St"
    address_line2 VARCHAR(255),
    -- e.g., "Suite 4B" or "Building A"
    address_line3 VARCHAR(255),
    -- Can be used for additional details
    city VARCHAR(100) NOT NULL,
    state_province_region VARCHAR(100),
    -- This can be a state, province, region, etc.
    postal_code VARCHAR(20),
    -- This can be zip code, postal code, etc.
    country VARCHAR(100) NOT NULL,
    -- Ideally a standardized country name or code
    phone VARCHAR(50),
    -- Phone number associated with this address
    address_type_id INT NOT NULL,
    -- Any other info or special instructions
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (address_type_id) REFERENCES address_types(id)
);