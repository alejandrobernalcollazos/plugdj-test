install postgresql:
   cmd.run:
    - name: |
        apt-get update
        apt-get install postgresql postgresql-contrib -y

set sql files into the machine:
   file.recurse:
    - name: /opt/postgresql
    - source: salt://resources/postgresql

set the bootstrap that create user db and tables:
   cmd.run:
    - name: |
        sudo -u postgres psql --file=/opt/postgresql/bootstrap.sql 

bind postgresql to * instead of localhost and allow remote connections:
   cmd.run:
    - name: |
        cp -f /opt/postgresql/postgresql.conf /etc/postgresql/9.5/main/postgresql.conf 
        cp -f /opt/postgresql/pg_hba.conf /etc/postgresql/9.5/main/pg_hba.conf

restart postgresql:
   module.run:
    - name: service.restart
    - m_name: postgresql
