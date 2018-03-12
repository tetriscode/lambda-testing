GOOS=linux go build -o main
zip deployment.zip main

aws lambda update-function-code \
--region us-east-1 \
--function-name wfs3 \
--zip-file fileb://./deployment.zip \
