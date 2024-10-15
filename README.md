- this is the first time i tried to build a login session.
- to run this project. First, make sure you understand basic concepts of backend and frontend development. Build a database in postgre (well, you can also use other type
of SQL like mysql, mssql,..., by changing the connection in main.go (psqlconn)).
- Also, make sure you have downloaded/get necessary library. Figure them out by reading golang files imports.
- After finishing them:
  + First, open the terminal, type cd server. After that, in terminal of react-auth directory, run 'air'. It will run the backend server
which is necessary for connecting APIs to.
  + Then, create a new terminal (ctrl shift `) and type: cd react-auth => type : npm run dev.
  + if nothing bad happens, it will show a very basic screen for login/register/logout testings. It will directly update registered accounts into the database.
