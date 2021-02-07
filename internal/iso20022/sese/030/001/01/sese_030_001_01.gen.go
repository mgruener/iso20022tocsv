// Code generated by main. DO NOT EDIT.

package sese_030_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AdditionalInformation3 struct {
	AcctOwnrTxId Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AcctOwnrTxId,omitempty"`
	ClssfctnTp   ClassificationType1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 ClssfctnTp,omitempty"`
	SfkpgAcct    SecuritiesAccount13                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 SfkpgAcct,omitempty"`
	FinInstrmId  SecurityIdentification11           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 FinInstrmId,omitempty"`
	Qty          FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Qty,omitempty"`
	FctvDt       DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 FctvDt,omitempty"`
	XpryDt       DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 XpryDt,omitempty"`
	CutOffDt     DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 CutOffDt,omitempty"`
	Invstr       PartyIdentification10Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Invstr,omitempty"`
	DlvrgPty1    PartyIdentificationAndAccount16    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 DlvrgPty1,omitempty"`
	RcvgPty1     PartyIdentificationAndAccount16    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 RcvgPty1,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateIdentification1 struct {
	Id    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 IdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of LAMI, NBOR, YBOR, RTRN
type AutoBorrowing2Code string

type AutomaticBorrowing2Choice struct {
	Cd    AutoBorrowing2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Prtry"`
}

// Must match the pattern [A-Z]{1,6}
type CFIIdentifier string

type ClassificationType1Choice struct {
	ClssfctnFinInstrm CFIIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AltrnClssfctn"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 DtTm"`
}

type Document struct {
	SctiesSttlmCondsModReq SecuritiesSettlementConditionsModificationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 SctiesSttlmCondsModReq"`
}

type DocumentIdentification11 struct {
	Id       Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Id"`
	CreDtTm  DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 CreDtTm,omitempty"`
	CpyDplct CopyDuplicate1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 CpyDplct,omitempty"`
}

type DocumentNumber1Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 LngNb"`
	PrtryNb GenericIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

type Extension2 struct {
	PlcAndNm   Max350Text         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PlcAndNm,omitempty"`
	XtnsnEnvlp ExtensionEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 XtnsnEnvlp"`
}

type ExtensionEnvelope1 struct {
	Item string `xml:",any"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AmtsdVal"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Issr,omitempty"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

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

type IdentificationSource1Choice struct {
	Dmst  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Dmst"`
	Prtry Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Prtry"`
}

type LinkageType1Choice struct {
	Cd    LinkageType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Prtry"`
}

// May be one of LINK, UNLK, SOFT
type LinkageType1Code string

type Linkages3 struct {
	PrcgPos ProcessingPosition2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PrcgPos,omitempty"`
	MsgNb   DocumentNumber1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 MsgNb,omitempty"`
	Ref     References6Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Ref"`
}

type MatchingDenied1Choice struct {
	Cd    MatchingProcess1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Prtry"`
}

// May be one of UNMT, MTRE
type MatchingProcess1Code string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Adr,omitempty"`
}

type PartyIdentification10Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 NmAndAdr"`
}

type PartyIdentification13Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PrtryId"`
}

type PartyIdentificationAndAccount16 struct {
	Id        PartyIdentification10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Id"`
	SfkpgAcct SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PrcgId,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Ctry"`
}

type PriorityNumeric1Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Nmrc"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Prtry"`
}

// May be one of AFTR, WITH, BEFR
type ProcessingPosition1Code string

type ProcessingPosition2Choice struct {
	Cd    ProcessingPosition1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Prtry"`
}

type References1 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AcctOwnrTxId,omitempty"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 MktInfrstrctrTxId,omitempty"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PoolId,omitempty"`
	CmonId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 CmonId,omitempty"`
	TradId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 TradId,omitempty"`
}

type References6Choice struct {
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AcctSvcrTxId"`
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AcctOwnrTxId"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 MktInfrstrctrTxId"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PoolId"`
	CmonId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 CmonId"`
	TradId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 TradId"`
}

type RequestDetails1 struct {
	Ref          References1               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Ref"`
	AutomtcBrrwg AutomaticBorrowing2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AutomtcBrrwg,omitempty"`
	RtnInd       bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 RtnInd,omitempty"`
	Lkg          LinkageType1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Lkg,omitempty"`
	Prty         PriorityNumeric1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Prty,omitempty"`
	OthrPrcg     []GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 OthrPrcg,omitempty"`
	PrtlSttlmInd bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 PrtlSttlmInd,omitempty"`
	SctiesRTGS   SecuritiesRTGS1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 SctiesRTGS,omitempty"`
	HldInd       bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 HldInd,omitempty"`
	MtchgDnl     MatchingDenied1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 MtchgDnl,omitempty"`
	UnltrlSplt   UnilateralSplit1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 UnltrlSplt,omitempty"`
	Lnkgs        []Linkages3               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Lnkgs,omitempty"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Nm,omitempty"`
}

type SecuritiesRTGS1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Prtry"`
}

type SecuritiesSettlementConditionsModificationRequestV01 struct {
	Id        DocumentIdentification11    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Id"`
	AcctOwnr  PartyIdentification13Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AcctOwnr,omitempty"`
	SfkpgAcct SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 SfkpgAcct"`
	ReqDtls   []RequestDetails1           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 ReqDtls"`
	AddtlInf  []AdditionalInformation3    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 AddtlInf,omitempty"`
	MsgOrgtr  PartyIdentification10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 MsgOrgtr,omitempty"`
	MsgRcpt   PartyIdentification10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 MsgRcpt,omitempty"`
	Xtnsn     []Extension2                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Xtnsn,omitempty"`
}

// May be one of TRAD
type SecuritiesTransactionType5Code string

type SecurityIdentification11 struct {
	Id   SecurityIdentification11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Id"`
	Desc Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Desc,omitempty"`
}

type SecurityIdentification11Choice struct {
	ISIN   ISINIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 ISIN"`
	OthrId AlternateIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 OthrId"`
}

type UnilateralSplit1Choice struct {
	Cd    SecuritiesTransactionType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Prtry"`
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