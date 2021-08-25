# Golioth Challange

Project developed by João Miguel to solve the Golioth Challenge.

Following it will be described how the project was structured and hot to run it.

## Project Structure


A API de trsnferência entre contas possui as seguintes entidades (modelos):

    * Account
    * Transfer
    * Login
    * Token 
    * StorageInMemory

### Entidade Account

Esta entidade é responsável por agrupar os dados referentes à uma conta.

    * Id         (int)     -> id da conta
    * Name       (string)  -> Nome do titular da conta
    * Cpf        (string)  -> Cpf do titular da conta
    * Secret     (string)  -> Senha da conta
    * Balance    (float64) -> Saldo atual da conta
    * Created_at (time)    -> Data e hora que a conta foi criada


### Entidade Transfer

Esta entidade é responsável por agrupar os dados referentes a uma transferência entre contas.

    * Id                     (string) -> UUID da transferência
    * Account_origin_id      (int)    -> conta de origem da transferência
    * Account_destination_id (int)    -> conta de destino da transferência
    * Ammount                (float)  -> valor da transferência
    * Created_at             (time)   -> Data e hora que a conta foi criada

Toda vez que uma nova transferência tentar ser realizada, o saldo da conta de origem será avaliado e, se houver saldo maior ou igual ao que se deseja transferir, então será permitido. Do contrário não será permitido e um erro será retornado ao usuário da API.

Quando for possível realizar a transferência entre contas, o sistema irá registrar duas componentes de Transferência no banco de dados: débito e crédito.

Registrará na lista de transferências da conta de origem uma nova transferência com valor negativo, significando débito. E, de forma simétrica, registrará na lista de transferẽncias da conta de destino uma nova transferência com o mesmo valor, porém positivo, isto é, um crédito.

Assim, dessa forma, é possível rastrear o caminho qual o atual salde de uma determinada conta percorreu, entre débitos e créditos.

### Entidade Login

Esta entidade é responsável por agrupar os dados referentes a um login.

    * Cpf    (string) -> cpf da conta
    * secret (string) -> senha da conta

Toda vez que um usuário se logar na API, o cpf e a senha deverão ser fornecidos. O sistema então autenticará o usuário conferindo se para aquele CPF fornecido, a senha fornecida bate com a senha armazenada na conta do usuário. Se sim, então ele será autenticado e autorizado.


### Entidade Token

Esta entidade é responsável por agrupar os dados referentes a um token.

    * Token           (string) -> string do token gerado
    * Cpf             (string) -> cpf da conta
    * AccountOriginId (int)    -> id da conta

Toda vez que um usuário se logar na API, um novo token será gerado, a string do token será inserida no slice de tokens do banco de dados em memória e a string do token será retornada ao usuário para realizar futuras requisições.

### Entidade StorageInMemory

A entidade StorageInMemory chama-se assim porque optei por realizar um banco de dados em memória. A desvantagem dessa abordagem é que ao encerrar o sistema da API, todos os dados serão perdidos, porém, como não é um sistema que necessita manter dados persisitidos para além de uma avaliação, então não haverá problema nisso.

A entidade StorageInMemory armazena 3 estruturas de dados necessárias para a API toda funcionar:

    * accounts  ([]Accounts)         -> um slice de Accounts
    * transfers (map[int][]Transfer) -> um map cujas chaves são os id's de cada conta cadastrada e o valor é um slice de Transfer (array de Transfer)
    * tokens    ([]tokens)           -> um slice de Tokens

Além disso a entidade StorageInMemory implementa uma interface *Storage* com funções típicas de banco de dados como SaveAccount, SaveTransfer, FindToken, etc.

A estrutura accounts armazena todas as contas criadas.

A estrutura transfers armazenará sempre 2 componentes (débito/crédito) para cada transferência realizada.

toda vez que uma nova conta for criada na API, uma chave nova será criada no map de transfer sendo o id dessa nova conta a chave a ser usada no map e um slice vazio será criado como valor dessa chave.

Por exemplo:

    A estrutura transfers possui o seguinte estado no momento:

        { }

    Alguem criou uma nova conta com id 1 e, depois, com id 2. Então duas novas chaves serão criadas em transfers:

        {
            "1": [],
            "2": []
        }

    E assim, sucessivamente.

    No momento em que uma nova transferencia for realizada, por exemplo, da conta 2 para a conta 1, duas componentes de trasnfers serão criadas:

        {
            "1': [
                {
                    "Id": 087c5544-2504-434c-8a78-dd6346879547,
                    "Account_origin_id": 1,
                    "Account_destination_id": 2,
                    "Ammount": 256.67,
                    "Created_at": "2021-07-20T18:36:50.728821618-03:00"
                }
            ],
            "2': [
                {
                    "Id": 087c5544-2504-434c-8a78-dd6346879547,
                    "Account_origin_id": 1,
                    "Account_destination_id": 2,
                    "Ammount": -256.67,
                    "Created_at": "2021-07-20T18:36:50.728821618-03:00"
                }
            ],
        }

    Repare que na conta 2 o valor é negativo enquanto que na conta 1 é positivo: débito e crédito.
        

## How to run the project

The project is containerized using a Golang image and for run it you must follow the next steps:

* First of all, clone the repository into your machine
> git clone https://github.com/miguel91it/backend-challenge.git

* Next you must enter in the root folder:
> cd backend-challenge

* Once inside at the root folder of the project, you must build and run all the containers (mongodb, gateway and devices) using docker compose:
> sudo docker-compose up --build gateway devices

The command above will build and run all the 3 containers but only gateway and devices containers will be attached and logged into standard output. This way you'll be able to see both logs during the gateway and devices containers executions.

* After the succesfull conclusion of the above command, you'll be able to attach to the gateway container and see its logs:
> sudo docker attach gateway

* Finally, to stop and remove all containers
> sudo docker-compose rm -sfv


## How to test the Gateway Server

It's possible to test manualy the telemetry endpoint of the API using command line tools like Curl or GUI tools like Postman.

Following we have some example of how to send a telemetry data to the gateway using `Curl`:

> POST: localhost:28000/api/v1/WeatherTelemetry

```
    Example of Request body:

        {
            "id":              "1",
            "timestamp":       12345678,
            "soil_moisture":   10,
            "ext_temperature": 10.1,
            "ext_humidity":    1.3
	    }
```

Full `Curl` command:

> curl --request POST 'localhost:28000/api/v1/WeatherTelemetry' \ \
--header 'Content-Type: application/json' \ \
--data-raw '{ \
		"id":              "02:42:ac:17:00:03", \
		"timestamp":       1629855252, \
		"soil_moisture":   10,\
		"ext_temperature": 10.1,\
		"ext_humidity":    1.3\
	}'

Experiment change soil_moisture, ext_temperature and ext_humidity values.