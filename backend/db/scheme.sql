CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    password TEXT NOT NULL, 
    email TEXT NOT NULL UNIQUE,
    role TEXT NOT NULL CHECK(role IN ('customer', 'roaster', 'admin')),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE coffee (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL, -- Coffee bean name
    description TEXT, -- Optional description of coffee
    roast_type TEXT NOT NULL, -- e.g., 'light', 'medium', 'dark'
    stock INTEGER NOT NULL DEFAULT 0, -- Amount of coffee in stock (grams, pounds, etc.)
    price DECIMAL(10, 2) NOT NULL, -- Price per unit (e.g., per 100g)
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL, -- Customer placing the order
    total_amount DECIMAL(10, 2) NOT NULL, -- Total order price
    status TEXT NOT NULL CHECK(status IN ('pending', 'in_progress', 'shipped', 'completed', 'cancelled')), -- Order status
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE order_items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    order_id INTEGER NOT NULL, -- Associated order
    coffee_id INTEGER NOT NULL, -- Coffee product in the order
    quantity INTEGER NOT NULL, -- Quantity ordered
    unit_price DECIMAL(10, 2) NOT NULL, -- Price per unit at the time of order
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (coffee_id) REFERENCES coffee(id)
);

-- Add this to your database schema
CREATE TABLE sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    expires_at DATETIME NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Add an index for faster token lookups
CREATE INDEX idx_sessions_token ON sessions(token);

-- Add an index for cleanup of expired sessions
CREATE INDEX idx_sessions_expires_at ON sessions(expires_at);