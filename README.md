# api-go


## iniciando o projeto
```
go mod init github.com/alexiaassumpcao/api-go
```

## instalar fiber
```
go get -u github.com/gofiber/fiber/v2
```

## instalar dependencias 
```
go mod download
```

## atualizar dependencias
```
go mod tidy
```

## rodar projeto na porta 8080
```
go run cmd/main.go
```


## rodar atraves do Makefile
```
make run
```

- OBS: eu nao sabia que era dificil rodar o comando make no windows então não consegui testar