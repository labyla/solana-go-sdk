package token2022

import (
	"encoding/binary"

	"github.com/labyla/solana-go-sdk/common"
)

// ExtensionType represents Token-2022 extension types
type ExtensionType uint16

const (
	ExtensionTypeUninitialized                 ExtensionType = 0
	ExtensionTypeTransferFeeConfig             ExtensionType = 1
	ExtensionTypeTransferFeeAmount             ExtensionType = 2
	ExtensionTypeMintCloseAuthority            ExtensionType = 3
	ExtensionTypeConfidentialTransferMint      ExtensionType = 4
	ExtensionTypeConfidentialTransferAccount   ExtensionType = 5
	ExtensionTypeDefaultAccountState           ExtensionType = 6
	ExtensionTypeImmutableOwner                ExtensionType = 7
	ExtensionTypeMemoTransfer                  ExtensionType = 8
	ExtensionTypeNonTransferable               ExtensionType = 9
	ExtensionTypeInterestBearingConfig         ExtensionType = 10
	ExtensionTypeCpiGuard                      ExtensionType = 11
	ExtensionTypePermanentDelegate             ExtensionType = 12
	ExtensionTypeNonTransferableAccount        ExtensionType = 13
	ExtensionTypeTransferHook                  ExtensionType = 14
	ExtensionTypeTransferHookAccount           ExtensionType = 15
	ExtensionTypeConfidentialTransferFeeConfig ExtensionType = 16
	ExtensionTypeConfidentialTransferFeeAmount ExtensionType = 17
	ExtensionTypeMetadataPointer               ExtensionType = 18
	ExtensionTypeTokenMetadata                 ExtensionType = 19
	ExtensionTypeGroupPointer                  ExtensionType = 20
	ExtensionTypeTokenGroup                    ExtensionType = 21
	ExtensionTypeGroupMemberPointer            ExtensionType = 22
	ExtensionTypeTokenGroupMember              ExtensionType = 23
	ExtensionTypeConfidentialMintBurn          ExtensionType = 24
	ExtensionTypeScaledUiAmount                ExtensionType = 25
	ExtensionTypePausable                      ExtensionType = 26
	ExtensionTypePausableAccount               ExtensionType = 27
)

// AccountType represents the type of Token-2022 account
type AccountType uint8

const (
	AccountTypeUninitialized AccountType = 0
	AccountTypeMint          AccountType = 1
	AccountTypeAccount       AccountType = 2
)

// BaseAccountLength is the length after which extensions start
const BaseAccountLength = 165

// Extension sizes
const (
	MintCloseAuthoritySize     = 32
	TransferFeeConfigSize      = 108
	TransferFeeAmountSize      = 8
	DefaultAccountStateSize    = 1
	ImmutableOwnerSize         = 0
	MemoTransferSize           = 1
	NonTransferableSize        = 0
	NonTransferableAccountSize = 0
	InterestBearingConfigSize  = 52
	CpiGuardSize               = 1
	PermanentDelegateSize      = 32
	TransferHookSize           = 64
	TransferHookAccountSize    = 1
	MetadataPointerSize        = 64
	GroupPointerSize           = 64
	GroupMemberPointerSize     = 64
	TokenGroupSize             = 72
	TokenGroupMemberSize       = 72
	ScaledUiAmountConfigSize   = 26
	PausableConfigSize         = 1
	PausableAccountSize        = 0
)

// MintCloseAuthority extension - allows closing a mint account
type MintCloseAuthority struct {
	CloseAuthority common.PublicKey
}

// TransferFee represents a transfer fee configuration
type TransferFee struct {
	Epoch                  uint64
	MaximumFee             uint64
	TransferFeeBasisPoints uint16
}

// TransferFeeConfig extension - enables transfer fees on the mint
type TransferFeeConfig struct {
	TransferFeeConfigAuthority common.PublicKey
	WithdrawWithheldAuthority  common.PublicKey
	WithheldAmount             uint64
	OlderTransferFee           TransferFee
	NewerTransferFee           TransferFee
}

