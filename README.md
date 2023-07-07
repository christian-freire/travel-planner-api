# Travel Planner API 🛫
by: Christian Freire

- [Description](#description)
- [Requirements](#requirements)
- [Setting up the development environment](#setting-up-the-development-environment)
- [Project Structure](#project-structure)
- [Packages](#packages)
- [Run the project](#run-the-project)

## Description

A trip planner API, using an OPENAI API implementation. The user makes a request with the start date of the trip, the end date of the trip, the city of origin and the city of destination, and receives from the API a planner with things to do and places to visit for every day of the trip.


## Requirements

To run the project in your local machine, will be necessary the following requirements:

- #### `Go (Programming language)`
The project is developed in Go. To install the language in your machine, acess https://go.dev/doc/install and install the most recent version of the language. To verify the version and if the installation was run correctly, open a prompt and run `go version`.

- #### `Insomnia (Requests and API tests)`
To make the requests to the server and test the API endpoint, we use Insomnia. To install, acess https://insomnia.rest/download and download the most recent version.  

## Setting up the development environment

Follow the next steps to configure the development environment:

#### `1. Clone the project repository in your local machine: `

```bash
git clone https://github.com/christian-freire/travel-planner-api.git
```

#### `2. Update the project dependencies in your local machine: `
Open a terminal inside the project directory and run:

```bash
go mod tidy
```
This command line will update the [packages](#packages) used in the project.

#### `3. API KEY: `
Inside the project directory, create a `configs` directory, and then, create a `env.yml` file. This file will contain the OPENAI API KEY, that you can generate by creating or logging into your account on the OPENAI platform, through this link: https://platform.openai.com/account/api-keys

After, take the API KEY and implement it in the `env.yml` you created: 

```bash
APIKEY: "YOUR_API_KEY"
```

⚠️ `Recommendation:` Is not a good pratice to leave the API KEY in the hard-code or commit it to the remote repository. So, you can create a `.gitignore` file and drop the `env.yml` inside it. 

#### `4. Configure Insomnia: `
Open your Insomnia, then create a `new request collection`. In the endpoint request method, set `POST`.

## Project Structure
The structure of the project must follow the layout recommended by https://github.com/golang-standards/project-layout/tree/master

```
├── cmd
|   └── main.go
├── configs
|   └── env.yml
├── lib
|   └── travel.go
├── routes
|   └── routes.go
├── services
|   └── open_ai.go 
├── .gitignore
├── go.mod
├── go.sum     
└── README.md
```

- `cmd`: This directory contains the main file for the project, the `main.go` file. This file contains the router initialization for the endpoint, and call the `Travel` function.
- `configs`: This directory contains the `env.yml` file, that contains the API KEY.
- `lib`: This directory contains the `travel.go` file. This file is responsible for creating the struct with the data that we must receive from our user, and create the standard question that will be passed to CHAT GPT, already implementing the data that was received from the request.
- `routes`: This directory contains the `routes.go` file. This file contais our POST endpoint implement.
- `services`: This directory contains the `open_ai.go` file. This file implements the OpenAI function, who read the question and print the response in terminal.

## Packages

In this project, i have used some Go packages, they are: 

- `GIN`: Gin is a web framework written in Go. It features a Martini-like API, but with performance up to 40 times faster than Martini. I have used Gin to create the router and the POST endpoint `"/trip"`

- `VIPER`: The Viper package in Go is a library used to handle application configurations. It allows reading and writing configuration files, as well as accessing and manipulating these configurations at runtime. Viper makes loading and using settings in a Go app flexible and convenient. I have used Viper to read my API KEY, in `env.yml`.

- `OPENAI CLIENT`: This library provides unofficial Go clients for OpenAI API. I have used to recieve the CHAT GPT response for the question.

## Run the project
To run the project is simple:

#### `1. Open a terminal in the project directory, and run: `
```bash
go run cmd/main.go
```
This command will upload the server to recieve requests. 

#### `2. Open Insomnia to push the request: `
In Insomnia, configure the endpoint adress to our `/trip` endpoint. Acording to our configuration, the server is in `http://localhost:8080` port. 

#### `3. Send the request: `
So, in the field `JSON` in Insomnia, fill in a valid `JSON` for the request, according to what is configured in the `travel.go`, like this:

```bash
{
	"start_date": "01/07/2023",
	"end_date": "30/07/2023",
	"city_origin": "São Paulo",
	"city_destination": "San Diego"
}
```
Then, send the request.

#### `4. Response: `
You will recieve your response, in the terminal that you have open to upload the server!

```bash
[GIN-debug] Listening and serving HTTP on :8080
Claro! Aqui está um roteiro de viagem para você aproveitar cada um dos dias em San Diego:

Dia 1 (01/07/2023):
- Chegada em San Diego
- Check-in no hotel e descanso
- Passeio pela região de Gaslamp Quarter
- Jantar em um dos restaurantes locais

Dia 2 (02/07/2023):
- Visita ao Balboa Park e seus museus
- Passeio pelo San Diego Zoo
- Tarde relaxante nas praias de La Jolla
- Jantar em um restaurante à beira-mar

Dia 3 (03/07/2023):
- Excursão para o USS Midway Museum
- Passeio pelo Seaport Village
- Aproveite o fim do dia em Ocean Beach
- Jantar em um restaurante de frutos do mar

Dia 4 (04/07/2023):
- Celebração do Dia da Independência dos Estados Unidos
- Participação nos eventos e fogos de artifício na região
- Passeio noturno pela praia de Coronado

Dia 5 (05/07/2023):
- Passeio de barco pela Baía de San Diego
- Visita ao Maritime Museum of San Diego
- Caminhada pelo calçadão de Mission Beach
- Jantar em um restaurante mexicano

Dia 6 (06/07/2023):
- Conheça o Parque Estadual de Anza-Borrego Desert
- Faça uma trilha pela paisagem deslumbrante
- Desfrute de um piquenique no parque
- Retorno a San Diego e jantar em um restaurante local

Dia 7 (07/07/2023):
- Passeio pelo bairro de Little Italy
- Visita ao Museu do Centro de História de San Diego
- Relaxe nas praias de Pacific Beach
- Jantar em um restaurante com vista para o pôr do sol

Dia 8 (08/07/2023):
- Explore o bairro histórico de Old Town San Diego State Historic Park
- Almoço em um autêntico restaurante mexicano
- Visita ao San Diego Natural History Museum
- Jantar em um dos restaurantes tex-mex locais

Dia 9 (09/07/2023):
- Excursão ao Cabrillo National Monument
- Aprecie as vistas panorâmicas do Farol de Point Loma
- Visite o Museu do Parque Nacional Cabrillo
- Jantar em um restaurante de frutos do mar

Dia 10 (10/07/2023):
- Aproveite um dia de compras no Fashion Valley Mall
- Faça uma pausa para o almoço em um dos restaurantes do shopping
- Visite o Museu de Arte de San Diego
- Jantar em um restaurante com culinária internacional
```

# And that's the end of the project! To the next 🚀

Made and created by: Christian Freire

