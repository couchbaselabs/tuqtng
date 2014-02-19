## Cross Document JOINS with KEY/KEYs operator

N1QL supports the JOIN clause that allows you to create new input objects by combining two or more source objects. 

For e.g. Lets assume we have two buckets users_with_orders and orders_with_users. The bucket users_with_orders contains user profiles along with the orders Ids of the orders that they placed. 

To see what a typical document in the users_with_orders bucket looks like, copy paste the following query to the window on the right

```
SELECT  * FROM  users_with_orders KEY  "Elinor_33313792"
```

The bucket orders_with_users contains the description of a particular order placed by a user. Now in order to view a user-profile along with the orders that she has placed in the past we can use the JOIN clause to accomplish that. 

The example on the right combines a users profile document referenced by the key "Elinor_33313792", with the description of the orders that the user has placed. Note that corss document joins must use the KEY/KEYs clause. In the example on the right we get a list of orders by unrolling the shipped_order_history array and pass that as in input to the JOIN clause which causes the users personal details to be combined with the order descriptions.

<pre id="example">
    SELECT user.personal_details, orders
        FROM users_with_orders user 
            KEY "Elinor_33313792" 
                JOIN users_with_orders orders 
                    KEYS ARRAY s.order_id FOR s IN user.shipped_order_history END
</pre> 
