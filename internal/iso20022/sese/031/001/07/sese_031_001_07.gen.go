// Code generated by main. DO NOT EDIT.

package sese_031_001_07

type AcknowledgedAcceptedStatus21Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Rsn"`
}

type AcknowledgementReason12Choice struct {
	Cd    AcknowledgementReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

// May be one of ADEA, SMPG, OTHR, CDCY, CDRG, CDRE, NSTP, RQWV, LATE
type AcknowledgementReason5Code string

type AcknowledgementReason9 struct {
	Cd          AcknowledgementReason12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AddtlRsnInf,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of LAMI, NBOR, YBOR, RTRN
type AutoBorrowing2Code string

type AutomaticBorrowing7Choice struct {
	Cd    AutoBorrowing2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

type DeniedReason10 struct {
	Cd          DeniedReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AddtlRsnInf,omitempty"`
}

type DeniedReason15Choice struct {
	Cd    DeniedReason6Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

// May be one of ADEA, CDCY, CDRE, DCAN, DSET, DPRG, DREP, LATE, OTHR, CDRG
type DeniedReason6Code string

type DeniedStatus15Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 NoSpcfdRsn"`
	Rsn        []DeniedReason10 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Rsn"`
}

type Document struct {
	SctiesSttlmCondModStsAdvc SecuritiesSettlementConditionModificationStatusAdviceV07 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 SctiesSttlmCondModStsAdvc"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 SchmeNm,omitempty"`
}

type HoldIndicator6 struct {
	Ind bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Ind"`
	Rsn []RegistrationReason5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Rsn,omitempty"`
}

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type Identification14 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Id"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be one of LINK, UNLK, SOFT
type LinkageType1Code string

type LinkageType3Choice struct {
	Cd    LinkageType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

type Linkages39 struct {
	PrcgPos ProcessingPosition8Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PrcgPos,omitempty"`
	MsgNb   DocumentNumber5Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 MsgNb,omitempty"`
	Ref     References46Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Ref"`
	RefOwnr PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 RefOwnr,omitempty"`
}

type MatchingDenied3Choice struct {
	Cd    MatchingProcess1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

// May be one of UNMT, MTRE
type MatchingProcess1Code string

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
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PrtryId"`
}

type PartyIdentification98 struct {
	Id  PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 LEI,omitempty"`
}

type PendingReason16 struct {
	Cd          PendingReason28Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AddtlRsnInf,omitempty"`
}

type PendingReason28Choice struct {
	Cd    PendingReason6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

// May be one of ADEA, CONF, OTHR, CDRG, CDCY, CDRE
type PendingReason6Code string

type PendingStatus38Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 NoSpcfdRsn"`
	Rsn        []PendingReason16 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Rsn"`
}

type PriorityNumeric4Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Nmrc"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

// May be one of AFTE, BEFO, WITH
type ProcessingPosition4Code string

type ProcessingPosition8Choice struct {
	Cd    ProcessingPosition4Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

type ProcessingStatus50Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AckdAccptd"`
	Rjctd      RejectionOrRepairStatus31Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Rjctd"`
	Cmpltd     ProprietaryReason4                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cmpltd"`
	Dnd        DeniedStatus15Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Dnd"`
	Pdg        PendingStatus38Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Pdg"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PrtryRsn,omitempty"`
}

type References18 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AcctOwnrTxId,omitempty"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PrcrTxId,omitempty"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PoolId,omitempty"`
	CmonId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 CmonId,omitempty"`
	TradId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 TradId,omitempty"`
}

type References46Choice struct {
	SctiesSttlmTxId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 SctiesSttlmTxId"`
	IntraPosMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 IntraPosMvmntId"`
	IntraBalMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 IntraBalMvmntId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AcctSvcrTxId"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 MktInfrstrctrTxId"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PoolId"`
	CmonId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 CmonId"`
	TradId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 TradId"`
	OthrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 OthrTxId"`
}

type Registration10Choice struct {
	Cd    Registration2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

// May be one of PTYH, CSDH, CDEL, CVAL
type Registration2Code string

type RegistrationReason5 struct {
	Cd       Registration10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	AddtlInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AddtlInf,omitempty"`
}

type RejectionAndRepairReason25Choice struct {
	Cd    RejectionReason27Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

type RejectionOrRepairReason25 struct {
	Cd          RejectionAndRepairReason25Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus31Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason25 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Rsn"`
}

// May be one of ADEA, LATE, SAFE, NRGM, NRGN, OTHR, REFE, INVM, INVL
type RejectionReason27Code string

type RequestDetails15 struct {
	Ref          References18                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Ref"`
	AutomtcBrrwg AutomaticBorrowing7Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AutomtcBrrwg,omitempty"`
	RtnInd       bool                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 RtnInd,omitempty"`
	Lkg          LinkageType3Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Lkg,omitempty"`
	Prty         PriorityNumeric4Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prty,omitempty"`
	OthrPrcg     []GenericIdentification30           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 OthrPrcg,omitempty"`
	PrtlSttlmInd SettlementTransactionCondition5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PrtlSttlmInd,omitempty"`
	SctiesRTGS   SecuritiesRTGS4Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 SctiesRTGS,omitempty"`
	HldInd       HoldIndicator6                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 HldInd,omitempty"`
	MtchgDnl     MatchingDenied3Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 MtchgDnl,omitempty"`
	UnltrlSplt   UnilateralSplit3Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 UnltrlSplt,omitempty"`
	Lnkgs        []Linkages39                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Lnkgs,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Nm,omitempty"`
}

type SecuritiesRTGS4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}

type SecuritiesSettlementConditionModificationStatusAdviceV07 struct {
	ReqRef      Identification14         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 ReqRef"`
	AcctOwnr    PartyIdentification98    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount19      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 SfkpgAcct,omitempty"`
	ReqDtls     RequestDetails15         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 ReqDtls,omitempty"`
	PrcgSts     ProcessingStatus50Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PrcgSts"`
	SplmtryData []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 SplmtryData,omitempty"`
}

// May be one of TRAD
type SecuritiesTransactionType5Code string

// May be one of PART, NPAR, PARC, PARQ
type SettlementTransactionCondition5Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UnilateralSplit3Choice struct {
	Cd    SecuritiesTransactionType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Prtry"`
}