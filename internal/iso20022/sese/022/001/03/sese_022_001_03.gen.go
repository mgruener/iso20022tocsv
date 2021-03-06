// Code generated by main. DO NOT EDIT.

package sese_022_001_03

type AcknowledgedAcceptedStatus12Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Rsn"`
}

type AcknowledgementReason1 struct {
	Cd          AcknowledgementReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Cd"`
	AddtlRsnInf Max210Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason1Choice struct {
	Cd    AcknowledgementReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Prtry"`
}

// May be one of ADEA, SMPG, OTHR
type AcknowledgementReason3Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type Document struct {
	SctiesStsOrStmtQryStsAdvc SecuritiesStatusOrStatementQueryStatusAdviceV03 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 SctiesStsOrStmtQryStsAdvc"`
}

type DocumentIdentification24 struct {
	MsgNb DocumentNumber1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 MsgNb,omitempty"`
	Ref   Identification1       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Ref"`
}

type DocumentNumber1 struct {
	Nb DocumentNumber1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Nb"`
}

type DocumentNumber1Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 LngNb"`
	PrtryNb GenericIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 PrtryNb"`
}

type DocumentNumber9 struct {
	Nb   DocumentNumber1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Nb"`
	Refs []Identification11    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Refs"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 SchmeNm,omitempty"`
}

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type Identification1 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Id"`
}

type Identification11 struct {
	AcctOwnrTxId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 PrcrTxId,omitempty"`
	CmonId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 CmonId,omitempty"`
	TradId            []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 TradId,omitempty"`
	MstrId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 MstrId,omitempty"`
	BsktId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 BsktId,omitempty"`
	IndxId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 IndxId,omitempty"`
	ListId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 ListId,omitempty"`
	PrgmId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 PrgmId,omitempty"`
	PoolId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 PoolId,omitempty"`
	CorpActnEvtId     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 CorpActnEvtId,omitempty"`
}

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

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 PrtryId"`
}

type ProcessingStatus22Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 AckdAccptd"`
	Rjctd      RejectionOrRepairStatus18Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Rjctd"`
	Prtry      ProprietaryStatusAndReason1        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Prtry"`
}

type ProprietaryReason1 struct {
	Rsn         GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason1 struct {
	PrtrySts GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 PrtrySts"`
	PrtryRsn []ProprietaryReason1    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 PrtryRsn,omitempty"`
}

type RejectionAndRepairReason2Choice struct {
	Cd    RejectionReason24Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Prtry"`
}

type RejectionOrRepairReason2 struct {
	Cd          RejectionAndRepairReason2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Cd"`
	AddtlRsnInf Max210Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus18Choice struct {
	NoSpcfdRsn NoReasonCode               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Rsn"`
}

// May be one of SAFE, DSEC, LATE, REFE, ADEA, OTHR, MISM
type RejectionReason24Code string

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Nm,omitempty"`
}

type SecuritiesStatusOrStatementQueryStatusAdviceV03 struct {
	QryDtls       DocumentIdentification24    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 QryDtls"`
	AcctOwnr      PartyIdentification36Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 SfkpgAcct,omitempty"`
	StsOrStmtReqd StatusOrStatement5Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 StsOrStmtReqd,omitempty"`
	PrcgSts       ProcessingStatus22Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 PrcgSts"`
	SplmtryData   []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 SplmtryData,omitempty"`
}

type StatusOrStatement5Choice struct {
	StsAdvc DocumentNumber9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 StsAdvc"`
	Stmt    DocumentNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Stmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
