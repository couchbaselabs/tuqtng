#  Messages sent by a user

List all the messages sent by a user zid-pork-0001 to all other users 

<pre id="example">
    SELECT 
        user.name, 
        inbox.messages
            FROM porkville AS user 
                KEY "zid-pork-0001" 
                    LEFT JOIN porkville_inbox AS inbox 
                        KEY "zid-pork-inbox-" || SUBSTR(user.uuid, 9)
</pre>