// TransferFeeAmount extension - holds withheld transfer fees on token accounts
type TransferFeeAmount struct {
	WithheldAmount uint64
}

// DefaultAccountState extension - sets default state for new token accounts
type DefaultAccountState struct {
	State uint8 // 0 = Uninitialized, 1 = Initialized, 2 = Frozen
}

// ImmutableOwner extension - prevents ownership transfer of token accounts
type ImmutableOwner struct{}

// MemoTransfer extension - requires memo on incoming transfers
type MemoTransfer struct {
	RequireIncomingTransferMemos bool
}

// NonTransferable extension - marks mint as non-transferable (soul-bound)
type NonTransferable struct{}

// NonTransferableAccount extension - marks account as belonging to non-transferable mint
type NonTransferableAccount struct{}

// InterestBearingConfig extension - enables interest accrual on tokens
type InterestBearingConfig struct {
	RateAuthority           common.PublicKey
	InitializationTimestamp int64
	PreUpdateAverageRate    int16
	LastUpdateTimestamp     int64
	CurrentRate             int16
}

// CpiGuard extension - restricts certain CPI operations
type CpiGuard struct {
	LockCpi bool
}

// PermanentDelegate extension - permanent delegate authority on all accounts
type PermanentDelegate struct {
	Delegate common.PublicKey
}

// TransferHook extension - enables custom transfer hook program
type TransferHook struct {
	Authority common.PublicKey
	ProgramId common.PublicKey
}

// TransferHookAccount extension - indicates account belongs to mint with transfer hook
type TransferHookAccount struct {
	Transferring bool
}

// MetadataPointer extension - points to metadata account
type MetadataPointer struct {
	Authority       common.PublicKey
	MetadataAddress common.PublicKey
}

// GroupPointer extension - points to group configuration
type GroupPointer struct {
	Authority    common.PublicKey
	GroupAddress common.PublicKey
}

// GroupMemberPointer extension - points to group member configuration
type GroupMemberPointer struct {
	Authority     common.PublicKey
	MemberAddress common.PublicKey
}

// TokenGroup extension - group configuration stored in mint
type TokenGroup struct {
	UpdateAuthority common.PublicKey
	Mint            common.PublicKey
	Size            uint32
	MaxSize         uint32
}

// TokenGroupMember extension - group member configuration stored in mint
type TokenGroupMember struct {
	Mint         common.PublicKey
	Group        common.PublicKey
	MemberNumber uint32
}

// ScaledUiAmountConfig extension - UI amount scaling
type ScaledUiAmountConfig struct {
	Authority                       common.PublicKey
	Multiplier                      float64 // stored as f64
	NewMultiplierEffectiveTimestamp int64
	NewMultiplier                   float64
}

// PausableConfig extension - allows pausing mint operations
type PausableConfig struct {
	Paused bool
}

// PausableAccount extension - indicates account belongs to pausable mint
type PausableAccount struct{}

// TokenMetadata extension - variable length metadata
type TokenMetadata struct {
	UpdateAuthority    common.PublicKey
	Mint               common.PublicKey
	Name               string
	Symbol             string
	Uri                string
	AdditionalMetadata []struct {
		Key   string
		Value string
	}
}

// MintExtensions holds all possible mint extensions
type MintExtensions struct {
	MintCloseAuthority    *MintCloseAuthority
	TransferFeeConfig     *TransferFeeConfig
	DefaultAccountState   *DefaultAccountState
	NonTransferable       *NonTransferable
	InterestBearingConfig *InterestBearingConfig
	PermanentDelegate     *PermanentDelegate
	TransferHook          *TransferHook
	MetadataPointer       *MetadataPointer
	TokenMetadata         *TokenMetadata
	GroupPointer          *GroupPointer
	TokenGroup            *TokenGroup
	GroupMemberPointer    *GroupMemberPointer
	TokenGroupMember      *TokenGroupMember
	ScaledUiAmountConfig  *ScaledUiAmountConfig
	PausableConfig        *PausableConfig
}

