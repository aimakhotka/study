package app

import (
	"yana_project/internal/services/xmlparse"
)

func Run() {
	//db, err := repository.GetPgConnection(repository.Config{
	//	Host:     os.Getenv("POSTGRES_HOST"),
	//	Port:     os.Getenv("POSTGRES_PORT"),
	//	Username: os.Getenv("POSTGRES_USER"),
	//	DBName:   os.Getenv("POSTGRES_DB"),
	//	SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	//	Password: os.Getenv("POSTGRES_PASSWORD"),
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	xmlparse.XmlParsing()
	//repos := repository.NewRepository(db)
	//port := ":8081"
	//fmt.Sprintf("Server is listening on port %s...", port)
	//http.ListenAndServe(port, http.FileServer(http.Dir("web")))

}
