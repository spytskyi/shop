CREATE TABLE IF NOT EXISTS types (
                                           id UUID PRIMARY KEY NOT NULL,
                                           name VARCHAR(255) NOT NULL,
                                           category UUID NOT NULL,
                                           FOREIGN KEY (category) REFERENCES categories(id)
);