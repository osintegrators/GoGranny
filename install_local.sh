#!/bin/bash

echo "Please enter the password you established for MySQL access (if you didn't create one, just press Enter): "
read -s pass

if ["$pass" == ""]
then
   mysql -uroot < grannyL.sql
   mysql -uroot granny < granny.sql
else
   mysql -uroot -p$pass < grannyL.sql
   mysql -uroot -p$pass granny < granny.sql
fi

