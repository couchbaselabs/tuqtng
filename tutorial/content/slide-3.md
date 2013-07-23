## Complex Data - Basic Operators

As documents can have nested child elements and embedded arrays, a few additional operators are needed. The '.' operator is used to refer to children, and the '[]' is used to refer to an element in an array. 

In the example on the right, we fetch the first child of each contact. Notice that the child traversal and array element dereference happens in the FROM clause. Here, the FROM clause is used to synthesize the input to the query as a collection of first child of each contact.

An alternate way of doing this would be to set the FROM clause to merely fetch children. The SELECT clause could then fetch the first child and extract the name attribute from it. 

Try rewriting the query to do the path expression in the SELECT clause. The answer is on the next slide.



