#!/bin/sh

cat <<EOT > /usr/local/bin/gotest
#!/bin/sh

if [ ! -z "\$1" ]
then
    go test -failfast ./\`dirname \$1\`/... -p 1
fi
EOT

chmod +x /usr/local/bin/gotest
gomon -build="gotest \$FILE" -command="go run main.go"
