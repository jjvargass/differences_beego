# differences_beego
It's un repo compare the create a model and controller based on the database schema y with the command **bee generate**

1. The table sql

  ![Table user bd](/img/001.png)

2. Create API beego  based on the database schema by providing database conn:

       $ bee api differences_beego -driver=postgres -conn=postgres://MyUsuarioBD:MyPassDB@127.0.0.1/differences?sslmode=disable

      get

       ______
       | ___ \
       | |_/ /  ___   ___
       | ___ \ / _ \ / _ \
       | |_/ /|  __/|  __/
       \____/  \___| \___| v1.10.0
       2019/04/09 15:07:01 INFO     ▶ 0001 Creating API...
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/conf
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/controllers
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/tests
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/conf/app.conf
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/main.go
       2019/04/09 15:07:01 INFO     ▶ 0002 Using 'postgres' as 'driver'
       2019/04/09 15:07:01 INFO     ▶ 0003 Using 'postgres://postgres:MyPassDB@127.0.0.1 differences?sslmode=disable' as 'conn'
       2019/04/09 15:07:01 INFO     ▶ 0004 Using '' as 'tables'
       2019/04/09 15:07:01 INFO     ▶ 0005 Analyzing database tables...
       2019/04/09 15:07:01 INFO     ▶ 0006 Creating model files...
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/models usuario.go
       2019/04/09 15:07:01 INFO     ▶ 0007 Creating controller files...
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/controllers usuario.go
       2019/04/09 15:07:01 INFO     ▶ 0008 Creating router files...
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/routers router.go
       2019/04/09 15:07:01 SUCCESS  ▶ 0009 New API successfully created!

      $ tree

       .
       ├── conf
       │   └── app.conf
       ├── controllers
       │   └── usuario.go
       ├── main.go
       ├── models
       │   └── usuario.go
       ├── routers
       │   └── router.go
       └── tests

3. Now create model and controller with scaffold command

       bee generate scaffold usuario2 -fields="id:int, nombre:string, apellido:string" -driver=postgres -conn=postgres://postgres:MyPassDB@127.0.0.1/differences?sslmode=disable

      get

       ______
       | ___ \
       | |_/ /  ___   ___
       | ___ \ / _ \ / _ \
       | |_/ /|  __/|  __/
       \____/  \___| \___| v1.10.0
       2019/04/09 16:04:35 INFO     ▶ 0001 Do you want to create a 'usuario2' model? [Yes|No]
       yes
       2019/04/09 16:04:40 INFO     ▶ 0002 Using 'Usuario2' as model name
       2019/04/09 16:04:40 INFO     ▶ 0003 Using 'models' as package name
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/models/usuario2.go
       2019/04/09 16:04:40 INFO     ▶ 0004 Do you want to create a 'usuario2' controller? [Yes|No]
       yes
       2019/04/09 16:04:50 INFO     ▶ 0005 Using 'Usuario2' as controller name
       2019/04/09 16:04:50 INFO     ▶ 0006 Using 'controllers' as package name
       2019/04/09 16:04:50 INFO     ▶ 0007 Using matching model 'Usuario2'
       	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/controllers/usuario2.go
       2019/04/09 16:04:50 INFO     ▶ 0008 Do you want to create views for this 'usuario2' resource? [Yes|No]
       no
       2019/04/09 16:04:52 INFO     ▶ 0009 Do you want to create a 'usuario2' migration and schema for this resource? [Yes|No]
       no
       2019/04/09 16:04:53 INFO     ▶ 0010 Do you want to migrate the database? [Yes|No]
       no
       2019/04/09 16:04:56 SUCCESS  ▶ 0011 All done! Don't forget to add  beego.Router("/usuario2" ,&controllers.Usuario2Controller{}) to routers/route.go

       2019/04/09 16:04:56 SUCCESS  ▶ 0012 Scaffold successfully generated!

      $ tree

       .
       ├── conf
       │   └── app.conf
       ├── controllers
       │   ├── usuario2.go
       │   └── usuario.go
       ├── img
       │   └── 001.png
       ├── main.go
       ├── models
       │   ├── usuario2.go
       │   └── usuario.go
       ├── README.md
       ├── routers
       │   └── router.go
       └── tests

4. The difference I find:
