// Code generated by main. DO NOT EDIT.

package secl_008_001_03

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

type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternatePartyIdentification4 struct {
	IdTp    IdentificationType6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 IdTp"`
	Ctry    CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Ctry"`
	AltrnId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 AltrnId"`
}

type AmountAndDirection27 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms17            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 FXDtls,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BuyIn3 struct {
	BuyInNtfctnId Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 BuyInNtfctnId"`
	ReqForDelyInd bool                               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 ReqForDelyInd"`
	NbOfDays      float64                            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 NbOfDays"`
	InitlQty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 InitlQty"`
	CvrdQty       FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 CvrdQty"`
	UcvrdQty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 UcvrdQty"`
}

type BuyInResponseV03 struct {
	TxId                 Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 TxId,omitempty"`
	BuyInRspnDtls        BuyIn3                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 BuyInRspnDtls"`
	OrgnlSttlmOblgtnDtls SettlementObligation7 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 OrgnlSttlmOblgtnDtls,omitempty"`
	SplmtryData          []SupplementaryData1  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 SplmtryData,omitempty"`
}

// May be one of HOUS, CLIE, LIPR
type ClearingAccountType1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type Document struct {
	BuyInRspn BuyInResponseV03 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 BuyInRspn"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 AmtsdVal"`
}

type ForeignExchangeTerms17 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 RsltgAmt"`
}

type GenericIdentification29 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 SchmeNm,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 SchmeNm,omitempty"`
}

type GenericIdentification40 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 SchmeNm,omitempty"`
}

type GenericIdentification58 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id,omitempty"`
	Tp GenericIdentification40 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Tp"`
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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Prtry"`
}

type IdentificationType6Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Prtry"`
}

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
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Adr,omitempty"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Adr"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Tp"`
}

type PartyIdentification33Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 AnyBIC"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 PrtryId"`
	NmAndAdr NameAndAddress6         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 NmAndAdr"`
}

type PartyIdentification34Choice struct {
	BIC      AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 BIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Ctry"`
}

type PartyIdentification35Choice struct {
	BIC     AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 BIC"`
	PrtryId GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 PrtryId"`
}

type PartyIdentificationAndAccount31 struct {
	Id       PartyIdentification33Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id"`
	AltrnId  AlternatePartyIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 AltrnId,omitempty"`
	AddtlInf PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 AddtlInf,omitempty"`
	ClrAcct  SecuritiesAccount18           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 ClrAcct,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 RegnDtls,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Ctry"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Ctry"`
}

type Price4 struct {
	Val PriceRateOrAmountChoice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Val"`
	Tp  PriceValueType7Code     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Tp,omitempty"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Amt"`
}

// May be one of DISC, PREM, PARV, YIEL, SPRE, PEUN, ABSO, TEDP, TEDY, FICT, VACT, PRCT, ACTU
type PriceValueType7Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat7Choice struct {
	Id      SafekeepingPlaceTypeAndText1             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 TpAndId"`
	Prtry   GenericIdentification58                  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id"`
}

type SafekeepingPlaceTypeAndText1 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id,omitempty"`
}

type SecuritiesAccount18 struct {
	Id Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id"`
	Tp ClearingAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Tp"`
	Nm Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Nm,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Nm,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Desc,omitempty"`
}

type SettlementObligation7 struct {
	CSDTxId          Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 CSDTxId,omitempty"`
	CntrlCtrPtyTxId  Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 CntrlCtrPtyTxId,omitempty"`
	PrvsBuyInId      Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 PrvsBuyInId,omitempty"`
	DlvryAcct        SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 DlvryAcct,omitempty"`
	SfkpgPlc         SafekeepingPlaceFormat7Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 SfkpgPlc,omitempty"`
	SfkpgAcct        SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 SfkpgAcct,omitempty"`
	ClrSgmt          PartyIdentification35Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 ClrSgmt,omitempty"`
	NonClrMmb        PartyIdentificationAndAccount31    `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 NonClrMmb,omitempty"`
	IntnddSttlmDt    ISODate                            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 IntnddSttlmDt,omitempty"`
	FinInstrmId      SecurityIdentification14           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 FinInstrmId"`
	TradDt           ISODate                            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 TradDt,omitempty"`
	DealPric         Price4                             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 DealPric,omitempty"`
	Qty              FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Qty"`
	Dpstry           PartyIdentification34Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Dpstry,omitempty"`
	RmngQtyToBeSttld FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 RmngQtyToBeSttld,omitempty"`
	SttlmAmt         AmountAndDirection27               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 SttlmAmt"`
	RmngAmtToBeSttld AmountAndDirection27               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 RmngAmtToBeSttld,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

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
