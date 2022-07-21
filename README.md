# catalyst_dev
Catalyst project development


Catalyst Admin Panel
Admin panel consists of login and user token modules.
● Login:- Admin can login, which consists of Md5 algorithm to encrypt
password.(Username- super , password- super).
● User Token:- which consists of grid viewing all existing tokens

    ADD button is used to generate new token for users after adding
    username and mobile number, click on submit then a pop-up window will
    appear with user token which we can copy.
    In grid page we have option to copy and recall the generated token.The
    token will expire after 7 days.

Steps to Admin panel deploy:

1. Extract the folder CatalystAdmin.tar.gz
2. Save the folder in your go running path.
3. Import the given Mysql file (mysql_file/catalyst.sql) into database
4. Change dbuser,dbname and dbpassword in main/conf/app.conf file.
5. Open terminal go to the project path, inside main folder execute the following
command
“go run main.go”

6. Open browser type localhost:2021    