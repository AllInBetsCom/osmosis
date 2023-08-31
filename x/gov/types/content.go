package types

import (
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// Constants pertaining to a Content object
const (
	MaxDescriptionLength int = 10000
	MaxTitleLength       int = 140
)

// Content defines an interface that a proposal must implement. It contains
// information such as the title and description along with the type and routing
// information for the appropriate handler to process the proposal. Content can
// have additional fields, which will handled by a proposal's Handler.
// TODO Try to unify this interface with types/module/simulation
// https://github.com/cosmos/cosmos-sdk/issues/5853
type Content govtypes.Content

// Handler defines a function that handles a proposal after it has passed the
// governance process.
type Handler govtypes.Handler

// ValidateAbstract validates a proposal's abstract contents returning an error
// if invalid.
func ValidateAbstract(c Content) error {
	title := c.GetTitle()
	if len(strings.TrimSpace(title)) == 0 {
		return sdkerrors.Wrap(ErrInvalidProposalContent, "proposal title cannot be blank")
	}
	if len(title) > MaxTitleLength {
		return sdkerrors.Wrapf(ErrInvalidProposalContent, "proposal title is longer than max length of %d", MaxTitleLength)
	}

	description := c.GetDescription()
	if len(description) == 0 {
		return sdkerrors.Wrap(ErrInvalidProposalContent, "proposal description cannot be blank")
	}
	if len(description) > MaxDescriptionLength {
		return sdkerrors.Wrapf(ErrInvalidProposalContent, "proposal description is longer than max length of %d", MaxDescriptionLength)
	}

	return nil
}
