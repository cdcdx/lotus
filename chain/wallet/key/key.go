package key

import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/filecoin-project/lotus/lib/sigs"
)

func GenerateKey(typ types.KeyType) (*Key, error) {
	ctyp := ActSigType(typ)
	if ctyp == crypto.SigTypeUnknown {
		return nil, xerrors.Errorf("unknown sig type: %s", typ)
	}
	pk, err := sigs.Generate(ctyp)
	if err != nil {
		return nil, err
	}

	// wallet-security 加密
	pk1, err := MakeByte(pk)
	if err != nil {
		return nil, err
	}

	ki := types.KeyInfo{
		Type:       typ,
		PrivateKey: pk1, //PrivateKey: pk,
	}
	return NewKey(ki)
}

type Key struct {
	types.KeyInfo

	PublicKey []byte
	Address   address.Address
}

func NewKey(keyinfo types.KeyInfo) (*Key, error) {
	k := &Key{
		KeyInfo: keyinfo,
	}

	// wallet-security 解密
	// var err error
	// k.PublicKey, err = sigs.ToPublic(ActSigType(k.Type), k.PrivateKey)
	pk, err := UnMakeByte(k.PrivateKey)
	if err != nil {
		return nil, err
	}
	k.PublicKey, err = sigs.ToPublic(ActSigType(k.Type), pk)

	if err != nil {
		return nil, err
	}

	switch k.Type {
	case types.KTSecp256k1:
		k.Address, err = address.NewSecp256k1Address(k.PublicKey)
		if err != nil {
			return nil, xerrors.Errorf("converting Secp256k1 to address: %w", err)
		}
	case types.KTDelegated:
		// Transitory Delegated signature verification as per FIP-0055
		ethAddr, err := ethtypes.EthAddressFromPubKey(k.PublicKey)
		if err != nil {
			return nil, xerrors.Errorf("failed to calculate Eth address from public key: %w", err)
		}

		ea, err := ethtypes.CastEthAddress(ethAddr)
		if err != nil {
			return nil, xerrors.Errorf("failed to create ethereum address from bytes: %w", err)
		}

		k.Address, err = ea.ToFilecoinAddress()
		if err != nil {
			return nil, xerrors.Errorf("converting Delegated to address: %w", err)
		}
	case types.KTBLS:
		k.Address, err = address.NewBLSAddress(k.PublicKey)
		if err != nil {
			return nil, xerrors.Errorf("converting BLS to address: %w", err)
		}
	default:
		return nil, xerrors.Errorf("unsupported key type: %s", k.Type)
	}

	return k, nil

}

func ActSigType(typ types.KeyType) crypto.SigType {
	switch typ {
	case types.KTBLS:
		return crypto.SigTypeBLS
	case types.KTSecp256k1:
		return crypto.SigTypeSecp256k1
	case types.KTDelegated:
		return crypto.SigTypeDelegated
	default:
		return crypto.SigTypeUnknown
	}
}
