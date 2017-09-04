install postgresql:
   cmd.run:
    - name: |
        apt-get update
        apt-get install postgresql postgresql-contrib -y

set sql files into the machine:
   file.recurse:
    - name: /opt/postgresql
    - source: salt://resources/postgresql

set the user and database:
   cmd.run:
    - name: |
        sudo -u postgres psql --file=/opt/postgresql/create_user.sql 
        sudo -u postgres psql --file=/opt/postgresql/create_db.sql 
