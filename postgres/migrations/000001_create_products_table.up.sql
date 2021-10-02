CREATE TABLE IF NOT EXISTS products (
                                           id INTEGER PRIMARY KEY NOT NULL,
                                           vendor_code VARCHAR(255) NOT NULL,
                                           name VARCHAR(255) NOT NULL,
                                           category INTEGER NOT NULL,
                                           type INTEGER NOT NULL,
                                           description VARCHAR(255) NOT NULL,
                                           stock_balance INTEGER NOT NULL,
                                           specifications jsonb,
                                           created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
                                           updated_at TIMESTAMPTZ,
                                           FOREIGN KEY (category) REFERENCES categories(id),
                                           FOREIGN KEY (type) REFERENCES types(id)


);