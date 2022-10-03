SELECT users.public_address, balances.amount, c.ticker
FROM
    users
        INNER JOIN balances
                   ON users.userid = balances.userid
        INNER JOIN coins c on c.coinid = balances.coinid
WHERE  users.public_address = '0x8a8cB99FBE417c2fBED13B4982e4fE1BE364d58C';
UPDATE base_balances SET amount = 1000 WHERE userid = 2;
insert into buy_history(userid, coinid, sum)
values
    (1,1,1000),
    (1,1,1020),
    (1,1,1400),
    (1,2,1000);
SELECT ticker, sum(sum) as sum FROM
    buy_history
    inner join coins c on c.coinid = buy_history.coinid
                        where userid = 1 group by ticker;