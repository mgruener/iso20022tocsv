// Code generated by main. DO NOT EDIT.

package sese_040_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmountAndDirection8 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 CdtDbtInd"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 OrgnlCcyAndOrdrdAmt,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type ConsentOrRejectionReason2Choice struct {
	Cd    CounterpartyResponseStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Cd"`
	Prtry GenericIdentification20               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Prtry"`
}

type ConsentReason2 struct {
	Cd          ConsentOrRejectionReason2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Cd"`
	AddtlRsnInf Max210Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AddtlRsnInf,omitempty"`
}

type ConsentStatus2Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 NoSpcfdRsn"`
	Rsn        []ConsentReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Rsn"`
}

// May be one of CPTR, CPCX, CPMD
type CounterpartyResponseStatusReason1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 DtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of OPEN, UKWN
type DateType4Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesSttlmTxCtrPtyRspn SecuritiesSettlementTransactionCounterpartyResponseV01 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SctiesSttlmTxCtrPtyRspn"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AmtsdVal"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type NoSpecifiedReason1 struct {
	NoSpcfdRsn NoReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 NoSpcfdRsn"`
}

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Tp"`
}

type PartyIdentification37Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Ctry"`
}

type PartyIdentification44Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Ctry"`
}

type PartyIdentification45Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 NmAndAdr"`
}

type PartyIdentification46 struct {
	Id     PartyIdentification44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Id"`
	PrcgId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 PrcgId,omitempty"`
}

type PartyIdentificationAndAccount44 struct {
	Id        PartyIdentification45Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Id"`
	SfkpgAcct SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 PrcgId,omitempty"`
}

type PendingStatus20Choice struct {
	Fwdd        NoSpecifiedReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Fwdd"`
	UdrInvstgtn NoSpecifiedReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 UdrInvstgtn"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Ctry"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 OrgnlAndCurFace"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type RejectionReason12 struct {
	Cd          ConsentOrRejectionReason2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Cd"`
	AddtlRsnInf Max210Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AddtlRsnInf,omitempty"`
}

type RejectionStatus7Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 NoSpcfdRsn"`
	Rsn        RejectionReason12 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Rsn"`
}

type ResponseStatus3Choice struct {
	Cnsntd ConsentStatus2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Cnsntd"`
	Rjctd  RejectionStatus7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Rjctd"`
	Pdg    PendingStatus20Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Pdg"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Nm,omitempty"`
}

type SecuritiesSettlementTransactionCounterpartyResponseV01 struct {
	TxId        TransactionIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 TxId"`
	RspnSts     ResponseStatus3Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 RspnSts"`
	TxDtls      TransactionDetails40       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 TxDtls,omitempty"`
	SplmtryData []SupplementaryData1       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SplmtryData,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Desc,omitempty"`
}

type SettlementDate2Choice struct {
	Dt   DateAndDateTimeChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Dt"`
	DtCd SettlementDateCode2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 DtCd"`
}

type SettlementDateCode2Choice struct {
	Cd    DateType4Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Prtry"`
}

type SettlementParties13 struct {
	Dpstry PartyIdentification46           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Pty5,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeDate1Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Dt"`
	DtCd TradeDateCode1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 DtCd"`
}

type TradeDateCode1Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Prtry"`
}

type TransactionDetails40 struct {
	FinInstrmId     SecurityIdentification14    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 FinInstrmId"`
	SctiesMvmntTp   ReceiveDelivery1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SctiesMvmntTp"`
	Pmt             DeliveryReceiptType2Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Pmt"`
	SttlmQty        Quantity6Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SttlmQty"`
	SfkpgAcct       SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SfkpgAcct"`
	SttlmAmt        AmountAndDirection8         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SttlmAmt,omitempty"`
	SttlmDt         SettlementDate2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 SttlmDt"`
	TradDt          TradeDate1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 TradDt,omitempty"`
	DlvrgSttlmPties SettlementParties13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties  SettlementParties13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 RcvgSttlmPties,omitempty"`
	Invstr          PartyIdentification37Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Invstr,omitempty"`
}

type TransactionIdentification2 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AcctOwnrTxId,omitempty"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 PrcrTxId,omitempty"`
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
