
	
-- psql -U username -f setup.sql -d dbname -W

CREATE TABLE todos(
    id SERIAL PRIMARY KEY NOT NULL,
    content  TEXT NOT NULL,
    done  BOOLEAN 
)

