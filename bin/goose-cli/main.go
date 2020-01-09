package main

import (
	"fmt"
	"log"
	"os"

  "net/http"
	"github.com/urfave/cli/v2"
)

var goose_ascii = 
`
        Honk Honk                  ___
                               ,-""   '.
                             ,'  _   e )'-._
                            /  ,' '-._<.===-'
                           /  /
                          /  ;
              _          /   ;
 ('._    _.-"" ""--..__,'    |
 <_  '-""                     \
  <'-                          :
   (__   <__.                  ;
     '-.   '-.__.      _.'    /
        \      '-.__,-'    _,'
         '._    ,    /__,-'
            ""._\__,'< <____
                 | |  '----.'.
                 | |        \ '.
                 ; |___      \-''
                 \   --<
                  '.'.<
                    '-'
`

func main() {
	app := &cli.App{
		Name: "Goose CLI",
		Usage: "Goose App Controller",
		Action: func(c *cli.Context) error {
			fmt.Println(goose_ascii)
			return nil
    },

		Commands: []*cli.Command{
			{
        Name: "add",
        Aliases: []string{"a"},
        Usage: "Add a goose to flock",
        Action: func(c *cli.Context) error {
          resp, err := http.Get("http://127.0.0.1:8080/add")
          if (err != nil) {
            log.Fatalln(err)
          }
          defer resp.Body.Close()
          fmt.Println("Goose Added")
          return nil
        },
      },
      {
        Name: "remove",
        Aliases: []string{"r"},
        Usage: "Kick a goose from flock",
        Action: func(c *cli.Context) error {
          resp, err := http.Get("http://127.0.0.1:8080/remove")
          if (err != nil) {
            log.Fatalln(err)
          }
          defer resp.Body.Close()
          fmt.Println("Goose Removed")
          return nil
        },
      },
		},
	}

	err := app.Run(os.Args)
		if err != nil {
			log.Fatal(err)
		}
}
