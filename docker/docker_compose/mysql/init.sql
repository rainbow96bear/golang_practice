-- "Books" 데이터베이스를 생성합니다.
CREATE DATABASE IF NOT EXISTS Books;

-- "Books" 데이터베이스를 사용합니다.
USE Books;

-- "books" 테이블을 생성합니다.
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    kind VARCHAR(50),
    title VARCHAR(50),
    author VARCHAR(50)
);