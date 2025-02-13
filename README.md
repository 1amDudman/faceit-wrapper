# Faceit SDK

The Faceit SDK provides a user-friendly interface for interacting with the Faceit API, including working with players and matches. The SDK is built with an emphasis on ease of use, flexibility, and scalability, allowing developers to integrate Faceit functionality into their applications.
**Mainly focused on CS2 and different player statistics.**

## Possibilities

- **Players**: Getting information about players.
- **Matches**: Access match data.
- **Flexibility**: Support for custom services via interfaces.
- **Future Development:** tests, webhooks, demo parsing, and other Faceit API endpoints related to a player statistics.

## Install

`go get github.com/1amDudman/faceit-wrapper`

## Requirements
- **Go 1.18+**
- **Active Faceit API Key**

## Examples of usage

**Client with default configs**
```go
faceitClient, err := client.NewFaceitClient("api-url", "api-key")
if err != nil {
        log.Fatalf("FaceitClient initialization error: %v", err)
    }
log.Println("Client was successufully created!")
```

**Client with custom configs**
```go
faceitClient, err := client.NewFaceitClient(
    "api-url",
    "api-key",
    client.WithTimeout(5 * time.Second),
    client.WithCustomPlayerService(someCustomPlayerService), // should satisfy interfaces.PlayerService
	client.WithCustomMatchService(someCustomMatchService), // should satisfy interfaces.MatchService
)
if err != nil {
        log.Fatalf("FaceitClient initialization error: %v", err)
    }
log.Println("Client was successufully created!")
```

**Services**
```go
player, err := faceitClient.PlayerService.GetByPlayerID(ctx, "playerID")
	if err != nil {
		log.Fatalf("Failed to fetch player: %v", err)
	}

match, err := faceitClient.MatchService.GetByMatchID(ctx, "matchID")
	if err != nil {
		log.Fatalf("Failed to fetch match: %v", err)
	}

fmt.Printf("Person: %+v\n", player)
fmt.Printf("Match: %+v\n", match)
```

## Design principles
- **Initialization:** Simple client initialization with API URL and API key.
- **Configuration:** Customizable via options, support for custom services.
- **Usage:** Logical and structured API for working with basic entities and services.

## License
This project is licensed under the MIT License. See the LICENSE file for more information.
