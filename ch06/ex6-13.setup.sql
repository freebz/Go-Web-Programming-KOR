-- 예제 6-13. 두 개의 관계된 테이블 설정

drop table if exists posts cascade;
drop table if exists comments;

create table posts (
  id     serial primary key,
  content text,
  author  varchar(255)
);

create table comments (
  id     serial primary key,
  content text,
  author  varchar(255),
  post_id integer references posts(id)
);
