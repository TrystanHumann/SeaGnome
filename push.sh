pscp -pw $2 main momam@$1:/home/momam
pscp -pw $2 -r Client/dist/* momam@$1:/home/momam/momam
pscp -pw $2 appsettings.json momam@$1:/home/momam
pscp -pw $2 dbconfig.json momam@$1:/home/momam