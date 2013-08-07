## Performing Simple Arithmetic

We can also perform basic arithmetic in our expressions.

In the example on the right we give each user a minimum score of 100 by adding that to their score.

The common arithmetic operations +, -, *, / and % are supported.

Try using a different arithmetic operation.

<pre id="example">
SELECT name, score + 100 
    FROM game 
        WHERE score + 100 > 108
</pre>