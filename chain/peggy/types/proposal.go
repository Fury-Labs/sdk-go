package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// constants
const (
	ProposalTypeBlacklistEthereumAddresses string = "ProposalTypeBlacklistEthereumAddresses"
	ProposalTypeRevokeEthereumBlacklist    string = "ProposalTypeRevokeEthereumBlacklist"
)

func init() {
	gov.RegisterProposalType(ProposalTypeBlacklistEthereumAddresses)
	gov.RegisterProposalTypeCodec(&BlacklistEthereumAddressesProposal{}, "kaiju/BlacklistEthereumAddressesProposal")
	gov.RegisterProposalType(ProposalTypeRevokeEthereumBlacklist)
	gov.RegisterProposalTypeCodec(&RevokeEthereumBlacklistProposal{}, "kaiju/RevokeEthereumBlacklistProposal")
}

// Implements Proposal Interface
var _ gov.Content = &BlacklistEthereumAddressesProposal{}
var _ gov.Content = &RevokeEthereumBlacklistProposal{}

// GetTitle returns the title of this proposal.
func (p *BlacklistEthereumAddressesProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *BlacklistEthereumAddressesProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *BlacklistEthereumAddressesProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *BlacklistEthereumAddressesProposal) ProposalType() string {
	return ProposalTypeBlacklistEthereumAddresses
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *BlacklistEthereumAddressesProposal) ValidateBasic() error {
	for _, blacklistAddress := range p.BlacklistAddresses {
		if _, err := NewEthAddress(blacklistAddress); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *RevokeEthereumBlacklistProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *RevokeEthereumBlacklistProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *RevokeEthereumBlacklistProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *RevokeEthereumBlacklistProposal) ProposalType() string {
	return ProposalTypeRevokeEthereumBlacklist
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *RevokeEthereumBlacklistProposal) ValidateBasic() error {
	for _, blacklistAddress := range p.BlacklistAddresses {
		if _, err := NewEthAddress(blacklistAddress); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)
}
