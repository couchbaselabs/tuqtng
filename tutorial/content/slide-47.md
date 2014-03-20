##  Messages sent by a user

List all the messages sent by a user zid-jungle-0001 to all other users 

<pre id="example">
    SELECT 
        user.name, 
        inbox.messages
            FROM jungleville AS user 
                KEY "zid-jungle-0001" 
                    LEFT JOIN jungleville_inbox AS inbox 
                        KEY "zid-jungle-inbox-" || SUBSTR(user.uuid, 11)
</pre>


