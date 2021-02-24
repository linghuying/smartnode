package tndao

import (
    "github.com/rocket-pool/rocketpool-go/dao/trustednode"
    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/services"
    "github.com/rocket-pool/smartnode/shared/types/api"
)


func canReplace(c *cli.Context) (*api.CanReplaceTNDAOPositionResponse, error) {

    // Get services
    if err := services.RequireNodeTrusted(c); err != nil { return nil, err }
    w, err := services.GetWallet(c)
    if err != nil { return nil, err }
    rp, err := services.GetRocketPool(c)
    if err != nil { return nil, err }

    // Response
    response := api.CanReplaceTNDAOPositionResponse{}

    // Get node account
    nodeAccount, err := w.GetNodeAccount()
    if err != nil {
        return nil, err
    }

    // Check proposal expired status
    proposalExpired, err := getProposalExpired(rp, nodeAccount.Address, "replace")
    if err != nil {
        return nil, err
    }
    response.ProposalExpired = proposalExpired

    // Update & return response
    response.CanReplace = !response.ProposalExpired
    return &response, nil

}


func replace(c *cli.Context) (*api.ReplaceTNDAOPositionResponse, error) {

    // Get services
    if err := services.RequireNodeTrusted(c); err != nil { return nil, err }
    w, err := services.GetWallet(c)
    if err != nil { return nil, err }
    rp, err := services.GetRocketPool(c)
    if err != nil { return nil, err }

    // Response
    response := api.ReplaceTNDAOPositionResponse{}

    // Get transactor
    opts, err := w.GetNodeAccountTransactor()
    if err != nil {
        return nil, err
    }

    // Replace position
    txReceipt, err := trustednode.Replace(rp, opts)
    if err != nil {
        return nil, err
    }
    response.TxHash = txReceipt.TxHash

    // Return response
    return &response, nil

}
