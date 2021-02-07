// Code generated by main. DO NOT EDIT.

package caaa_001_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorAuthorisationRequest2 struct {
	Envt  CardPaymentEnvironment9  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Envt"`
	Cntxt CardPaymentContext1      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Cntxt"`
	Tx    CardPaymentTransaction11 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Tx"`
}

type AcceptorAuthorisationRequestV02 struct {
	Hdr        Header1                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Hdr"`
	AuthstnReq AcceptorAuthorisationRequest2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AuthstnReq"`
	SctyTrlr   ContentInformationType6       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SctyTrlr"`
}

type Acquirer2 struct {
	Id         GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id,omitempty"`
	ParamsVrsn Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 ParamsVrsn"`
}

type AddressVerification1 struct {
	AdrDgts    Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AdrDgts,omitempty"`
	PstlCdDgts Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PstlCdDgts,omitempty"`
}

// May be one of EA2C, E3DC, DKPT, DKP9, UKPT, UKA1
type Algorithm2Code string

// May be one of MACC, MCCS, CMA1, CMD1
type Algorithm3Code string

// May be one of HS25, HS38, HS51
type Algorithm5Code string

// May be one of EA2C, E3DC
type Algorithm6Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification2 struct {
	Algo  Algorithm2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Param,omitempty"`
}

type AlgorithmIdentification3 struct {
	Algo  Algorithm3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Param,omitempty"`
}

type AlgorithmIdentification6 struct {
	Algo  Algorithm6Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Param,omitempty"`
}

type AlgorithmIdentification7 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Algo"`
	Param Parameter2     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Param,omitempty"`
}

type AlgorithmIdentification8 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Algo"`
	Param Parameter3     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Param,omitempty"`
}

// May be one of ATTD, SATT, UATT
type AttendanceContext1Code string

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData2 struct {
	Vrsn        float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Vrsn,omitempty"`
	Rcpt        []Recipient2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Rcpt"`
	MACAlgo     AlgorithmIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 MACAlgo"`
	NcpsltdCntt EncapsulatedContent1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 NcpsltdCntt"`
	MAC         Max35Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 MAC"`
}

// May be one of ICCD, AGNT, MERC
type AuthenticationEntity1Code string

// May be one of BYPS, MANU, MERC, FPIN, NPIN, PPSG, PSWD, SCRT, SCNL, SNCT, CPSG, UKNW
type AuthenticationMethod2Code string

// May be one of PRST, BYPS, UNRD, NCSC
type CSCManagement1Code string

// May be one of DFLT, SVNG, CHCK, CRDT, UVRL, INVS, EPRS
type CardAccountType1Code string

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL
type CardDataReading1Code string

type CardPaymentContext1 struct {
	PmtCntxt  PaymentContext1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PmtCntxt"`
	SaleCntxt SaleContext1    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SaleCntxt,omitempty"`
}

type CardPaymentEnvironment9 struct {
	Acqrr   Acquirer2           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Acqrr,omitempty"`
	Mrchnt  Organisation8       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Mrchnt,omitempty"`
	POI     PointOfInteraction2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 POI"`
	Card    PaymentCard5        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Card"`
	Crdhldr Cardholder3         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Crdhldr,omitempty"`
}

// May be one of AGGR, DCCV, GRTT, INSP, LOYT, NRES, PUCO, RECP, SOAF, UNAF, VCAU
type CardPaymentServiceType2Code string

// May be one of IRES, URES, PRES, ARES, FREC, RREC
type CardPaymentServiceType3Code string

// May be one of TP2P, TP2B, BALC, CACT, CRDP, CAFT, CAVR, CSHW, CSHB, CSHD, DEFR, LOAD, ORCR, PINC, QUCH, RFND, RESA, VALC
type CardPaymentServiceType4Code string

type CardPaymentTransaction11 struct {
	TxCaptr      bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxCaptr"`
	TxTp         CardPaymentServiceType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxTp"`
	AddtlSvc     []CardPaymentServiceType2Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AddtlSvc,omitempty"`
	SvcAttr      CardPaymentServiceType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SvcAttr,omitempty"`
	MrchntCtgyCd Min3Max4Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 MrchntCtgyCd"`
	TxId         TransactionIdentifier1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxId"`
	OrgnlTx      CardPaymentTransaction17       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 OrgnlTx,omitempty"`
	InitrTxId    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 InitrTxId,omitempty"`
	RcncltnId    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RcncltnId,omitempty"`
	TxDtls       CardPaymentTransactionDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxDtls"`
	AddtlTxData  []Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AddtlTxData,omitempty"`
}

