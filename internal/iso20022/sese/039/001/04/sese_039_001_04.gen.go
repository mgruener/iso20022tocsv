// Code generated by main. DO NOT EDIT.

package sese_039_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcknowledgedAcceptedStatus23Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason11 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Rsn"`
}

type AcknowledgementReason11 struct {
	Cd          AcknowledgementReason14Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason14Choice struct {
	Cd    AcknowledgementReason6Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

// May be one of ADEA, SMPG, OTHR, NSTP, LATE
type AcknowledgementReason6Code string

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

type AmountAndDirection51 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 CdtDbtInd"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 OrgnlCcyAndOrdrdAmt,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 DtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of OPEN, UKWN
type DateType4Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type DeniedReason10 struct {
	Cd          DeniedReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AddtlRsnInf,omitempty"`
}

type DeniedReason15Choice struct {
	Cd    DeniedReason6Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

// May be one of ADEA, CDCY, CDRE, DCAN, DSET, DPRG, DREP, LATE, OTHR, CDRG
type DeniedReason6Code string

type DeniedStatus15Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 NoSpcfdRsn"`
	Rsn        []DeniedReason10 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Rsn"`
}

type Document struct {
	SctiesSttlmTxModReqStsAdvc SecuritiesSettlementTransactionModificationRequestStatusAdviceV04 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SctiesSttlmTxModReqStsAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AmtsdVal"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type Identification14 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Id"`
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

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

type ModificationProcessingStatus7Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus23Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AckdAccptd"`
	PdgPrcg    PendingProcessingStatus13Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PdgPrcg"`
	Dnd        DeniedStatus15Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Dnd"`
	Rjctd      RejectionStatus18Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Rjctd"`
	Rprd       RepairStatus13Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Rprd"`
	Modfd      ModificationStatus4Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Modfd"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

type ModificationReason4 struct {
	Cd          ModificationReason4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	AddtlRsnInf Max210Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AddtlRsnInf,omitempty"`
}

type ModificationReason4Choice struct {
	Cd    ModifiedStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

type ModificationStatus4Choice struct {
	NoSpcfdRsn NoReasonCode          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 NoSpcfdRsn"`
	Rsn        []ModificationReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Rsn,omitempty"`
}

// May be one of MDBY, OTHR
type ModifiedStatusReason1Code string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Tp"`
}

type PartyIdentification44Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Ctry"`
}

type PartyIdentification71Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 NmAndAdr"`
}

type PartyIdentification91 struct {
	Id     PartyIdentification44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Id"`
	LEI    LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 LEI,omitempty"`
	PrcgId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PrcgId,omitempty"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PrtryId"`
}

type PartyIdentification93Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Ctry"`
}

type PartyIdentification98 struct {
	Id  PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 LEI,omitempty"`
}

type PartyIdentification99 struct {
	Id  PartyIdentification93Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 LEI,omitempty"`
}

type PartyIdentificationAndAccount117 struct {
	Id        PartyIdentification71Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Id"`
	LEI       LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 LEI,omitempty"`
	SfkpgAcct SecuritiesAccount19         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PrcgId,omitempty"`
}

type PendingProcessingReason10 struct {
	Cd          PendingProcessingReason12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	AddtlRsnInf Max210Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AddtlRsnInf,omitempty"`
}

type PendingProcessingReason12Choice struct {
	Cd    PendingProcessingReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	Prtry GenericIdentification30      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

// May be one of ADEA, BLOC, MUNO, NEXT, MINO, OTHR, DENO, CERT
type PendingProcessingReason3Code string

type PendingProcessingStatus13Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 NoSpcfdRsn"`
	Rsn        []PendingProcessingReason10 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Rsn"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Ctry"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PrtryRsn,omitempty"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 OrgnlAndCurFace"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type RejectionReason24Choice struct {
	Cd    RejectionReason31Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

type RejectionReason26 struct {
	Cd          RejectionReason24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AddtlRsnInf,omitempty"`
}

// May be one of SAFE, DQUA, ADEA, DSEC, LATE, CASH, DDEA, DTRD, PLCE, RTGS, NCRR, PHYS, REFE, DMON, MINO, BATC, MUNO, TXST, SETS, IIND, CAEV, CASY, DDAT, SETR, SDUT, INPS, OTHR, ICUS, ICAG, DEPT, IEXE, INVL, INVB, INVN, VALR
type RejectionReason31Code string

type RejectionStatus18Choice struct {
	NoSpcfdRsn NoReasonCode        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 NoSpcfdRsn"`
	Rsn        []RejectionReason26 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Rsn"`
}

type RepairReason10Choice struct {
	Cd    RepairReason4Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

// May be one of BATC, CAEV, CASH, CASY, DDAT, DDEA, DMON, DQUA, DSEC, DTRD, IIND, MINO, MUNO, NCRR, PHYS, PLCE, REFE, RTGS, SAFE, SETR, SETS, TXST, INPS, SDUT, OTHR, IEXE, ICAG, DEPT, ICUS
type RepairReason4Code string

type RepairReason9 struct {
	Cd          RepairReason10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AddtlRsnInf,omitempty"`
}

type RepairStatus13Choice struct {
	NoSpcfdRsn NoReasonCode    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 NoSpcfdRsn"`
	Rsn        []RepairReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Rsn"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Nm,omitempty"`
}

type SecuritiesSettlementTransactionModificationRequestStatusAdviceV04 struct {
	ModReqRef   Identification14                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 ModReqRef"`
	AcctOwnr    PartyIdentification98               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount19                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SfkpgAcct"`
	TxId        TransactionIdentifications33        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 TxId,omitempty"`
	ModPrcgSts  ModificationProcessingStatus7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 ModPrcgSts"`
	TxDtls      TransactionDetails81                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 TxDtls,omitempty"`
	SplmtryData []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SplmtryData,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Desc,omitempty"`
}

type SettlementDate10Choice struct {
	Dt   DateAndDateTimeChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Dt"`
	DtCd SettlementDateCode8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 DtCd"`
}

type SettlementDateCode8Choice struct {
	Cd    DateType4Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

type SettlementParties40 struct {
	Dpstry PartyIdentification91            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Pty5,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeDate5Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Dt"`
	DtCd TradeDateCode3Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 DtCd"`
}

type TradeDateCode3Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Prtry"`
}

type TransactionDetails81 struct {
	FinInstrmId     SecurityIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 FinInstrmId"`
	SctiesMvmntTp   ReceiveDelivery1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SctiesMvmntTp"`
	Pmt             DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Pmt"`
	SttlmQty        Quantity6Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SttlmQty"`
	SttlmAmt        AmountAndDirection51     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SttlmAmt,omitempty"`
	SttlmDt         SettlementDate10Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 SttlmDt"`
	TradDt          TradeDate5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 TradDt,omitempty"`
	DlvrgSttlmPties SettlementParties40      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties  SettlementParties40      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 RcvgSttlmPties,omitempty"`
	Invstr          PartyIdentification99    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Invstr,omitempty"`
}

type TransactionIdentifications33 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 PrcrTxId,omitempty"`
	OthrId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 OthrId,omitempty"`
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
