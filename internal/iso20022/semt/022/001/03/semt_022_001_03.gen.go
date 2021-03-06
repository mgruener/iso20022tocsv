// Code generated by main. DO NOT EDIT.

package semt_022_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcknowledgedAcceptedStatus21Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type AcknowledgedAcceptedStatus23Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason11 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type AcknowledgedAcceptedStatus24Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason12 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type AcknowledgementReason11 struct {
	Cd          AcknowledgementReason14Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason12 struct {
	Cd          AcknowledgementReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason12Choice struct {
	Cd    AcknowledgementReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type AcknowledgementReason14Choice struct {
	Cd    AcknowledgementReason6Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type AcknowledgementReason15Choice struct {
	Cd    AcknowledgementReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

// May be one of ADEA, SMPG, OTHR
type AcknowledgementReason3Code string

// May be one of ADEA, SMPG, OTHR, CDCY, CDRG, CDRE, NSTP, RQWV, LATE
type AcknowledgementReason5Code string

// May be one of ADEA, SMPG, OTHR, NSTP, LATE
type AcknowledgementReason6Code string

type AcknowledgementReason9 struct {
	Cd          AcknowledgementReason12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancellationReason10 struct {
	Cd          CancellationReason21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type CancellationReason19Choice struct {
	Cd    CancelledStatusReason13Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type CancellationReason21Choice struct {
	Cd    CancelledStatusReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type CancellationReason9 struct {
	Cd          CancellationReason19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type CancellationStatus14Choice struct {
	NoSpcfdRsn NoReasonCode          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []CancellationReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type CancellationStatus15Choice struct {
	NoSpcfdRsn NoReasonCode           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []CancellationReason10 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

// May be one of CANI, CANS, CSUB, CXLR, CANT, CANZ, CORP, SCEX, OTHR, CTHP
type CancelledStatusReason13Code string

// May be one of CANI, OTHR
type CancelledStatusReason5Code string

type DeniedReason10 struct {
	Cd          DeniedReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type DeniedReason11 struct {
	Cd          DeniedReason16Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type DeniedReason15Choice struct {
	Cd    DeniedReason6Code       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type DeniedReason16Choice struct {
	Cd    DeniedReason4Code       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

// May be one of ADEA, DCAN, DPRG, DREP, DSET, LATE, OTHR, CDRG, CDCY, CDRE
type DeniedReason4Code string

// May be one of ADEA, CDCY, CDRE, DCAN, DSET, DPRG, DREP, LATE, OTHR, CDRG
type DeniedReason6Code string

type DeniedStatus15Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []DeniedReason10 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type DeniedStatus16Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []DeniedReason11 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type Document struct {
	SctiesSttlmTxAudtTrlRpt SecuritiesSettlementTransactionAuditTrailReportV03 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 SctiesSttlmTxAudtTrlRpt"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// May be one of AWMO, BYIY, CLAT, ADEA, CANR, CAIS, OBJT, AWSH, PHSE, STCD, DOCY, MLAT, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, PART, NOFX, CMON, YCOL, COLL, DEPO, FLIM, INCA, LINK, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, IAAD, OTHR, PHCK, BENO, BOTH, CLHT, DENO, DISA, DKNY, FROZ, LAAW, LATE, LIQU, PRCY, REGT, SETS, CERT, PRSY, CDLR, CSDH, CVAL, INBC
type FailingReason2Code string

type FailingReason8 struct {
	Cd          FailingReason8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type FailingReason8Choice struct {
	Cd    FailingReason2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type FailingStatus10Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []FailingReason8 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Issr,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type Identification14 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Id"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type MatchingStatus25Choice struct {
	Mtchd  ProprietaryReason4          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Mtchd"`
	Umtchd UnmatchedStatus17Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Umtchd"`
	Prtry  ProprietaryStatusAndReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

type ModificationProcessingStatus7Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus23Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AckdAccptd"`
	PdgPrcg    PendingProcessingStatus13Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PdgPrcg"`
	Dnd        DeniedStatus15Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Dnd"`
	Rjctd      RejectionStatus18Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rjctd"`
	Rprd       RepairStatus13Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rprd"`
	Modfd      ModificationStatus4Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Modfd"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type ModificationReason4 struct {
	Cd          ModificationReason4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type ModificationReason4Choice struct {
	Cd    ModifiedStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type ModificationStatus4Choice struct {
	NoSpcfdRsn NoReasonCode          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []ModificationReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn,omitempty"`
}

// May be one of MDBY, OTHR
type ModifiedStatusReason1Code string

// May be one of NORE
type NoReasonCode string

type OrganisationIdentification7 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 LastPgInd"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PrtryId"`
}

type PartyIdentification98 struct {
	Id  PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 LEI,omitempty"`
}

type PendingProcessingReason10 struct {
	Cd          PendingProcessingReason12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type PendingProcessingReason10Choice struct {
	Cd    PendingProcessingReason2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type PendingProcessingReason12Choice struct {
	Cd    PendingProcessingReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

// May be one of ADEA, CAIS, DOCY, NOFX, BLOC, MUNO, GLOB, YCOL, COLL, FLIM, NEXT, LACK, LALO, MONY, MINO, OTHR, DENO, LIQU, CERT, CSDH, CVAL, CDEL, CDLR, CDAC, INBC
type PendingProcessingReason2Code string

// May be one of ADEA, BLOC, MUNO, NEXT, MINO, OTHR, DENO, CERT
type PendingProcessingReason3Code string

type PendingProcessingReason8 struct {
	Cd          PendingProcessingReason10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type PendingProcessingStatus11Choice struct {
	NoSpcfdRsn NoReasonCode               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []PendingProcessingReason8 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type PendingProcessingStatus13Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []PendingProcessingReason10 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type PendingReason15 struct {
	Cd          PendingReason27Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type PendingReason16 struct {
	Cd          PendingReason28Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type PendingReason17 struct {
	Cd          PendingReason30Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type PendingReason27Choice struct {
	Cd    PendingReason2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type PendingReason28Choice struct {
	Cd    PendingReason6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

// May be one of AWMO, ADEA, CAIS, REFU, AWSH, PHSE, TAMM, DOCY, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, PART, NMAS, NOFX, CMON, YCOL, COLL, DEPO, FLIM, INCA, LINK, FUTU, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, IAAD, OTHR, PHCK, BENO, BOTH, CLHT, DENO, DISA, DKNY, FROZ, LAAW, LATE, LIQU, PRCY, REGT, SETS, CERT, PRSY, CSDH, CVAL, CDLR, INBC
type PendingReason2Code string

type PendingReason30Choice struct {
	Cd    PendingReason9Code      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

// May be one of ADEA, CONF, OTHR, CDRG, CDCY, CDRE
type PendingReason6Code string

// May be one of ADEA, CONF, OTHR, CDRG, CDCY, CDRE, CDAC, INBC
type PendingReason9Code string

type PendingStatus37Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []PendingReason15 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type PendingStatus38Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []PendingReason16 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type PendingStatus39Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []PendingReason17 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type ProcessingStatus49Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AckdAccptd"`
	PdgPrcg    PendingProcessingStatus11Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PdgPrcg"`
	Rjctd      RejectionStatus17Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rjctd"`
	Rpr        RepairStatus12Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rpr"`
	Canc       CancellationStatus14Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Canc"`
	PdgCxl     PendingStatus38Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PdgCxl"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
	CxlReqd    ProprietaryReason4                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 CxlReqd"`
	ModReqd    ProprietaryReason4                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 ModReqd"`
}

type ProcessingStatus53Choice struct {
	PdgCxl     PendingStatus39Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PdgCxl"`
	Rjctd      RejectionOrRepairStatus30Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rjctd"`
	Rpr        RejectionOrRepairStatus31Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rpr"`
	AckdAccptd AcknowledgedAcceptedStatus24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AckdAccptd"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
	Dnd        DeniedStatus16Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Dnd"`
	Canc       CancellationStatus15Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Canc"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PrtryRsn,omitempty"`
}

type RejectionAndRepairReason24Choice struct {
	Cd    RejectionReason32Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type RejectionAndRepairReason25Choice struct {
	Cd    RejectionReason27Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type RejectionOrRepairReason24 struct {
	Cd          RejectionAndRepairReason24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairReason25 struct {
	Cd          RejectionAndRepairReason25Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus30Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason24 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type RejectionOrRepairStatus31Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason25 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type RejectionReason23Choice struct {
	Cd    RejectionReason30Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type RejectionReason24Choice struct {
	Cd    RejectionReason31Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type RejectionReason25 struct {
	Cd          RejectionReason23Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type RejectionReason26 struct {
	Cd          RejectionReason24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, LATE, SAFE, NRGM, NRGN, OTHR, REFE, INVM, INVL
type RejectionReason27Code string

// May be one of SAFE, DQUA, ADEA, DSEC, LATE, CASH, DDEA, DTRD, PLCE, RTGS, NCRR, PHYS, REFE, DMON, MINO, BATC, MUNO, TXST, SETS, IIND, CAEV, CASY, DDAT, SETR, SDUT, INPS, OTHR, ICUS, ICAG, DEPT, IEXE, INVB, INVL, INVN, VALR
type RejectionReason30Code string

// May be one of SAFE, DQUA, ADEA, DSEC, LATE, CASH, DDEA, DTRD, PLCE, RTGS, NCRR, PHYS, REFE, DMON, MINO, BATC, MUNO, TXST, SETS, IIND, CAEV, CASY, DDAT, SETR, SDUT, INPS, OTHR, ICUS, ICAG, DEPT, IEXE, INVL, INVB, INVN, VALR
type RejectionReason31Code string

// May be one of SAFE, ADEA, LATE, NRGN, REFE, NRGM, OTHR
type RejectionReason32Code string

type RejectionStatus17Choice struct {
	NoSpcfdRsn NoReasonCode        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []RejectionReason25 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type RejectionStatus18Choice struct {
	NoSpcfdRsn NoReasonCode        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []RejectionReason26 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type RepairReason10Choice struct {
	Cd    RepairReason4Code       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

// May be one of BATC, CAEV, CASH, CASY, DDAT, DDEA, DMON, DQUA, DSEC, DTRD, IIND, MINO, MUNO, NCRR, PHYS, PLCE, REFE, RTGS, SAFE, SETR, SETS, TXST, INPS, SDUT, OTHR, IEXE, ICAG, DEPT, ICUS
type RepairReason4Code string

type RepairReason8 struct {
	Cd          RepairReason10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type RepairReason9 struct {
	Cd          RepairReason10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type RepairStatus12Choice struct {
	NoSpcfdRsn NoReasonCode    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []RepairReason8 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type RepairStatus13Choice struct {
	NoSpcfdRsn NoReasonCode    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []RepairReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
}

type SecuritiesAccount24 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Nm,omitempty"`
}

type SecuritiesSettlementTransactionAuditTrailReportV03 struct {
	Pgntn     Pagination                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Pgntn"`
	QryRef    Identification14             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 QryRef,omitempty"`
	TxId      TransactionIdentifications29 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 TxId,omitempty"`
	SfkpgAcct SecuritiesAccount24          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 SfkpgAcct"`
	AcctOwnr  PartyIdentification98        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AcctOwnr,omitempty"`
	StsTrl    []StatusTrail6               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 StsTrl,omitempty"`
}

type SettlementStatus17Choice struct {
	Pdg   PendingStatus37Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Pdg"`
	Flng  FailingStatus10Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Flng"`
	Prtry ProprietaryStatusAndReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type StatusTrail6 struct {
	StsDt         ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 StsDt"`
	SndgOrgId     OrganisationIdentification7         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 SndgOrgId,omitempty"`
	UsrId         Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 UsrId,omitempty"`
	PrcgSts       ProcessingStatus49Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PrcgSts,omitempty"`
	IfrrdMtchgSts MatchingStatus25Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 IfrrdMtchgSts,omitempty"`
	MtchgSts      MatchingStatus25Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 MtchgSts,omitempty"`
	SttlmSts      SettlementStatus17Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 SttlmSts,omitempty"`
	ModPrcgSts    ModificationProcessingStatus7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 ModPrcgSts,omitempty"`
	CxlSts        ProcessingStatus53Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 CxlSts,omitempty"`
	Sttld         ProprietaryReason4                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Sttld,omitempty"`
	SplmtryData   []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TransactionIdentifications29 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 PrcrTxId,omitempty"`
}

// May be one of ADEA, ACRU, IIND, CPCA, CLAT, NCRR, DDEA, DSEC, DQUA, LEOG, LATE, MIME, CMIS, NMAS, DTRA, OTHR, FRAP, PHYS, INPS, PLCE, PODU, DEPT, ICAG, ICUS, IEXE, REGD, RTGS, SAFE, DMON, SETS, SETR, TXST, DTRD, DELN, UNBR, DDAT, DMCT, DCMX
type UnmatchedReason12Code string

type UnmatchedReason16 struct {
	Cd          UnmatchedReason23Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 AddtlRsnInf,omitempty"`
}

type UnmatchedReason23Choice struct {
	Cd    UnmatchedReason12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Prtry"`
}

type UnmatchedStatus17Choice struct {
	NoSpcfdRsn NoReasonCode        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 NoSpcfdRsn"`
	Rsn        []UnmatchedReason16 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Rsn"`
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
