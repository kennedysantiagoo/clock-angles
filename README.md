# clock-angles
A simple rest API that, given an hour and minute number, retrieves the angle between the hour pointer and the minute pointer. 

# This API contains a GET endpoit which accept 'hour' and 'minute' in the url path.
 Ex: localhost:3000/clock/{hour}/{minute}  

# This API is meant to run locally and the default port is set to ':3000'

# This API is meant to use a conteinerized Postgres image and the connection information can be found in the connection file. 

# There's a sql dump file in a directory called 'sql'

# Steps to get database ready and populated:

# 1 Download a Postgres image in docker

# 2 Open your terminal

# 3 Create, run and export a postgres container: 
 'docker run --name meu-banco -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres'

# 4 Open the Postgres image bash terminal: 
   'sudo docker exec -it yourPostgresContainerID bash' 

# 5 Enter in the postgres user profile: 
   'psql -U postgres'

# 6 Create a postgres database:
   'CREATE DATABASE mybanco'

# 7 Exit the postgres user:
  '\q'

# 8 Create a table and import data to database:
 'sudo -u postgres clock_angles < path/on/your/machine/to/the/dump/file/dump.sql'



 
   
   