type CardPaymentTransaction17 struct {
	TxId      TransactionIdentifier1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxId"`
	POIId     GenericIdentification32       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 POIId,omitempty"`
	InitrTxId Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 InitrTxId,omitempty"`
	RcptTxId  Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RcptTxId,omitempty"`
	TxTp      CardPaymentServiceType4Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxTp"`
	AddtlSvc  []CardPaymentServiceType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AddtlSvc,omitempty"`
	SvcAttr   CardPaymentServiceType3Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SvcAttr,omitempty"`
	TxRslt    CardPaymentTransactionResult1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxRslt,omitempty"`
}

type CardPaymentTransactionDetails1 struct {
	Ccy            CurrencyCode          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Ccy"`
	TtlAmt         float64               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TtlAmt"`
	AmtQlfr        TypeOfAmount1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AmtQlfr,omitempty"`
	DtldAmt        []DetailedAmount1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 DtldAmt,omitempty"`
	VldtyDt        ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 VldtyDt,omitempty"`
	OnLineRsn      OnLineReason1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 OnLineRsn,omitempty"`
	UattnddLvlCtgy Max35NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 UattnddLvlCtgy,omitempty"`
	AcctTp         CardAccountType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AcctTp,omitempty"`
	RcrngTx        RecurringTransaction1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RcrngTx,omitempty"`
	Pdct           []Product1            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Pdct,omitempty"`
	ICCRltdData    Max10000Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 ICCRltdData,omitempty"`
}

type CardPaymentTransactionResult1 struct {
	AuthstnNtty   GenericIdentification33 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AuthstnNtty,omitempty"`
	RspnToAuthstn ResponseType1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RspnToAuthstn"`
	AuthstnCd     Min6Max8Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AuthstnCd,omitempty"`
}

type CardSecurityInformation1 struct {
	CSCMgmt CSCManagement1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CSCMgmt"`
	CSCVal  Min3Max4NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CSCVal,omitempty"`
}

type Cardholder3 struct {
	Id        []CardholderIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id,omitempty"`
	Nm        Max45Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Nm,omitempty"`
	Lang      ISO2ALanguageCode           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Lang,omitempty"`
	Authntcn  []CardholderAuthentication3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Authntcn,omitempty"`
	AdrVrfctn AddressVerification1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AdrVrfctn,omitempty"`
	PrsnlData Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PrsnlData,omitempty"`
}

type CardholderAuthentication3 struct {
	AuthntcnMtd       AuthenticationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AuthntcnMtd"`
	AuthntcnNtty      AuthenticationEntity1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AuthntcnNtty,omitempty"`
	AuthntcnVal       Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AuthntcnVal,omitempty"`
	CrdhldrOnLinePIN  OnLinePIN2                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CrdhldrOnLinePIN,omitempty"`
	AuthntcnColltnInd Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AuthntcnColltnInd,omitempty"`
}

type CardholderIdentification1 struct {
	CrdhldrIdVal Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CrdhldrIdVal"`
	CrdhldrIdTp  PersonIdentificationType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CrdhldrIdTp"`
}

// May be one of MNSG, NPIN, FCPN, FEPN, FDSG, FBIO, MNVR, FBIG, APKI, PKIS, CHDT, SCEC
type CardholderVerificationCapability1Code string

type CertificateIdentifier1 struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 IssrAndSrlNb"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RltvDstngshdNm"`
}

type CommunicationCharacteristics1 struct {
	ComTp   POICommunicationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 ComTp"`
	RmotPty PartyType7Code            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RmotPty"`
	Actv    bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Actv"`
}

type ContentInformationType5 struct {
	CnttTp     ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CnttTp"`
	EnvlpdData EnvelopedData2   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 EnvlpdData"`
}

