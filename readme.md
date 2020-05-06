# Odin - Project Creator
Odin template for creating your project creator

### Configure Odin
```bash
Odin start
```
This command will create Odin settings on your computer 
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

Odin uses modules for project struct. The config default it is **$HOME/.odin/modules/default.yml**
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
   - ${resource}/default/main.tpl:cmd/main.go
resource: ./ # Absolute Path of templates
```

- create > directories: Created directories when executing command for creating project
- create > files: Created files with  templates based in resources
- resources: Absolute path of templates

### Command to create my project
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
The command read the module's default settings to know which directories and files to create must be created and which files to create

## Create your module
To create your  project module, use the command below
```bash
$ odin config module module1       
Created file $HOME/.odin/modules/module1.yml
Edit $HOME/.odin/modules/module1.yml
```
Edit your module to create your struct

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
**Is required create your  resource path  and alter config with absolute path**

For example:

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

When to execute **odin create handler --name MyHandler** your create:
```go
package handler

type MyHandler struct {
    
}

func NewMyHandler() *MyHandler {
    return &MyHandler{
        
    }
}
```

**use ${VARIABLE} for replacing arguments on the command**

### Create new Command
For creating a new command, on your **$HOME/.odin/modules/{MODULE}.yml** create next struct
```yaml
# ....
commands:
  - cmd: {NAME TO COMMAND (required)}
    description: {DESCRIPTION TO COMMAND (required)}
    args: # Array of arguments 
      - name: {NAME TO ARGUMENT (required)}
        description: {DESCRIPTION TO ARGUMENT (required)}
        value: {VALUE TO DEFAULT ARG}
    directories: # List of directories for created
      - 
      -
      -
    files: # List of files for created with the base template. Use ${resource} and ${VARIABLE NAME}. ${resource} is required.
      - ${resource}/handler.tpl:pkg/handler/${VARIABLE NAME}.go
```

Example:
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




