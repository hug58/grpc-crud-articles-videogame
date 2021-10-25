

CREATE TABLE IF NOT EXISTS articles (
    id serial NOT NULL,
    name VARCHAR(150) NOT NULL,
    price INT default 0, 
    description text NOT NULL,
    user_id INT NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    CONSTRAINT pk_notes PRIMARY KEY(id)
);