// AccountExtensions holds all possible token account extensions
type AccountExtensions struct {
	TransferFeeAmount      *TransferFeeAmount
	ImmutableOwner         *ImmutableOwner
	MemoTransfer           *MemoTransfer
	NonTransferableAccount *NonTransferableAccount
	CpiGuard               *CpiGuard
	TransferHookAccount    *TransferHookAccount
	PausableAccount        *PausableAccount
}

// ParseMintExtensions parses extensions from mint account data
func ParseMintExtensions(data []byte) (*MintExtensions, error) {
	if len(data) <= BaseAccountLength {
		return nil, nil
	}

	// Account type is at BaseAccountLength
	if len(data) <= BaseAccountLength+1 {
		return nil, nil
	}

	accountType := AccountType(data[BaseAccountLength])
	if accountType != AccountTypeMint {
		return nil, ErrInvalidAccountDataSize
	}

	extensions := &MintExtensions{}
	tlvData := data[BaseAccountLength+1:]

	if err := parseTLVExtensions(tlvData, func(extType ExtensionType, extData []byte) error {
		return parseMintExtension(extensions, extType, extData)
	}); err != nil {
		return nil, err
	}

	return extensions, nil
}

// ParseAccountExtensions parses extensions from token account data
func ParseAccountExtensions(data []byte) (*AccountExtensions, error) {
	if len(data) <= BaseAccountLength {
		return nil, nil
	}

	// Account type is at BaseAccountLength
	if len(data) <= BaseAccountLength+1 {
		return nil, nil
	}

	accountType := AccountType(data[BaseAccountLength])
	if accountType != AccountTypeAccount {
		return nil, ErrInvalidAccountDataSize
	}

	extensions := &AccountExtensions{}
	tlvData := data[BaseAccountLength+1:]

	if err := parseTLVExtensions(tlvData, func(extType ExtensionType, extData []byte) error {
		return parseAccountExtension(extensions, extType, extData)
	}); err != nil {
		return nil, err
	}

	return extensions, nil
}

// parseTLVExtensions iterates through TLV-encoded extension data
func parseTLVExtensions(data []byte, handler func(ExtensionType, []byte) error) error {
	offset := 0
	for offset+4 <= len(data) {
		extType := ExtensionType(binary.LittleEndian.Uint16(data[offset : offset+2]))
		if extType == ExtensionTypeUninitialized {
			break
		}

		length := int(binary.LittleEndian.Uint16(data[offset+2 : offset+4]))
		offset += 4

		if offset+length > len(data) {
			return ErrInvalidAccountDataSize
		}

		extData := data[offset : offset+length]
		if err := handler(extType, extData); err != nil {
			return err
		}

		offset += length
	}
	return nil
}

