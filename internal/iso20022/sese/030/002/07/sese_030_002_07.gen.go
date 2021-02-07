// Code generated by main. DO NOT EDIT.

package sese_030_002_07

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AdditionalInformation14 struct {
	AcctOwnrTxId RestrictedFINXMax16Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AcctOwnrTxId,omitempty"`
	ClssfctnTp   ClassificationType33Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 ClssfctnTp,omitempty"`
	SfkpgAcct    SecuritiesAccount30                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SfkpgAcct,omitempty"`
	FinInstrmId  SecurityIdentification20            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 FinInstrmId,omitempty"`
	Qty          FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Qty,omitempty"`
	FctvDt       DateAndDateTimeChoice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 FctvDt,omitempty"`
	XpryDt       DateAndDateTimeChoice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 XpryDt,omitempty"`
	CutOffDt     DateAndDateTimeChoice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 CutOffDt,omitempty"`
	Invstr       PartyIdentification111              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Invstr,omitempty"`
	DlvrgPty1    PartyIdentificationAndAccount146    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 DlvrgPty1,omitempty"`
	RcvgPty1     PartyIdentificationAndAccount146    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 RcvgPty1,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of LAMI, NBOR, YBOR, RTRN
type AutoBorrowing2Code string

type AutomaticBorrowing11Choice struct {
	Cd    AutoBorrowing2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prtry"`
}

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type ClassificationType33Choice struct {
	ClssfctnFinInstrm CFIOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification86 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AltrnClssfctn"`
}

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 DtTm"`
}

type Document struct {
	SctiesSttlmCondsModReq SecuritiesSettlementConditionsModificationRequest002V07 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SctiesSttlmCondsModReq"`
}

type DocumentNumber16Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 LngNb"`
	PrtryNb GenericIdentification163          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity15Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AmtsdVal"`
}

type GenericIdentification163 struct {
	Id      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Id"`
	Issr    Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Issr"`
	SchmeNm Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SchmeNm,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SchmeNm,omitempty"`
}

type GenericIdentification86 struct {
	Id      RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SchmeNm,omitempty"`
}

type HoldIndicator7 struct {
	Ind bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Ind"`
	Rsn []RegistrationReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Rsn,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type IdentificationSource4Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Cd"`
	Prtry RestrictedFINExact2Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be one of LINK, UNLK, SOFT
type LinkageType1Code string

type LinkageType4Choice struct {
	Cd    LinkageType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prtry"`
}

type Linkages44 struct {
	PrcgPos ProcessingPosition18Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PrcgPos,omitempty"`
	MsgNb   DocumentNumber16Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 MsgNb,omitempty"`
	Ref     References54Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Ref"`
	RefOwnr PartyIdentification103Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 RefOwnr,omitempty"`
}

type MatchingDenied4Choice struct {
	Cd    MatchingProcess1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prtry"`
}

// May be one of UNMT, MTRE
type MatchingProcess1Code string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress12 struct {
	Nm RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Nm"`
}

type OtherIdentification2 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Sfx,omitempty"`
	Tp  IdentificationSource4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Tp"`
}

type PartyIdentification103Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PrtryId"`
}

type PartyIdentification104Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AnyBIC"`
	PrtryId  GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PrtryId"`
	NmAndAdr NameAndAddress12        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 NmAndAdr"`
}

type PartyIdentification111 struct {
	Id  PartyIdentification104Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 LEI,omitempty"`
}

type PartyIdentification119 struct {
	Id  PartyIdentification103Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 LEI,omitempty"`
}

type PartyIdentificationAndAccount146 struct {
	Id        PartyIdentification104Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Id"`
	LEI       LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 LEI,omitempty"`
	SfkpgAcct SecuritiesAccount30          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SfkpgAcct,omitempty"`
	PrcgId    RestrictedFINXMax16Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PrcgId,omitempty"`
}

type PriorityNumeric5Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Nmrc"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prtry"`
}

type ProcessingPosition18Choice struct {
	Cd    ProcessingPosition4Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prtry"`
}

// May be one of AFTE, BEFO, WITH
type ProcessingPosition4Code string

type References21 struct {
	AcctOwnrTxId      RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AcctOwnrTxId,omitempty"`
	AcctSvcrTxId      RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PrcrTxId,omitempty"`
	PoolId            RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PoolId,omitempty"`
	CmonId            RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 CmonId,omitempty"`
	TradId            RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 TradId,omitempty"`
}

type References54Choice struct {
	SctiesSttlmTxId   RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SctiesSttlmTxId"`
	IntraPosMvmntId   RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 IntraPosMvmntId"`
	IntraBalMvmntId   RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 IntraBalMvmntId"`
	AcctSvcrTxId      RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AcctSvcrTxId"`
	MktInfrstrctrTxId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 MktInfrstrctrTxId"`
	PoolId            RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PoolId"`
	CmonId            RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 CmonId"`
	TradId            RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 TradId"`
	OthrTxId          RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 OthrTxId"`
}

type Registration12Choice struct {
	Cd    Registration2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prtry"`
}

// May be one of PTYH, CSDH, CDEL, CVAL
type Registration2Code string

type RegistrationReason6 struct {
	Cd       Registration12Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Cd"`
	AddtlInf RestrictedFINXMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AddtlInf,omitempty"`
}

type RequestDetails16 struct {
	Ref          References21                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Ref"`
	AutomtcBrrwg AutomaticBorrowing11Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AutomtcBrrwg,omitempty"`
	RtnInd       bool                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 RtnInd,omitempty"`
	Lkg          LinkageType4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Lkg,omitempty"`
	Prty         PriorityNumeric5Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prty,omitempty"`
	OthrPrcg     []GenericIdentification47           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 OthrPrcg,omitempty"`
	PrtlSttlmInd SettlementTransactionCondition5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PrtlSttlmInd,omitempty"`
	SctiesRTGS   SecuritiesRTGS5Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SctiesRTGS,omitempty"`
	HldInd       HoldIndicator7                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 HldInd,omitempty"`
	MtchgDnl     MatchingDenied4Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 MtchgDnl,omitempty"`
	UnltrlSplt   UnilateralSplit4Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 UnltrlSplt,omitempty"`
	Lnkgs        []Linkages44                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Lnkgs,omitempty"`
}

// Must match the pattern XX|TS
type RestrictedFINExact2Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,140}
type RestrictedFINXMax140Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,210}
type RestrictedFINXMax210Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax30Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,31}
type RestrictedFINXMax31Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,35}
type RestrictedFINXMax35Text string

type SecuritiesAccount30 struct {
	Id RestrictedFINXMax35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Id"`
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Nm,omitempty"`
}

type SecuritiesRTGS5Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Ind"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prtry"`
}

type SecuritiesSettlementConditionsModificationRequest002V07 struct {
	AcctOwnr    PartyIdentification119    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount30       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SfkpgAcct"`
	ReqDtls     []RequestDetails16        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 ReqDtls"`
	AddtlInf    []AdditionalInformation14 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 AddtlInf,omitempty"`
	SplmtryData []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 SplmtryData,omitempty"`
}

// May be one of TRAD
type SecuritiesTransactionType5Code string

type SecurityIdentification20 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 ISIN,omitempty"`
	OthrId []OtherIdentification2   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Desc,omitempty"`
}

// May be one of PART, NPAR, PARC, PARQ
type SettlementTransactionCondition5Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UnilateralSplit4Choice struct {
	Cd    SecuritiesTransactionType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Prtry"`
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
