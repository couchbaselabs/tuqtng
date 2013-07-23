## A friendly language

If you know SQL, you already know a lot of Couchbase Query Language.

A basic query has three parts to it:

* SELECT - Portions of document to return
* FROM - The bucket (data store) to work with
* WHERE - Conditions the document must satisfy

Only SELECT clause is required. The wildcard * selects all parts of the document. Queries can return a collection of differently shaped documents or fragments. However, they will all match the conditions you specified in the FROM clause.

Remember, there **IS NO SCHEMA** in Couchbase. You don't lose any flexibility you love about Couchbase.

If you change the * to an element like 'children', you will see the query return a collection of appropriate fragments of each document.