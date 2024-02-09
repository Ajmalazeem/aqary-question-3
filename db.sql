CREATE DATABASE CLASS;

CREATE TABLE Seat (
    id SERIAL PRIMARY KEY,
    student VARCHAR
);

-- if number of students is odd

INSERT INTO Seat (student) VALUES ('Abbot');
INSERT INTO Seat (student) VALUES ('Doris');
INSERT INTO Seat (student) VALUES ('Emerson');
INSERT INTO Seat (student) VALUES ('Green');
INSERT INTO Seat (student) VALUES ('Jeames');

-- -- if number of students is even

INSERT INTO Seat (student) VALUES ('Abbot');
INSERT INTO Seat (student) VALUES ('Doris');
INSERT INTO Seat (student) VALUES ('Emerson');
INSERT INTO Seat (student) VALUES ('Green');


-- SELECT id, student FROM seat 