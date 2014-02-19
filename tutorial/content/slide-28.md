## Chaining JOINs

JOINS can be chanined and combined with UNNEST as many times as desired. 

In the example on the right we perform an indocument join to get an object that contains the order Ids along with the users details and that inturn is then JOINED with the order decriptions from the users_with_orders bucket.


<pre id="example">
    SELECT  u.personal_details.display_name name, s as order_no, o.product_details  
        FROM users_with_orders u KEY "Aide_48687583" 
            UNNEST u.shipped_order_history s 
                JOIN users_with_orders o KEY s.order_id
</pre>
