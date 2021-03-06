package main 
import (
   "context"
   "log"

   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
   "github.com/konrawitAEK/app/controllers"
   _ "github.com/konrawitAEK/app/docs"
   "github.com/konrawitAEK/app/ent"
)

type Physicians struct{
    Physician []Physician
}

type Physician struct{
    NAME string
    EMAIL string
}

type Departments struct{
    Department []Department
}

type Department struct{
    Departmentname string
}

type Positions struct{
    Position []Position
}

type Position struct{
    Nameposition string
}


// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host localhost:8080
// @BasePath /api/v1
 
// @securityDefinitions.basic BasicAuth
 
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
 
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
   router := gin.Default()
   router.Use(cors.Default())
 
   client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
   if err != nil {
       log.Fatalf("fail to open sqlite3: %v", err)
   }
   defer client.Close()
 
   if err := client.Schema.Create(context.Background()); err != nil {
       log.Fatalf("failed creating schema resources: %v", err)
   }
 
   v1 := router.Group("/api/v1")
   controllers.NewPositionassingmentController(v1, client)
   controllers.NewPositionController(v1, client)
   controllers.NewPhysicianController(v1, client)
   controllers.NewDepartmentController(v1, client)
   
   physicians := Physicians{
    Physician: []Physician{
                Physician{"Dang Dang","Dang@gmail.com"},
                Physician{"AEK Dang","AEK@gmail.com"},
                Physician{"PANG Dang","PANG@gmail.com"},
                Physician{"NW Dang","NW@gmail.com"},
        },
   }

   for _, p := range physicians.Physician {
        client.Physician.
            Create().
            SetNAME(p.NAME).
            SetEMAIL(p.EMAIL).
            Save(context.Background())
   }

   departments := Departments{
    Department: []Department{
            Department{"ICU"},
            Department{"Radiology Department"},
            Department{"Emergency Room"},
            Department{"Outpatient Department"},
            Department{"Inpatient Department"},
            Department{"Laboratory Department"},
            Department{"Surgical Department"},
        },
   }

   for _, d := range departments.Department {  
        client.Department.
            Create().
            SetDepartmentname(d.Departmentname).
            Save(context.Background())
   }

   positions := Positions{
    Position: []Position{
        Position{"Department head"},
        Position{"member"},
        },
   }

   for _, po := range positions.Position {
        client.Position.
            Create().
            SetNameposition(po.Nameposition).
            Save(context.Background())
   }
   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}
