## Use Functions on the Data

Built-in functions allow greater flexibility when working with the data.

One useful built-in function is LENGTH().  If the argument passed to this function is an array or object, it returns the number of items inside.  If the argument is a string it returns the number of characters in the string.

In the example on the right we look for contacts having at least one child.

<pre id="example">
SELECT * 
    FROM contacts 
        WHERE LENGTH(children) > 0
</pre>