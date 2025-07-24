# Codivy - Automatic repository creator (I know this is in the description)

[![Go Version](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/github/license/JuaanReis/Codivy)](https://github.com/JuaanReis/Codivy/blob/main/LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen)]()
[![Made By](https://img.shields.io/badge/made%20by-juan%20reis-black)](https://github.com/JuaanReis)
[![Codivy](https://img.shields.io/badge/project-codivy-critical)]()

> create your repositories automatically without much work (just answer a few interview-style questions)

## How it works
1. Codivy makes a POST request to the GitHub API with your information (no one else has access except GitHub)
2. Then just access the link (it is in the terminal at the end of the request)

## Why use it?
1. why did I do
2. facilitates the creation of repositories
3. you just need a token and a cool name idea (not my case)

## How use it?
1. Create a token with admin permission
2. Clone this repository
```bash
git clone https://github.com/JuaanReis/Codevy.git
go run main.go
```
3. Answer the questions
4. To continue your project

## Project structure
```
Codivy/
├── .cli/
│     └── .app/ 
│          └── main.go       <- main app entry point
├── internal/
│      └── github/
│            └── requests/
│                  └── request.go    <- GitHub request logic
├── go.mod
└── README.md
```

## Contributions
If you want to help with this project feel free to do so, no rules, just know Go.