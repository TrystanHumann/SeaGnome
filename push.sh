pscp -pw $2 ./main root@$1:/root/momam
pscp -pw $2 -r ./Client/dist/* root@$1:/root/momam/momam
pscp -pw $2 ./appsettings.json root@$1:/root
pscp -pw $2 ./dbconfig.json root@$1:/root