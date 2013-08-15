## Document Joins

Document joins allow us to take the contents of nested arrays and join them with the parent object.

Our source bucket has documents with the structure:

	{
	    "email": "dave@gmail.com",
	    "name": "dave",
	    "type": "contact"
	    "children": [ ... nested child objects ... ],
	}

After the join we will have new objects representing parent/child pairs.

	{
	  "child": {
	    "age": 17,
	    "gender": "m",
	    "name": "aiden"
	  },
	  "contact": {
	    "age": 45,
	    "children": [ ... full original array of children ... ],
	    "email": "dave@gmail.com",
	    "name": "dave",
	    "type": "contact"
	  }
	}

<pre id="example">
SELECT * 
    FROM tutorial AS contact
    	OVER contact.children AS child
        	WHERE contact.name = 'dave' 
</pre>