CREATE TABLE users(
    userid bigserial primary key ,
    public_address varchar(100) UNIQUE
);
insert into users(public_address)
values
('0x8a8cB99FBE932c2fBED13B4982e4fE1BE364d58C');

SELECT * FROM users WHERE public_address = '0x8a8cB99FBE932c2fBED13B4982e4fE1BE364d58C';

CREATE TABLE balances(
    balanceid bigserial primary key ,
    userid bigserial,
    amount bigint,
    coinid bigint,
    CONSTRAINT fk_user FOREIGN KEY (userid) REFERENCES users(userid)
);