type ContentInformationType6 struct {
	CnttTp       ContentType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CnttTp"`
	AuthntcdData []AuthenticatedData2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AuthntcdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, ECRP, AUTH
type ContentType1Code string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DetailedAmount1 struct {
	Tp  TypeOfAmount2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Tp"`
	Val float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Val"`
}

type DisplayCapabilities1 struct {
	DispTp    UserInterface2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 DispTp"`
	NbOfLines Max3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 NbOfLines"`
	LineWidth Max3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 LineWidth"`
}

type Document struct {
	AccptrAuthstnReq AcceptorAuthorisationRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AccptrAuthstnReq"`
}

type EncapsulatedContent1 struct {
	CnttTp ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CnttTp"`
	Cntt   Max10000Binary   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Cntt,omitempty"`
}

type EncryptedContent2 struct {
	CnttTp         ContentType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CnttNcrptnAlgo"`
	NcrptdData     Max10000Binary           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 NcrptdData"`
}

type EnvelopedData2 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Vrsn,omitempty"`
	Rcpt       []Recipient2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Rcpt"`
	NcrptdCntt EncryptedContent2  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 NcrptdCntt"`
}

// May be no more than 10 items long
type Exact10Text string

// Must match the pattern [0-9]
type Exact1NumericText string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// May be one of DAIL, MNTH, YEAR
type Frequency4Code string

type GenericIdentification31 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Tp"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 ShrtNm,omitempty"`
}

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 ShrtNm,omitempty"`
}

type GenericIdentification33 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id,omitempty"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Tp"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 ShrtNm,omitempty"`
}

type GenericIdentification48 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id"`
	Vrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Vrsn"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Issr"`
}

type Header1 struct {
	MsgFctn    MessageFunction1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PrtcolVrsn"`
	XchgId     Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CreDtTm"`
	InitgPty   GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 InitgPty"`
	RcptPty    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RcptPty,omitempty"`
	Tracblt    []Traceability1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Tracblt,omitempty"`
}

// Must match the pattern [a-z]{2,2}
type ISO2ALanguageCode string

// Must match the pattern [0-9]{3,3}
type ISO3NumericCountryCode string

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

type ISOYearMonth time.Time

func (t *ISOYearMonth) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(t).UnmarshalText(text)
}
func (t ISOYearMonth) MarshalText() ([]byte, error) {
	return xsdGYearMonth(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SrlNb"`
}

type KEK2 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 NcrptdKey"`
}

type KEKIdentifier1 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 KeyId"`
	KeyVrsn   Exact10Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 KeyVrsn"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 DerivtnId,omitempty"`
}

type KeyTransport2 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Vrsn"`
	RcptId        CertificateIdentifier1   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 NcrptdKey"`
}

// May be one of FIXD, ABRD, NMDC, MOTO, HOME
type LocationCategory1Code string

type Max10000Binary []byte

func (t *Max10000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,2}
type Max2NumericText string

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{1,35}
type Max35NumericText string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max3Text string

// Must be at least 1 items long
type Max40Text string

// Must be at least 1 items long
type Max45Text string

type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

type MemoryCharacteristics1 struct {
	Id     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id"`
	TtlSz  float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TtlSz"`
	FreeSz float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 FreeSz"`
	Unit   MemoryUnit1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Unit"`
}

// May be one of BYTE, EXAB, GIGA, KILO, MEGA, PETA, TERA
type MemoryUnit1Code string

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT
type MessageFunction1Code string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

// Must match the pattern [0-9]{3,4}
type Min3Max4NumericText string

// Must be at least 3 items long
type Min3Max4Text string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 6 items long
type Min6Max8Text string

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

// May be one of OFLN, ONLN, SMON
type OnLineCapability1Code string

type OnLinePIN2 struct {
	NcrptdPINBlck ContentInformationType5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 NcrptdPINBlck"`
	PINFrmt       PINFormat2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PINFrmt"`
	AddtlInpt     Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AddtlInpt,omitempty"`
}

// May be one of RNDM, ICCF, MERF, TRMF, ISSF, FRLT, EXFL, TAMT, CBIN, UBIN, CPAN, FLOW, CRCY
type OnLineReason1Code string

type Organisation8 struct {
	Id        GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id,omitempty"`
	CmonNm    Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CmonNm,omitempty"`
	LctnCtgy  LocationCategory1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 LctnCtgy,omitempty"`
	Adr       Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Adr,omitempty"`
	CtryCd    ISO3NumericCountryCode  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CtryCd,omitempty"`
	SchmeData Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SchmeData,omitempty"`
}

