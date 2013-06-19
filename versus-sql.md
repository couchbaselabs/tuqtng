Introduction
-------------
This is a basic introduction of a mapping of SQL concepts to UnQL 2013. It is neither a tutorial nor a complete introduction to UnQL, and so is suitable for a casual reading. For anything more serious, pelase refer to [UnQL 2013 Specification](unql-2013.md) and the to be written tutorials.

Fundamental differences
------------------------
The most important difference versus traditional SQL and UnQL are not lingual but the data model. In traditional SQL, data is constrained to tables with a uniform structure, and many such tables exist.


    EMPLOYEE
    ------------------
    Name | SSN | Wage
    ------------------
    Siri | 234 | 123 |
    Steve| 123 | 456 |
    ------------------
    
    SCHEMA:
    Name -> String of width 100
    SSN -> Number of width 9
    Wage -> Number of width 10

    EMPLOYERS:
    -------------------------------------
    Name_Key | Company   | Start | End  |
    -------------------------------------
    Siri     | Yahoo     | 2005  | 2006 |
    Siri     | Oracle    | 2006  | 2012 |
    Siri     | Couchbase | 2012  | NULL |
    -------------------------------------


In UnQL, the data exists as free form documents, gathered as large collections called buckets. There is no uniformity and there is no logical proximity of objects of the same data shape in a bucket.  

    (HRDATA bucket)
    {
        'Name': 'Siri'
        'SSN': 234
        'Wage': 123
        'History': 
        [
          ['Yahoo', 2005, 2006],
          ['Oracle', 2006, 2012],
          ['Couchbase', 2012, null]
        ]
    },
    {
        'Name': Steve
        'SSN': 123,
        'Wage': 456,
    }



Result Set differences
----------------------
When one runs a query in SQL, a set of rows, consisting of one or more columns each is returned, and a header can be retrieved to obtain metadata about each column. It is generally not possible to get rowset where each row has a different set of columns.


    SELECT Name, Company 
    FROM Employee, Employers
    WHERE Name_Key = Name
    
    ------------------
    Name | Company   |
    ------------------
    Siri | Oracle    |
    Siri | Yahoo     |
    Siri | Couchbase |
    ------------------

In UnQL, a query returns a set of documents. The returned document set may be uniform, but it need not be. Typically, sepecifying a SELECT with fixed set of attribute ("column") names results in a uniform set of documents. SELECT with a wildcard("*") results in non-uniform result set. The only guarentee is that every returned document meets the query criteria.

    SELECT Name, History
    FROM HRDATA
    
    {
        'Name': Siri,
        'History':
        [
          ['Yahoo', 2005, 2006],
          ['Oracle', 2006, 2012],
          ['Couchbase', 2012]
        ]
    }
    {
        'Name': 'Steve'
    }


Like SQL, UnQL allows renaming fields using the AS keyword. However, UnQL also allows reshaping the data, which has no analog in SQL. This is done by embedding the attributes of the statement in the desired result object shape.


    SELECT { Name, History, {'FullTime': true} AS 'Status'}
    FROM HRDATA
    
    {
        'Name': 'Siri',
        'History':
        [
          ['Yahoo', 2005, 2006],
          ['Oracle', 2006, 2012],
          ['Couchbase', 2012]
        ],
        'Status': { 'FullTime': true }
    }
    {
        'Name': 'Steve',
        'Status': { 'FullTime': true }
    }


Selection differences
---------------------
