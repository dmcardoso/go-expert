CREATE TABLE categories(id UUID, name VARCHAR, description VARCHAR);
CREATE TABLE courses(id UUID, name VARCHAR, description VARCHAR, category_id UUID);
