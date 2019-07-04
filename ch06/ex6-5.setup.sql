-- 예제 6-5. 데이터베이스 생성을 위한 스크립트

create table posts (
  id     serial primary key,
  content text,
  author  varchar(255)
);
