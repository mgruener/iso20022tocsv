// Code generated by main. DO NOT EDIT.

package sese_022_002_05_0

type AcknowledgedAcceptedStatus31Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Rsn"`
}

type AcknowledgementReason19 struct {
	Cd          AcknowledgementReason22Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Cd"`
	AddtlRsnInf RestrictedFINXMax210Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason22Choice struct {
	Cd    AcknowledgementReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Cd"`
	Prtry GenericIdentification47    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Prtry"`
}

// May be one of ADEA, SMPG, OTHR
type AcknowledgementReason3Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type Document struct {
	SctiesStsOrStmtQryStsAdvc SecuritiesStatusOrStatementQueryStatusAdvice002V05 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 SctiesStsOrStmtQryStsAdvc"`
}

type DocumentIdentification48 struct {
	MsgNb DocumentNumber6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 MsgNb,omitempty"`
	Ref   Identification16      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Ref"`
}

type DocumentNumber14 struct {
	Nb DocumentNumber6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Nb"`
}

type DocumentNumber17 struct {
	Nb   DocumentNumber6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Nb"`
	Refs []Identification27    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Refs"`
}

type DocumentNumber6Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 LngNb"`
	PrtryNb GenericIdentification86           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 SchmeNm,omitempty"`
}

type GenericIdentification86 struct {
	Id      RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 SchmeNm,omitempty"`
}

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type Identification16 struct {
	Id RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Id"`
}

type Identification27 struct {
	AcctOwnrTxId      RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 AcctOwnrTxId"`
	AcctSvcrTxId      RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 PrcrTxId,omitempty"`
	CmonId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 CmonId,omitempty"`
	TradId            []RestrictedFINXMax52Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 TradId,omitempty"`
	MstrId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 MstrId,omitempty"`
	BsktId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 BsktId,omitempty"`
	IndxId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 IndxId,omitempty"`
	ListId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 ListId,omitempty"`
	PrgmId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 PrgmId,omitempty"`
	PoolId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 PoolId,omitempty"`
	CorpActnEvtId     RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 CorpActnEvtId,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max350Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

// May be one of NORE
type NoReasonCode string

type PartyIdentification136Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 PrtryId"`
}

type PartyIdentification156 struct {
	Id  PartyIdentification136Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 LEI,omitempty"`
}

type ProcessingStatus64Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus31Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 AckdAccptd"`
	Rjctd      RejectionOrRepairStatus37Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Rjctd"`
	Prtry      ProprietaryStatusAndReason7        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Prtry"`
}

type ProprietaryReason5 struct {
	Rsn         GenericIdentification47  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Rsn,omitempty"`
	AddtlRsnInf RestrictedFINXMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason7 struct {
	PrtrySts GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 PrtrySts"`
	PrtryRsn []ProprietaryReason5    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 PrtryRsn,omitempty"`
}

type RejectionAndRepairReason31Choice struct {
	Cd    RejectionReason24Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Prtry"`
}

type RejectionOrRepairReason31 struct {
	Cd          RejectionAndRepairReason31Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Cd"`
	AddtlRsnInf RestrictedFINXMax210Text         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus37Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason31 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Rsn"`
}

// May be one of SAFE, DSEC, LATE, REFE, ADEA, OTHR, MISM
type RejectionReason24Code string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,210}
type RestrictedFINXMax210Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax30Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,35}
type RestrictedFINXMax35Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,52}
type RestrictedFINXMax52Text string

type SecuritiesAccount30 struct {
	Id RestrictedFINXMax35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Id"`
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Nm,omitempty"`
}

type SecuritiesStatusOrStatementQueryStatusAdvice002V05 struct {
	QryDtls       DocumentIdentification48  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 QryDtls"`
	AcctOwnr      PartyIdentification156    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount30       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 SfkpgAcct,omitempty"`
	StsOrStmtReqd StatusOrStatement10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 StsOrStmtReqd,omitempty"`
	PrcgSts       ProcessingStatus64Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 PrcgSts"`
	SplmtryData   []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 SplmtryData,omitempty"`
}

type StatusOrStatement10Choice struct {
	StsAdvc DocumentNumber17 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 StsAdvc"`
	Stmt    DocumentNumber14 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Stmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
