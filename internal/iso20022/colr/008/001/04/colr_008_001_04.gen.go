// Code generated by main. DO NOT EDIT.

package colr_008_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Prtry"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CashCollateralResponse1 struct {
	RspnTp    Status4Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 RspnTp"`
	CollId    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CollId,omitempty"`
	AsstNb    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 AsstNb,omitempty"`
	CshAcctId AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CshAcctId,omitempty"`
	RjctnRsn  RejectionReasonV021Code      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 RjctnRsn,omitempty"`
	RjctnInf  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 RjctnInf,omitempty"`
}

type CollateralAccount2 struct {
	Id Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Id"`
	Tp CollateralAccountIdentificationType2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Tp,omitempty"`
	Nm Max70Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Nm,omitempty"`
}

type CollateralAccountIdentificationType2Choice struct {
	Tp    CollateralAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Tp,omitempty"`
	Prtry GenericIdentification36    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Prtry"`
}

// May be one of HOUS, CLIE, LIPR, MGIN, DFLT
type CollateralAccountType1Code string

// May be one of INPR, COPR
type CollateralProposalResponse1Code string

type CollateralProposalResponse2 struct {
	VartnMrgn       CollateralProposalResponseType2 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 VartnMrgn"`
	SgrtdIndpdntAmt CollateralProposalResponseType2 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 SgrtdIndpdntAmt,omitempty"`
}

type CollateralProposalResponse2Choice struct {
	CollPrpsl       CollateralProposalResponse2     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CollPrpsl"`
	SgrtdIndpdntAmt CollateralProposalResponseType2 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 SgrtdIndpdntAmt"`
}

type CollateralProposalResponseType2 struct {
	CollPrpslId Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CollPrpslId"`
	Tp          CollateralProposalResponse1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Tp"`
	Rspn        CollateralResponse1             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Rspn"`
}

type CollateralProposalResponseV04 struct {
	TxId        Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 TxId"`
	Oblgtn      Obligation4                       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Oblgtn"`
	PrpslRspn   CollateralProposalResponse2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 PrpslRspn"`
	SplmtryData []SupplementaryData1              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 SplmtryData,omitempty"`
}

type CollateralResponse1 struct {
	SctiesCollRspn []SecuritiesCollateralResponse1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 SctiesCollRspn,omitempty"`
	CshCollRspn    []CashCollateralResponse1       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CshCollRspn,omitempty"`
	OthrCollRspn   []OtherCollateralResponse1      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 OthrCollRspn,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 DtTm"`
}

type Document struct {
	CollPrpslRspn CollateralProposalResponseV04 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CollPrpslRspn"`
}

// May be one of BFWD, PAYM, CCPC, COMM, CRDS, CRTL, CRSP, CCIR, CRPR, EQUI, EQPT, EQUS, EXTD, EXPT, FIXI, FORX, FORW, FUTR, OPTN, LIQU, OTCD, REPO, RVPO, SLOA, SBSC, SCRP, SLEB, SHSL, SCIR, SCIE, SWPT, TBAS, TRBD, TRCP
type ExposureType5Code string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Issr,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Adr"`
}

type Obligation4 struct {
	PtyA       PartyIdentification100Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 PtyA"`
	SvcgPtyA   PartyIdentification100Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 SvcgPtyA,omitempty"`
	PtyB       PartyIdentification100Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 PtyB"`
	SvcgPtyB   PartyIdentification100Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 SvcgPtyB,omitempty"`
	CollAcctId CollateralAccount2           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CollAcctId,omitempty"`
	XpsrTp     ExposureType5Code            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 XpsrTp,omitempty"`
	ValtnDt    DateAndDateTimeChoice        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 ValtnDt"`
}

type OtherCollateralResponse1 struct {
	RspnTp   Status4Code             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 RspnTp"`
	CollId   Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CollId,omitempty"`
	AsstNb   Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 AsstNb,omitempty"`
	RjctnRsn RejectionReasonV021Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 RjctnRsn,omitempty"`
	RjctnInf Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 RjctnInf,omitempty"`
}

type PartyIdentification100Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 PrtryId"`
	NmAndAdr NameAndAddress6         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 NmAndAdr"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Ctry"`
}

// May be one of DSEC, EVNM, UKWN, ICOL, CONL, ELIG, INID, OTHR
type RejectionReasonV021Code string

type SecuritiesCollateralResponse1 struct {
	CollId   Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 CollId,omitempty"`
	AsstNb   Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 AsstNb,omitempty"`
	RspnTp   Status4Code             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 RspnTp"`
	RjctnRsn RejectionReasonV021Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 RjctnRsn,omitempty"`
	RjctnInf Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 RjctnInf,omitempty"`
}

// May be one of REJT, PACK
type Status4Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
