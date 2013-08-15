## A friendly language

If you know SQL then Couchbase Query Language will be familiar to you.

A basic query has three parts to it:

* SELECT - Parts of document to return
* FROM - The data bucket, or data store to work with
* WHERE - Conditions the document must satisfy

Only a SELECT clause is required in a query. The wildcard * selects all parts of the document. Queries can return a collection of different document structures or fragments. However, they will all match the conditions in the WHERE clause.

Remember there **IS NO SCHEMA** in Couchbase. You don't lose any flexibility you love about Couchbase.

If you change the * to a document field such as 'children', you will see the query return a collection of appropriate fragments of each document.

Try the next sample query where we find all documents where the name is 'ian.'

<pre id="example">
SELECT *
  FROM tutorial
    WHERE name = 'ian'
</pre>
