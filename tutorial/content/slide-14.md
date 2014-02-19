## Querying primary keys with KEY/KEYS expression

Specific primary keys within a bucket can be queried using the KEY/KEYS expression. The KEY clause is used to fetch exactly one document from the bucket. For fetching multiple documents the list of keys should be specified as an array.  If the KEYS clause is used then the argument must be an array, whereas if the KEY clause is used then its argument must evaluate to a single element

The query on the right fetches a list of keys from the bucket tutorial. To query a single key try changing the KEYS expression to KEY "dave"

An arbitary expression can be used as an argument to the KEY/KEYS clause as long as it evaluates to an array or a single element

<pre id="example">
    SELECT fname, email
        FROM tutorial 
            KEYS ["dave", "ian"]
</pre>