// May be one of ISO0, ISO1, ISO2, ISO3, ISO4
type PINFormat2Code string

// May be one of BLTH, ETHR, GPRS, GSMF, PSTN, RS23, USBD, USBH, WIFI
type POICommunicationType1Code string

// May be one of APPL, CERT, EVAL
type POIComponentAssessment1Code string

// May be one of WAIT, OUTD, OPER, DACT
type POIComponentStatus1Code string

// May be one of AQPP, APPR, TLPR, SCPR, SERV, TERM, DVCE, SECM, APLI, EMVK, EMVO, MDWR, DRVR, OPST, MRPR
type POIComponentType3Code string

type Parameter1 struct {
	InitlstnVctr Max500Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 InitlstnVctr,omitempty"`
}

type Parameter2 struct {
	DgstAlgo     Algorithm5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter3 struct {
	DgstAlgo Algorithm5Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 DgstAlgo,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

// May be one of ACQR, ITAG, PCPT, TMGT, SALE
type PartyType7Code string

type PaymentCard5 struct {
	PrtctdCardData ContentInformationType5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData1          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PlainCardData,omitempty"`
	CardCtryCd     Max3Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CardCtryCd,omitempty"`
	CardPdctPrfl   Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CardPdctPrfl,omitempty"`
	CardBrnd       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CardBrnd,omitempty"`
	AddtlCardData  Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AddtlCardData,omitempty"`
}

type PaymentContext1 struct {
	CardPres       bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CardPres,omitempty"`
	CrdhldrPres    bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CrdhldrPres,omitempty"`
	OnLineCntxt    bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 OnLineCntxt,omitempty"`
	AttndncCntxt   AttendanceContext1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AttndncCntxt,omitempty"`
	TxEnvt         TransactionEnvironment1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxEnvt,omitempty"`
	TxChanl        TransactionChannel1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxChanl,omitempty"`
	AttndntMsgCpbl bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AttndntMsgCpbl,omitempty"`
	AttndntLang    ISO2ALanguageCode           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AttndntLang,omitempty"`
	CardDataNtryMd CardDataReading1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CardDataNtryMd"`
	FllbckInd      bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 FllbckInd,omitempty"`
}

// May be one of PASS, DRLC, EEID, DRVR
type PersonIdentificationType4Code string

type PlainCardData1 struct {
	PAN        Min8Max28NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PAN"`
	CardSeqNb  Min2Max3NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CardSeqNb,omitempty"`
	FctvDt     ISOYearMonth             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 FctvDt,omitempty"`
	XpryDt     ISOYearMonth             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 XpryDt"`
	SvcCd      Exact3NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SvcCd,omitempty"`
	TrckData   []TrackData1             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TrckData,omitempty"`
	CardSctyCd CardSecurityInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CardSctyCd,omitempty"`
}

type PointOfInteraction2 struct {
	Id       GenericIdentification32         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id"`
	SysNm    Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SysNm,omitempty"`
	GrpId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 GrpId,omitempty"`
	Cpblties PointOfInteractionCapabilities1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Cpblties,omitempty"`
	Cmpnt    []PointOfInteractionComponent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Cmpnt,omitempty"`
}

type PointOfInteractionCapabilities1 struct {
	CardRdngCpblties      []CardDataReading1Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CardRdngCpblties,omitempty"`
	CrdhldrVrfctnCpblties []CardholderVerificationCapability1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CrdhldrVrfctnCpblties,omitempty"`
	OnLineCpblties        OnLineCapability1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 OnLineCpblties,omitempty"`
	DispCpblties          []DisplayCapabilities1                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 DispCpblties,omitempty"`
	PrtLineWidth          Max3NumericText                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PrtLineWidth,omitempty"`
}

type PointOfInteractionComponent3 struct {
	Tp       POIComponentType3Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Tp"`
	Id       PointOfInteractionComponentIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id"`
	Sts      PointOfInteractionComponentStatus1          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Sts,omitempty"`
	StdCmplc []GenericIdentification48                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 StdCmplc,omitempty"`
	Chrtcs   PointOfInteractionComponentCharacteristics1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Chrtcs,omitempty"`
	Assmnt   []PointOfInteractionComponentAssessment1    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Assmnt,omitempty"`
}

type PointOfInteractionComponentAssessment1 struct {
	Tp      POIComponentAssessment1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Tp"`
	Assgnr  []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Assgnr"`
	DlvryDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 DlvryDt,omitempty"`
	XprtnDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 XprtnDt,omitempty"`
	Nb      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Nb"`
}

type PointOfInteractionComponentCharacteristics1 struct {
	Mmry           []MemoryCharacteristics1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Mmry,omitempty"`
	Com            []CommunicationCharacteristics1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Com,omitempty"`
	SctyAccsMdls   float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SctyAccsMdls,omitempty"`
	SbcbrIdntyMdls float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SbcbrIdntyMdls,omitempty"`
	KeyChckVal     Max35Binary                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 KeyChckVal,omitempty"`
}

type PointOfInteractionComponentIdentification1 struct {
	ItmNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 ItmNb,omitempty"`
	PrvdrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PrvdrId,omitempty"`
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Id,omitempty"`
	SrlNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SrlNb,omitempty"`
}

type PointOfInteractionComponentStatus1 struct {
	VrsnNb Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 VrsnNb,omitempty"`
	Sts    POIComponentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Sts,omitempty"`
}

type Product1 struct {
	PdctCd       Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PdctCd"`
	UnitOfMeasr  UnitOfMeasure1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 UnitOfMeasr,omitempty"`
	PdctQty      float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PdctQty,omitempty"`
	UnitPric     float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 UnitPric,omitempty"`
	PdctAmt      float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PdctAmt"`
	TaxTp        Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TaxTp,omitempty"`
	AddtlPdctInf Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AddtlPdctInf,omitempty"`
}

type Recipient2Choice struct {
	KeyTrnsprt KeyTransport2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 KeyTrnsprt,omitempty"`
	KEK        KEK2          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 KEK,omitempty"`
}

type RecurringTransaction1 struct {
	SeqNb       Max2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SeqNb"`
	PrdUnit     Frequency4Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 PrdUnit"`
	InstlmtPrd  float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 InstlmtPrd"`
	TtlNbOfPmts float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TtlNbOfPmts"`
	IntrstChrgs float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 IntrstChrgs,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AttrVal"`
}

// May be one of DECL, APPR, PART, TECH
type Response1Code string

type ResponseType1 struct {
	Rspn    Response1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Rspn"`
	RspnRsn Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RspnRsn,omitempty"`
}

type SaleContext1 struct {
	SaleId        Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SaleId,omitempty"`
	SaleRefNb     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SaleRefNb,omitempty"`
	SaleRcncltnId Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 SaleRcncltnId,omitempty"`
	CshrId        Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 CshrId,omitempty"`
	ShftNb        Max2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 ShftNb,omitempty"`
	AddtlSaleData Max70Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 AddtlSaleData,omitempty"`
}

type Traceability1 struct {
	RlayId      GenericIdentification31 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TracDtTmOut"`
}

type TrackData1 struct {
	TrckNb  Exact1NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TrckNb,omitempty"`
	TrckVal Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TrckVal"`
}

// May be one of MAIL, TLPH, ECOM, TVPY
type TransactionChannel1Code string

// May be one of MERC, PRIV, PUBL
type TransactionEnvironment1Code string

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 TxRef"`
}

// May be one of ACTL, ESTM, MAXI, DFLT, RPLT
type TypeOfAmount1Code string

// May be one of CSHB, GRTY, FEES, RBTS, VATX
type TypeOfAmount2Code string

// May be one of PIEC, TONS, FOOT, GBGA, USGA, GRAM, INCH, KILO, PUND, METR, CMET, MMET, LITR, CELI, MILI, GBOU, USOU, GBQA, USQA, GBPI, USPI, MILE, KMET, YARD, SQKI, HECT, ARES, SMET, SCMT, SMIL, SQMI, SQYA, SQFO, SQIN, ACRE
type UnitOfMeasure1Code string

// May be one of MDSP, CDSP
type UserInterface2Code string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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

type xsdGYearMonth time.Time

func (t *xsdGYearMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYearMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
