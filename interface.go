package xcoin_client

import "context"

type IClient interface {
	// GetVersion retrieves the bot version and pair information
	GetVersion(ctx context.Context) (*VersionResponse, error)
	// GetPairList retrieves the list of trading pairs
	GetPairList2(ctx context.Context) ([]Pair, error)
	// GetPairCurrentSettings retrieves current settings for a trading pair
	GetPairCurrentSettings(ctx context.Context, keyWork int) (*KeyValueResponse, error)
	// GetPairOrders retrieves current orders for a trading pair
	GetPairOrders(ctx context.Context, keyWork int) ([]Order, error)
	// AddPair adds a new trading pair
	AddPair(ctx context.Context, val1, val2, birga, mode string) error
	// ExecuteCommand sends a command to a specific trading pair
	ExecuteCommand(ctx context.Context, keyWork int, command CommandType) error
	// UpdatePairSettings updates settings for a trading pair
	UpdatePairSettings(ctx context.Context, keyWork int, settings *KeyValueResponse) error
}
