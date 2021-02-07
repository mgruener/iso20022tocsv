// Code generated by main. DO NOT EDIT.

package fxtr_030_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// May be one of POST, PREA, UNAL
type AllocationIndicator1Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type ClearingBrokerIdentification1 struct {
	SdInd     SideIndicator1Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 SdInd"`
	ClrBrkrId Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ClrBrkrId"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Prtry"`
}

// May be one of FULL, ONEW, PART, UNCO
type CollateralisationIndicator1Code string

// May be one of L, A, C, I, F, O, R, U
type CorporateSectorIdentifier1Code string

type CounterpartySideTransactionReporting1 struct {
	RptgJursdctn     Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 RptgJursdctn,omitempty"`
	RptgPty          PartyIdentification73Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 RptgPty,omitempty"`
	CtrPtySdUnqTxIdr []UniqueTransactionIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CtrPtySdUnqTxIdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 DtTm"`
}

type Document struct {
	FXTradBlkStsNtfctn ForeignExchangeTradeBulkStatusNotificationV04 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 FXTradBlkStsNtfctn"`
}

// May be no more than 42 items long
type Exact42Text string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

type ForeignExchangeTradeBulkStatusNotificationV04 struct {
	StsDtls     TradeData12          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 StsDtls"`
	TradData    []TradeData11        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 TradData"`
	MsgPgntn    Pagination           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 MsgPgntn,omitempty"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 SplmtryData,omitempty"`
}

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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max10Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max52Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress8 struct {
	Nm         Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Nm"`
	Adr        PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Adr,omitempty"`
	AltrntvIdr []Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AltrntvIdr,omitempty"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 LastPgInd"`
}

type PartyIdentification44 struct {
	AnyBIC     AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AnyBIC"`
	AltrntvIdr []Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AltrntvIdr,omitempty"`
}

type PartyIdentification59 struct {
	PtyNm      Max34Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PtyNm,omitempty"`
	AnyBIC     PartyIdentification44               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AnyBIC,omitempty"`
	AcctNb     Max34Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AcctNb,omitempty"`
	Adr        Max105Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Adr,omitempty"`
	ClrSysId   ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ClrSysId,omitempty"`
	LglNttyIdr LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 LglNttyIdr,omitempty"`
}

type PartyIdentification73Choice struct {
	NmAndAdr NameAndAddress8       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 NmAndAdr"`
	AnyBIC   PartyIdentification44 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AnyBIC"`
	PtyId    PartyIdentification59 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PtyId"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Ctry"`
}

type RegulatoryReporting4 struct {
	TradgSdTxRptg          []TradingSideTransactionReporting1      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 TradgSdTxRptg,omitempty"`
	CtrPtySdTxRptg         []CounterpartySideTransactionReporting1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CtrPtySdTxRptg,omitempty"`
	CntrlCtrPtyClrHs       PartyIdentification73Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CntrlCtrPtyClrHs,omitempty"`
	ClrBrkr                PartyIdentification73Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ClrBrkr,omitempty"`
	ClrXcptnPty            PartyIdentification73Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ClrXcptnPty,omitempty"`
	ClrBrkrId              ClearingBrokerIdentification1           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ClrBrkrId,omitempty"`
	ClrThrshldInd          bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ClrThrshldInd,omitempty"`
	ClrdPdctId             Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ClrdPdctId,omitempty"`
	UndrlygPdctIdr         UnderlyingProductIdentifier1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 UndrlygPdctIdr,omitempty"`
	AllcnInd               AllocationIndicator1Code                `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AllcnInd,omitempty"`
	CollstnInd             CollateralisationIndicator1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CollstnInd,omitempty"`
	ExctnVn                Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ExctnVn,omitempty"`
	ExctnTmstmp            DateAndDateTimeChoice                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ExctnTmstmp,omitempty"`
	NonStdFlg              bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 NonStdFlg,omitempty"`
	LkSwpId                Exact42Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 LkSwpId,omitempty"`
	FinNtrOfTheCtrPtyInd   bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 FinNtrOfTheCtrPtyInd,omitempty"`
	CollPrtflInd           bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CollPrtflInd,omitempty"`
	CollPrtflCd            Max10Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CollPrtflCd,omitempty"`
	PrtflCmprssnInd        bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PrtflCmprssnInd,omitempty"`
	CorpSctrInd            CorporateSectorIdentifier1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CorpSctrInd,omitempty"`
	TradWthNonEEACtrPtyInd bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 TradWthNonEEACtrPtyInd,omitempty"`
	NtrgrpTradInd          bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 NtrgrpTradInd,omitempty"`
	ComrclOrTrsrFincgInd   bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 ComrclOrTrsrFincgInd,omitempty"`
	AddtlRptgInf           Max210Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 AddtlRptgInf,omitempty"`
}

