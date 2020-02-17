package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

func main() {
	// okenString := "eyJhbGciOiJSUzI1NiIsImtpZCI6Ijc2MmZhNjM3YWY5NTM1OTBkYjhiYjhhNjM2YmYxMWQ0MzYwYWJjOTgiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJ2ZXJpZnktand0LXB1YnN1YiIsImF6cCI6IjExODA3NTM5NTgzOTM2MjY2NDE0OSIsImVtYWlsIjoicHVic3ViLXZlcmlmaWNhdGlvbi1leGFtcGxlQHBva3V0dW5hLWRldi5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJleHAiOjE1ODE4NDY0NDAsImlhdCI6MTU4MTg0Mjg0MCwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tIiwic3ViIjoiMTE4MDc1Mzk1ODM5MzYyNjY0MTQ5In0.tNUAXJbbjxo5XlgL_ipwzm10b_e9HvabRtKvtLS6L6TR47ciJmm2i3Eh2ixAax9qp1GTLPUzdHKcxsnYiS1tpW-PQw7zMKwl5VmKQ9BXCOf9uSiloNw5DhukY5714umSyeA1D1EUEc7IhjEUnUzyVdCNMBwsFE7B2BGRfz7eoBmlDRcFVUmyDg0TlCDyb0GMZ-jaqxYt4txpU3DJueRsOopMdE0PTWEkmIZK-7Oi1QTOO6bviWTOpcU9pQw2Daf5NXHM-mGvEG7uISouC3QDO3ip6KpZi550lx7Ym4XZ3A5yvaAZlST8NEIh-qGG6Xie_m4I0aS6t_-VEVb5UqFWpg"
	tokenString := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjdkNjgwZDhjNzBkNDRlOTQ3MTMzY2JkNDk5ZWJjMWE2MWMzZDVhYmMiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJodHRwczovL2V4YW1wbGUuY29tIiwiYXpwIjoiMTEzNzc0MjY0NDYzMDM4MzIxOTY0IiwiZW1haWwiOiJnYWUtZ2NwQGFwcHNwb3QuZ3NlcnZpY2VhY2NvdW50LmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJleHAiOjE1NTAxODU5MzUsImlhdCI6MTU1MDE4MjMzNSwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tIiwic3ViIjoiMTEzNzc0MjY0NDYzMDM4MzIxOTY0In0.QVjyqpmadTyDZmlX2u3jWd1kJ68YkdwsRZDo-QxSPbxjug4ucLBwAs2QePrcgZ6hhkvdc4UHY4YF3fz9g7XHULNVIzX5xh02qXEH8dK6PgGndIWcZQzjSYfgO-q-R2oo2hNM5HBBsQN4ARtGK_acG-NGGWM3CQfahbEjZPAJe_B8M7HfIu_G5jOLZCw2EUcGo8BvEwGcLWB2WqEgRM0-xt5-UPzoa3-FpSPG7DHk7z9zRUeq6eB__ldb-2o4RciJmjVwHgnYqn3VvlX9oVKEgXpNFhKuYA-mWh5o7BCwhujSMmFoBOh6mbIXFcyf5UiVqKjpqEbqPGo_AvKvIQ9VTQ"
	token, err := jwt.ParseString(tokenString)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Printf("%+v\n", token)

	message, err := jws.ParseString(tokenString)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Printf("%+v\n", message.Signatures()[0].ProtectedHeaders())

	// fetch & serialize
	set, err := jwk.Fetch("https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	fmt.Printf("%+v\n", set)
	json, err := json.Marshal(set.Keys[1])
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Printf("%+v\n", string(json))

	// verify
	payload, err := jws.VerifyWithJWKSet([]byte(tokenString), set, nil)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Printf("%+v\n", string(payload))

}
