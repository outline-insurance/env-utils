apt install curl

curl -O https://dl.google.com/go/go1.14.linux-amd64.tar.gz

tar xvf go1.14.linux-amd64.tar.gz

chown -R root:root ./go
mv go /usr/local

export PATH=$PATH:/usr/local/go/bin

cd /tmp
go build 
