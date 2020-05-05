# Odin - Project Creator
Odin template for create your project creator

### Configure odin
```bash
odin start
```
This command will create odin settings on your computer 
```bash
$HOME/.odin/config.yml
```
### Cli
```bash
$ odin --help
Create your application

Usage:
  odin [command]

Available Commands:
  config      Configure odin
  create      Create new project
  help        Help about any command
  start       Use for start configurations odin

Flags:
  -h, --help   help for odin

Use "odin [command] --help" for more information about a command.
```

Odin use modules for struct projects. The confige default it is **$HOME/.odin/modules/default.yml**
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

- create > directories : Created directories when execute command for create project
- create > files: Created files with  templates based in resources
- resources: Absolute path of templates

### Command for create my project
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
To create your own project module, use the command below
```bash
$ odin config module module1       
Created file $HOME/.odin/modules/module1.yml
Edit $HOME/.odin/modules/module1.yml
```
Edit your module for create your struct

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
resource: ./ # Absolute Path of templates
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

