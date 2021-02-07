// Code generated by main. DO NOT EDIT.

package sese_022_001_04

type AcknowledgedAcceptedStatus24Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason12 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Rsn"`
}

type AcknowledgementReason12 struct {
	Cd          AcknowledgementReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason15Choice struct {
	Cd    AcknowledgementReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Prtry"`
}

// May be one of ADEA, SMPG, OTHR
type AcknowledgementReason3Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type Document struct {
	SctiesStsOrStmtQryStsAdvc SecuritiesStatusOrStatementQueryStatusAdviceV04 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 SctiesStsOrStmtQryStsAdvc"`
}

type DocumentIdentification30 struct {
	MsgNb DocumentNumber5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 MsgNb,omitempty"`
	Ref   Identification14      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Ref"`
}

type DocumentNumber12 struct {
	Nb   DocumentNumber5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Nb"`
	Refs []Identification15    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Refs"`
}

type DocumentNumber13 struct {
	Nb DocumentNumber5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Nb"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 SchmeNm,omitempty"`
}

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type Identification14 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Id"`
}

type Identification15 struct {
	AcctOwnrTxId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 PrcrTxId,omitempty"`
	CmonId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 CmonId,omitempty"`
	TradId            []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 TradId,omitempty"`
	MstrId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 MstrId,omitempty"`
	BsktId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 BsktId,omitempty"`
	IndxId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 IndxId,omitempty"`
	ListId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 ListId,omitempty"`
	PrgmId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 PrgmId,omitempty"`
	PoolId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 PoolId,omitempty"`
	CorpActnEvtId     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 CorpActnEvtId,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

// May be one of NORE
type NoReasonCode string

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 PrtryId"`
}

type PartyIdentification98 struct {
	Id  PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 LEI,omitempty"`
}

type ProcessingStatus55Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 AckdAccptd"`
	Rjctd      RejectionOrRepairStatus32Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Rjctd"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Prtry"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 PrtryRsn,omitempty"`
}

type RejectionAndRepairReason26Choice struct {
	Cd    RejectionReason24Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Prtry"`
}

type RejectionOrRepairReason26 struct {
	Cd          RejectionAndRepairReason26Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Cd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus32Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason26 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Rsn"`
}

// May be one of SAFE, DSEC, LATE, REFE, ADEA, OTHR, MISM
type RejectionReason24Code string

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Nm,omitempty"`
}

type SecuritiesStatusOrStatementQueryStatusAdviceV04 struct {
	QryDtls       DocumentIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 QryDtls"`
	AcctOwnr      PartyIdentification98    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount19      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 SfkpgAcct,omitempty"`
	StsOrStmtReqd StatusOrStatement7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 StsOrStmtReqd,omitempty"`
	PrcgSts       ProcessingStatus55Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 PrcgSts"`
	SplmtryData   []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 SplmtryData,omitempty"`
}

type StatusOrStatement7Choice struct {
	StsAdvc DocumentNumber12 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 StsAdvc"`
	Stmt    DocumentNumber13 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Stmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
