FRONT END
You have to build this Web Application:

Web App Framework:
Please use a Web App Framework of your choice.
The web application will receive a user input using format: NAME[space]AGE[space]CITY
For example: CUT MINI 28 BANDA ACEH
The web application then submits the data to the back-end application using JSON format.

---

BACK END
You have to build this RESTful Web Service:

Back-end App Framework:
Please use a Back-end App Framework of your choice. We prefer Golang or Java Spring Boot.
After the data is submitted, the back-end app parses it to get three data: Name, Age, and City.
Those data are inserted into a MySQL or PostgreSQL table whose fields are:
ID (auto increment)
NAME
AGE
CITY
String processing must be right-to-left, and done character by character. Each character may only be read once.
The back-end app must be able to handle names and cities using multiple words (no limit). No regex is allowed. The back-end app also has to convert the name and city data into upper case before inserting them into the table.
As for age data, the back-end app must be able to handle Indonesian user's common mistake, which is adding TAHUN, THN, or TH string after the age data. No regex is allowed. For example:
28 TAHUN, 28 THN, 28 TH
28TAHUN, 28THN, 28TH
You may not use string replacement functions anywhere in the parsing logic.
If the city name is a province's capital, please add its province's name after the city name. The list of provincial capitals must be provided at runtime via configuration.
Please make sure the back-end app is able to handle TAHUN, THN, and TH strings case-insensitively.
Create a solution without using global variables and limit the use of local variables to a maximum of 5 variables for the entire parsing process. You may exceed the local variable limit only if you explicitly explain (as comments/remarks) why it improves correctness or clarity.
The final result saved to the database must be in Fixed-Width String format. Name must be 30 characters, Age must be 3 characters, and City must be 20 characters.
