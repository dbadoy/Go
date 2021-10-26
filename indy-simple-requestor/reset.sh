rm -rf vendor
rm go.*

go mod init module
go mod vendor

cp -r connection ./vendor
cp -r did ./vendor
cp -r message ./vendor