func parseMintExtension(ext *MintExtensions, extType ExtensionType, data []byte) error {
	switch extType {
	case ExtensionTypeMintCloseAuthority:
		if len(data) >= 32 {
			ext.MintCloseAuthority = &MintCloseAuthority{
				CloseAuthority: common.PublicKeyFromBytes(data[:32]),
			}
		}
	case ExtensionTypeTransferFeeConfig:
		if len(data) >= TransferFeeConfigSize {
			ext.TransferFeeConfig = &TransferFeeConfig{
				TransferFeeConfigAuthority: common.PublicKeyFromBytes(data[0:32]),
				WithdrawWithheldAuthority:  common.PublicKeyFromBytes(data[32:64]),
				WithheldAmount:             binary.LittleEndian.Uint64(data[64:72]),
				OlderTransferFee: TransferFee{
					Epoch:                  binary.LittleEndian.Uint64(data[72:80]),
					MaximumFee:             binary.LittleEndian.Uint64(data[80:88]),
					TransferFeeBasisPoints: binary.LittleEndian.Uint16(data[88:90]),
				},
				NewerTransferFee: TransferFee{
					Epoch:                  binary.LittleEndian.Uint64(data[90:98]),
					MaximumFee:             binary.LittleEndian.Uint64(data[98:106]),
					TransferFeeBasisPoints: binary.LittleEndian.Uint16(data[106:108]),
				},
			}
		}
	case ExtensionTypeDefaultAccountState:
		if len(data) >= 1 {
			ext.DefaultAccountState = &DefaultAccountState{
				State: data[0],
			}
		}
	case ExtensionTypeNonTransferable:
		ext.NonTransferable = &NonTransferable{}
	case ExtensionTypeInterestBearingConfig:
		if len(data) >= InterestBearingConfigSize {
			ext.InterestBearingConfig = &InterestBearingConfig{
				RateAuthority:           common.PublicKeyFromBytes(data[0:32]),
				InitializationTimestamp: int64(binary.LittleEndian.Uint64(data[32:40])),
				PreUpdateAverageRate:    int16(binary.LittleEndian.Uint16(data[40:42])),
				LastUpdateTimestamp:     int64(binary.LittleEndian.Uint64(data[42:50])),
				CurrentRate:             int16(binary.LittleEndian.Uint16(data[50:52])),
			}
		}
	case ExtensionTypePermanentDelegate:
		if len(data) >= 32 {
			ext.PermanentDelegate = &PermanentDelegate{
				Delegate: common.PublicKeyFromBytes(data[:32]),
			}
		}
	case ExtensionTypeTransferHook:
		if len(data) >= TransferHookSize {
			ext.TransferHook = &TransferHook{
				Authority: common.PublicKeyFromBytes(data[0:32]),
				ProgramId: common.PublicKeyFromBytes(data[32:64]),
			}
		}
	case ExtensionTypeMetadataPointer:
		if len(data) >= MetadataPointerSize {
			ext.MetadataPointer = &MetadataPointer{
				Authority:       common.PublicKeyFromBytes(data[0:32]),
				MetadataAddress: common.PublicKeyFromBytes(data[32:64]),
			}
		}
	case ExtensionTypeGroupPointer:
		if len(data) >= GroupPointerSize {
			ext.GroupPointer = &GroupPointer{
				Authority:    common.PublicKeyFromBytes(data[0:32]),
				GroupAddress: common.PublicKeyFromBytes(data[32:64]),
			}
		}
	case ExtensionTypeGroupMemberPointer:
		if len(data) >= GroupMemberPointerSize {
			ext.GroupMemberPointer = &GroupMemberPointer{
				Authority:     common.PublicKeyFromBytes(data[0:32]),
				MemberAddress: common.PublicKeyFromBytes(data[32:64]),
			}
		}
	case ExtensionTypeTokenGroup:
		if len(data) >= TokenGroupSize {
			ext.TokenGroup = &TokenGroup{
				UpdateAuthority: common.PublicKeyFromBytes(data[0:32]),
				Mint:            common.PublicKeyFromBytes(data[32:64]),
				Size:            binary.LittleEndian.Uint32(data[64:68]),
				MaxSize:         binary.LittleEndian.Uint32(data[68:72]),
			}
		}
	case ExtensionTypeTokenGroupMember:
		if len(data) >= TokenGroupMemberSize {
			ext.TokenGroupMember = &TokenGroupMember{
				Mint:         common.PublicKeyFromBytes(data[0:32]),
				Group:        common.PublicKeyFromBytes(data[32:64]),
				MemberNumber: binary.LittleEndian.Uint32(data[64:68]),
			}
		}
	case ExtensionTypePausable:
		if len(data) >= 1 {
			ext.PausableConfig = &PausableConfig{
				Paused: data[0] == 1,
			}
		}
	case ExtensionTypeTokenMetadata:
		// Variable length - parse carefully
		if metadata, err := parseTokenMetadata(data); err == nil {
			ext.TokenMetadata = metadata
		}
	}
	return nil
}

