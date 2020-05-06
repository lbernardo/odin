# Odin - Project Creator
Odin usa modelos escritos em yaml para criar sua estrutura de projeto

### Configure Odin
O primeiro passo é executar o start 
```bash
$ odin start
```
O comando irá criar as configurações iniciais do odin no seu diretório $HOME
```bash
$HOME/.odin/config.yml
```

### CLI
```bash
$ odin --help
Create your application

Usage:
  odin [command]

Available Commands:
  config      Configure Odin
  create      Create a new project
  help        Help about any command
  start       Use for start configurations Odin

Flags:
  -h, --help   help for odin

Use "odin [command] --help" for more information about a command.
```

Odin usa modulos para estruturar seu projeto. A configuração inicial fica presente em **$HOME/.odin/modules/default.yml**
```yaml
create:
  directories: # Cria os diretorios do projeto ao executar odin create myproject
   - cmd
   - internal
   - internal/primary
   - internal/secondary
   - pkg
   - pkg/handler
   - pkg/models
   - pkg/repositories
  files: # Cria os arquivos do projeto ao executar odin create myproject
   - ${resource}/default/main.tpl:cmd/main.go
resource: ./ # Caminho absoluto de seus templates
```

### Criando um novo projeto
Ao executar o comando abaixo você irá criar um novo projeto com base no modulo padrão (**$HOME/.odin/modules/default.yml**)
```bash
$ odin create myproject  github.com/lbernardo/myproject
Created  myproject
Created  myproject/cmd
Created  myproject/internal
Created  myproject/internal/primary
Created  myproject/internal/secondary
Created  myproject/pkg
Created  myproject/pkg/handler
Created  myproject/pkg/models
Created  myproject/pkg/repositories
Created  myproject/cmd/main.go
```
O comando lê o arquivo yaml e vê o que precisa ser criado, tanto diretórios como arquivos

## Create your module
Para criar sua própria estrutura de projetos, você precisa criar um módulo. Para isso use o comando abaixo:
```bash
$ odin config module module1
Created file $HOME/.odin/modules/module1.yml
Edit $HOME/.odin/modules/module1.yml
```
Use o comando abaixo para definir como  modulo principal
```bash
$ odin config default module1
```

Agora você precisa editar o arquivo criado: 

**$HOME/.odin/modules/module1.yml**
```yaml
create:
  directories: # Create directories for project
   - cmd
   - internal
   - internal/primary
   - internal/secondary
   - pkg
   - pkg/handler
   - pkg/models
   - pkg/repositories
  files: # Create files for project
   - ${resource}/main.tpl:cmd/main.go
resource: /home/lbernardo/workspace/mytemplate # Absolute Path of templates
commands:
  - cmd: handler
    description: "Create new handler"
    args:
      - name: name
        description: Name of handler
    #directories: #create directories for command
    #  - 
    files:
      - ${resource}/handler.tpl:pkg/handler/${name}.go
  - cmd: model
    description: "Create new model"
    args:
      - name: name
        description: Name of model
    files:
      - ${resource}/model.tpl:pkg/models/${name}.go
```
**É obrigatório criar seu diretório resource. Ele deve conter os templates do seu projeto**

Por exemplo:

/home/lbernardo/workspace/mytemplate/handler.tpl
```
package handler

type ${name} struct {
    
}

func New${name}() *${name} {
    return &${name}{
        
    }
}
```

Editando o arquivo yml o parâmetro resource com o valor **/home/lbernardo/workspace/mytemplate**

Quando executar o comando **odin create handler --name MyHandler** irá criar o seguinte arquivo no seu projeto:

**pkg/handler/myhandler.go**
```go
package handler

type MyHandler struct {
    
}

func NewMyHandler() *MyHandler {
    return &MyHandler{
        
    }
}
```

**utilize ${VARIAVEL} para substituir os argumentos no conteudo do seu template**

### Create new Command
Para criar um novo comando no seu modulo **$HOME/.odin/modules/{MODULE}.yml**  utilize a seguinte estrutura:
```yaml
# ....
commands:
  - cmd: {Nome do comando (obrigatório)}
    description: {Descrição do comando (obrigatório)}
    args: # Lista de argumentos
      - name: {Nome do argumento (obrigatório)}
        description: {Descrição do argumento (obrigatório)}
        value: {Valor padrão do argumento}
    directories: # Lista de diretorios que deve ser criado
      - 
      -
      -
    files: # Lista de arquivos que deve ser criado com base no template. Use ${resource} e ${VARIABLE NAME}. ${resource} é obrigatório.
      - ${resource}/handler.tpl:pkg/handler/${VARIABLE NAME}.go
```

Exemplo
**$HOME/.odin/modules/mymodule.yml**
```
create:
  directories: # Create directories for project
   - cmd
   - internal
   - internal/primary
   - internal/secondary
   - pkg
   - pkg/handler
   - pkg/models
   - pkg/repositories
resource: /home/lbernardo/workspace/mytemplate # Absolute Path of templates
commands:
  - cmd: example
    description: "Create my example"
    args: # Array of arguments 
      - name: name
        description: "Name"
     -  name: lastname
        description: "Lastname"
    directories: # List of directories for created
      - cmd/cli
      - cmd/http
    files: # List of files for created with the base template. Use ${resource} and ${VARIABLE NAME}. ${resource} is required.
      - ${resource}/cli/main.go:cmd/cli/main.go
      - ${resource}/http/main.go:cmd/http/main.go
```

**/home/lbernardo/workspace/mytemplate/cli/main.go**
```
package main

import "fmt"

func main() {
    fmt.Println("${name} ${lastname}")
}
```
**/home/lbernardo/workspace/mytemplate/http/main.go**
```
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi ${name} ${lastname}")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Execute**
```
$ odin example --name "Lucas" --lastname "Bernardo"
```

**cmd/cli/main.go**
```
package main

import "fmt"

func main() {
    fmt.Println("Lucas Bernardo")
}
```

**cmd/http/main.go**
```
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi Lucas Bernardo")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```