// May be one of CCPL, CLNT
type SideIndicator1Code string

type Status27Choice struct {
	Cd    TradeStatus6Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Cd"`
	Prtry Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Prtry"`
}

type Status28Choice struct {
	Cd    TradeStatus7Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Cd"`
	Prtry Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Prtry"`
}

type StatusAndSubStatus2 struct {
	StsCd    Status27Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 StsCd"`
	SubStsCd Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 SubStsCd,omitempty"`
}

// May be one of SMDY
type StatusSubType2Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeData11 struct {
	OrgtrRef           Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 OrgtrRef,omitempty"`
	MtchgSysUnqRef     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 MtchgSysUnqRef"`
	MtchgSysMtchgRef   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 MtchgSysMtchgRef,omitempty"`
	MtchgSysMtchdSdRef Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 MtchgSysMtchdSdRef,omitempty"`
	CurSttlmDt         ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CurSttlmDt,omitempty"`
	NewSttlmDt         ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 NewSttlmDt,omitempty"`
	CurStsDtTm         ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CurStsDtTm,omitempty"`
	PdctTp             Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PdctTp,omitempty"`
	SttlmSsnIdr        Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 SttlmSsnIdr,omitempty"`
	RgltryRptg         RegulatoryReporting4   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 RgltryRptg,omitempty"`
}

type TradeData12 struct {
	MsgId        Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 MsgId"`
	StsOrgtr     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 StsOrgtr,omitempty"`
	CurSts       StatusAndSubStatus2    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CurSts"`
	CurStsSubTp  StatusSubType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CurStsSubTp,omitempty"`
	CurStsDtTm   ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 CurStsDtTm"`
	PrvsSts      Status28Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PrvsSts,omitempty"`
	PrvsStsSubTp StatusSubType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PrvsStsSubTp,omitempty"`
	PdctTp       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PdctTp,omitempty"`
	SttlmSsnIdr  Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 SttlmSsnIdr,omitempty"`
	LkdRptId     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 LkdRptId,omitempty"`
}

// May be one of INVA, FMTC, SMAP, RJCT, RSCD, STLD, SPLI, UMTC, SMAT, FUMT, NETT, PFIX, OMTC
type TradeStatus6Code string

// May be one of INVA, UMTC, FMTC, SMAT, SUSP, SMAP, PFIX, FUMT
type TradeStatus7Code string

type TradingSideTransactionReporting1 struct {
	RptgJursdctn    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 RptgJursdctn,omitempty"`
	RptgPty         PartyIdentification73Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 RptgPty,omitempty"`
	TradgSdUnqTxIdr []UniqueTransactionIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 TradgSdUnqTxIdr,omitempty"`
}

// May be one of FORW, NDFO, SPOT, SWAP
type UnderlyingProductIdentifier1Code string

type UniqueTransactionIdentifier2 struct {
	UnqTxIdr    Max52Text   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 UnqTxIdr"`
	PrrUnqTxIdr []Max52Text `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 PrrUnqTxIdr,omitempty"`
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
