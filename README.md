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
        2019/04/09 15:07:01 INFO     ▶ 0003 Using 'postgres://postgres:MyPassDB@127.0.0.1/differences?sslmode=disable' as 'conn'
        2019/04/09 15:07:01 INFO     ▶ 0004 Using '' as 'tables'
        2019/04/09 15:07:01 INFO     ▶ 0005 Analyzing database tables...
        2019/04/09 15:07:01 INFO     ▶ 0006 Creating model files...
        	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/models/usuario.go
        2019/04/09 15:07:01 INFO     ▶ 0007 Creating controller files...
        	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/controllers/usuario.go
        2019/04/09 15:07:01 INFO     ▶ 0008 Creating router files...
        	create	 /home/jjvargass/go/src/github.com/jotavargas/differences_beego/routers/router.go
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
