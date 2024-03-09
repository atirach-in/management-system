
-- create
CREATE TABLE Books (
  Id INTEGER AUTO_INCREMENT  PRIMARY KEY,
  title TEXT NOT NULL,
  author_id INTEGER NOT NULL,
  publisher_id INTEGER NOT NULL,
  publication_year INTEGER NOT NULL,
  price decimal(10,2) NOT NULL
);

CREATE TABLE Authors(
 Id INTEGER AUTO_INCREMENT  PRIMARY KEY,
 author_name TEXT NOT NULL
);

CREATE TABLE Publishers(
 Id INTEGER AUTO_INCREMENT  PRIMARY KEY,
 publishers_name TEXT NOT NULL
);

-- insert Authors
INSERT INTO Authors(author_name) VALUES ('Clark');
INSERT INTO Authors(author_name) VALUES ('Dave');
INSERT INTO Authors(author_name) VALUES ('Ava');

-- insert Publishers
INSERT INTO Publishers(publishers_name) VALUES ('John');
INSERT INTO Publishers(publishers_name) VALUES ('Danial');
INSERT INTO Publishers(publishers_name) VALUES ('Ken');

INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('JAVA', 1, 3, 2012, 189);
INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('PHP', 2, 2, 2013, 175);
INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('Nodejs', 3, 1, 2014, 299);
INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('Flutter', 1, 3, 2015, 235);
INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('Nosql', 2, 2, 2016, 499);
INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('Firebase', 3, 1, 2017, 99);

INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('Socket', 3, 1, 2017, 89);
INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('K8s', 3, 2, 2023, 899);
INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('Docker', 3, 3, 2023, 399);
INSERT INTO Books(title, author_id, publisher_id, publication_year, price) VALUES ('React', 3, 1, 2023, 589);

-- fetch 
SELECT * FROM Authors;
SELECT * FROM Publishers;
SELECT * FROM Books;

SELECT p.publishers_name, AVG(b.price) AS average_price
FROM Books AS b
JOIN Publishers AS p ON b.publisher_id = p.Id
GROUP BY b.publisher_id
ORDER BY AVG(b.price) DESC
LIMIT 3;

SELECT b.title, a.author_name
FROM Books AS b
JOIN Authors AS a ON b.author_id = a.Id
WHERE b.author_id IN (
    SELECT author_id
    FROM Books
    GROUP BY author_id
    HAVING COUNT(*) > 3
);
