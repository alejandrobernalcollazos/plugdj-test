install go:
   cmd.run:
    - name: |
        wget https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz
        tar -C /usr/local -xzf go1.9.linux-amd64.tar.gz
        echo "# Adding go binaries into the path" >> /etc/profile
        echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile

set the go http server executable:
   file.recurse:
     - name: /opt/go
     - source: salt://resources/go

generate go binary from source and run it:
   cmd.run:
     - name: |
         . /etc/profile
         go build
         ./simple_http_server >/dev/null 2>&1 < /dev/null &
     - cwd: /opt/go/src/simple_http_server
