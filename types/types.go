package types

import (
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

//go:generate cbor-gen-for DealProposalCbor ParamsVersion1

type DealProposalCbor struct {
	PieceCID     cid.Cid
	PieceSize    abi.PaddedPieceSize
	VerifiedDeal bool
	Client       address.Address

	Label DealLabel

	StartEpoch           abi.ChainEpoch
	EndEpoch             abi.ChainEpoch
	StoragePricePerEpoch abi.TokenAmount

	ProviderCollateral abi.TokenAmount
	ClientCollateral   abi.TokenAmount

	Version string
	Params  []byte
}

type ParamsVersion1 struct {
	LocationRef      string
	CarSize          uint64
	SkipIpniAnnounce bool
}

type DealLabel struct {
	bs        []byte
	notString bool
}

func NewLabelFromString(s string) (DealLabel, error) {
	return DealLabel{
		bs:        []byte(s),
		notString: false,
	}, nil
}

func (label *DealLabel) MarshalCBOR(w io.Writer) error {
	scratch := make([]byte, 9)

	// nil *DealLabel counts as EmptyLabel
	// on chain structures should never have a pointer to a DealLabel but the case is included for completeness
	if label == nil {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, 0); err != nil {
			return err
		}
		_, err := io.WriteString(w, string(""))
		return err
	}
	if len(label.bs) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("label is too long to marshal (%d), max allowed (%d)", len(label.bs), cbg.ByteArrayMaxLen)
	}

	majorType := byte(cbg.MajByteString)
	if label.IsString() {
		majorType = cbg.MajTextString
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, majorType, uint64(len(label.bs))); err != nil {
		return err
	}
	_, err := w.Write(label.bs)
	return err
}

func (label *DealLabel) UnmarshalCBOR(br io.Reader) error {
	if label == nil {
		return xerrors.Errorf("cannot unmarshal into nil pointer")
	}

	// reset fields
	label.bs = nil

	scratch := make([]byte, 8)

	maj, length, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajTextString && maj != cbg.MajByteString {
		return fmt.Errorf("unexpected major tag (%d) when unmarshaling DealLabel: only textString (%d) or byteString (%d) expected", maj, cbg.MajTextString, cbg.MajByteString)
	}
	if length > cbg.ByteArrayMaxLen {
		return fmt.Errorf("label was too long (%d), max allowed (%d)", length, cbg.ByteArrayMaxLen)
	}
	buf := make([]byte, length)
	_, err = io.ReadAtLeast(br, buf, int(length))
	if err != nil {
		return err
	}
	label.bs = buf
	label.notString = maj != cbg.MajTextString
	if !label.notString && !utf8.ValidString(string(buf)) {
		return fmt.Errorf("label string not valid utf8")
	}

	return nil
}

func (label DealLabel) IsString() bool {
	return !label.notString
}

func (label DealLabel) IsBytes() bool {
	return label.notString
}

func (label DealLabel) ToString() (string, error) {
	if !label.IsString() {
		return "", xerrors.Errorf("label is not string")
	}

	return string(label.bs), nil
}

func (label DealLabel) ToBytes() ([]byte, error) {
	if !label.IsBytes() {
		return nil, xerrors.Errorf("label is not bytes")
	}
	return label.bs, nil
}

func (label DealLabel) Length() int {
	return len(label.bs)
}

func (l DealLabel) Equals(o DealLabel) bool {
	return bytes.Equal(l.bs, o.bs) && l.notString == o.notString
}
