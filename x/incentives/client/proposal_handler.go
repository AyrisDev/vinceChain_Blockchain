package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/AyrisDev/vinceChain_Blockchain/x/incentives/client/cli"
	"github.com/AyrisDev/vinceChain_Blockchain/x/incentives/client/rest"
)

var (
	RegisterIncentiveProposalHandler = govclient.NewProposalHandler(cli.NewRegisterIncentiveProposalCmd, rest.RegisterIncentiveProposalRESTHandler)
	CancelIncentiveProposalHandler   = govclient.NewProposalHandler(cli.NewCancelIncentiveProposalCmd, rest.CancelIncentiveProposalRequestRESTHandler)
)