func parseAccountExtension(ext *AccountExtensions, extType ExtensionType, data []byte) error {
	switch extType {
	case ExtensionTypeTransferFeeAmount:
		if len(data) >= 8 {
			ext.TransferFeeAmount = &TransferFeeAmount{
				WithheldAmount: binary.LittleEndian.Uint64(data[0:8]),
			}
		}
	case ExtensionTypeImmutableOwner:
		ext.ImmutableOwner = &ImmutableOwner{}
	case ExtensionTypeMemoTransfer:
		if len(data) >= 1 {
			ext.MemoTransfer = &MemoTransfer{
				RequireIncomingTransferMemos: data[0] == 1,
			}
		}
	case ExtensionTypeNonTransferableAccount:
		ext.NonTransferableAccount = &NonTransferableAccount{}
	case ExtensionTypeCpiGuard:
		if len(data) >= 1 {
			ext.CpiGuard = &CpiGuard{
				LockCpi: data[0] == 1,
			}
		}
	case ExtensionTypeTransferHookAccount:
		if len(data) >= 1 {
			ext.TransferHookAccount = &TransferHookAccount{
				Transferring: data[0] == 1,
			}
		}
	case ExtensionTypePausableAccount:
		ext.PausableAccount = &PausableAccount{}
	}
	return nil
}

// parseTokenMetadata parses variable-length token metadata
func parseTokenMetadata(data []byte) (*TokenMetadata, error) {
	if len(data) < 68 { // minimum: 2 pubkeys + 4 length bytes
		return nil, ErrInvalidAccountDataSize
	}

	metadata := &TokenMetadata{
		UpdateAuthority: common.PublicKeyFromBytes(data[0:32]),
		Mint:            common.PublicKeyFromBytes(data[32:64]),
	}

	offset := 64

	// Parse name
	if offset+4 > len(data) {
		return metadata, nil
	}
	nameLen := int(binary.LittleEndian.Uint32(data[offset : offset+4]))
	offset += 4
	if offset+nameLen > len(data) {
		return metadata, nil
	}
	metadata.Name = string(data[offset : offset+nameLen])
	offset += nameLen

	// Parse symbol
	if offset+4 > len(data) {
		return metadata, nil
	}
	symbolLen := int(binary.LittleEndian.Uint32(data[offset : offset+4]))
	offset += 4
	if offset+symbolLen > len(data) {
		return metadata, nil
	}
	metadata.Symbol = string(data[offset : offset+symbolLen])
	offset += symbolLen

	// Parse URI
	if offset+4 > len(data) {
		return metadata, nil
	}
	uriLen := int(binary.LittleEndian.Uint32(data[offset : offset+4]))
	offset += 4
	if offset+uriLen > len(data) {
		return metadata, nil
	}
	metadata.Uri = string(data[offset : offset+uriLen])
	offset += uriLen

	// Parse additional metadata (key-value pairs)
	if offset+4 > len(data) {
		return metadata, nil
	}
	numPairs := int(binary.LittleEndian.Uint32(data[offset : offset+4]))
	offset += 4

	for i := 0; i < numPairs && offset < len(data); i++ {
		// Parse key
		if offset+4 > len(data) {
			break
		}
		keyLen := int(binary.LittleEndian.Uint32(data[offset : offset+4]))
		offset += 4
		if offset+keyLen > len(data) {
			break
		}
		key := string(data[offset : offset+keyLen])
		offset += keyLen

		// Parse value
		if offset+4 > len(data) {
			break
		}
		valueLen := int(binary.LittleEndian.Uint32(data[offset : offset+4]))
		offset += 4
		if offset+valueLen > len(data) {
			break
		}
		value := string(data[offset : offset+valueLen])
		offset += valueLen

		metadata.AdditionalMetadata = append(metadata.AdditionalMetadata, struct {
			Key   string
			Value string
		}{Key: key, Value: value})
	}

	return metadata, nil
}

// HasExtensions checks if the data contains Token-2022 extensions
func HasExtensions(data []byte) bool {
	return len(data) > BaseAccountLength
}

// GetExtensionTypes returns a list of extension types present in the data
func GetExtensionTypes(data []byte) ([]ExtensionType, error) {
	if len(data) <= BaseAccountLength+1 {
		return nil, nil
	}

	tlvData := data[BaseAccountLength+1:]
	var types []ExtensionType

	if err := parseTLVExtensions(tlvData, func(extType ExtensionType, _ []byte) error {
		types = append(types, extType)
		return nil
	}); err != nil {
		return nil, err
	}

	return types, nil
}
