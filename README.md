# user-services


colocar as migrations para rodar no dockerfile
 
Rodar o migrate usando a variável de ambiente
Agora, execute o comando referenciando a variável DATABASE_URL:

migrate -database "$DATABASE_URL" -path db/migrations up
