CREATE TABLE IF NOT EXISTS types (
                                           id SERIAL PRIMARY KEY,
                                           name VARCHAR(255) NOT NULL,
                                           category_id INT NOT NULL,
                                           FOREIGN KEY (category_id) REFERENCES categories(id)
);