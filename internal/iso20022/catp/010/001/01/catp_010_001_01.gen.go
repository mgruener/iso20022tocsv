// Code generated by main. DO NOT EDIT.

package catp_010_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// ATMConfigurationParameter1
//
// Configuration parameter version of the ATM.
type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Vrsn"`
}

// ATMContext7
//
// Context in which the transaction is performed.
type ATMContext7 struct {
	SsnRef Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SsnRef,omitempty"`
	Svc    ATMService8 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Svc"`
}

// ATMCustomer1
//
// Customer involved in a withdrawal transaction.
type ATMCustomer1 struct {
	Prfl         ATMCustomerProfile1              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Prfl,omitempty"`
	SelctdLang   string                           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SelctdLang,omitempty"`
	Authntcn     []CardholderAuthentication8      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Authntcn"`
	AuthntcnRslt []TransactionVerificationResult5 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AuthntcnRslt,omitempty"`
}

// ATMCustomerProfile1
//
// Profile of the customer selected by an ATM.
type ATMCustomerProfile1 struct {
	RtrvlMd ATMCustomerProfile1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 RtrvlMd"`
	PrflRef Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PrflRef,omitempty"`
	CstmrId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CstmrId,omitempty"`
}

// May be one of CRDF, OREQ, PREQ
type ATMCustomerProfile1Code string

// ATMEnvironment1
//
// Environment of the withdrawal transaction.
type ATMEnvironment1 struct {
	Acqrr    Acquirer7               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Acqrr,omitempty"`
	ATMMgrId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ATMMgrId,omitempty"`
	HstgNtty TerminalHosting1        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 HstgNtty,omitempty"`
	ATM      AutomatedTellerMachine1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ATM"`
	Cstmr    ATMCustomer1            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Cstmr"`
	Card     PaymentCard16           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Card,omitempty"`
}

// ATMEquipment1
//
// ATM terminal equipment.
type ATMEquipment1 struct {
	Manfctr    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Manfctr,omitempty"`
	Mdl        Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Mdl,omitempty"`
	SrlNb      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SrlNb,omitempty"`
	ApplPrvdr  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ApplPrvdr,omitempty"`
	ApplNm     Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ApplNm,omitempty"`
	ApplVrsn   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ApplVrsn,omitempty"`
	ApprvlNb   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ApprvlNb,omitempty"`
	CfgtnParam []ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CfgtnParam,omitempty"`
}

// ATMMessageFunction1
//
// Identifies the type of process related to an ATM message.
type ATMMessageFunction1 struct {
	Fctn     MessageFunction7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Fctn"`
	ATMSvcCd Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 HstSvcCd,omitempty"`
}

// ATMPINManagementRequest1
//
// Information related to the request of a PIN management from an ATM.
type ATMPINManagementRequest1 struct {
	Envt  ATMEnvironment1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Envt"`
	Cntxt ATMContext7     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Cntxt"`
	Tx    ATMTransaction9 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Tx"`
}

// ATMPINManagementRequestV01
//
// The ATMPINManagementRequest message is sent by an ATM to an ATM manager to request an operation on the cardholder PIN.
type ATMPINManagementRequestV01 struct {
	Hdr                 Header20                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Hdr"`
	PrtctdATMPINMgmtReq ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PrtctdATMPINMgmtReq,omitempty"`
	ATMPINMgmtReq       ATMPINManagementRequest1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ATMPINMgmtReq,omitempty"`
	SctyTrlr            ContentInformationType15 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SctyTrlr,omitempty"`
}

// ATMService8
//
// Context in which the transaction is performed.
type ATMService8 struct {
	SvcRef   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SvcRef,omitempty"`
	ATMSvcCd Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ATMSvcCd,omitempty"`
	SvcTp    ATMServiceType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SvcTp"`
}

// May be one of PINC, PINR, PINU
type ATMServiceType5Code string

// ATMTransaction9
//
// Transaction for which the service is requested.
type ATMTransaction9 struct {
	TxId          TransactionIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 TxId"`
	RcncltnId     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 RcncltnId,omitempty"`
	CrdhldrNewPIN OnLinePIN5             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CrdhldrNewPIN,omitempty"`
	ICCRltdData   Max10000Binary         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ICCRltdData,omitempty"`
}

// Acquirer7
//
// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Brnch,omitempty"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

// AlgorithmIdentification11
//
// Cryptographic algorithms and parameters for the protection of transported keys by an asymmetric key.
type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Param,omitempty"`
}

// AlgorithmIdentification12
//
// Mask generator function cryptographic algorithm and parameters.
type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Param,omitempty"`
}

// AlgorithmIdentification13
//
// Cryptographic algorithm and parameters for the protection of the transported key.
type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Param,omitempty"`
}

// AlgorithmIdentification14
//
// Cryptographic algorithm and parameters for encryptions with a symmetric cryptographic key.
type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Param,omitempty"`
}

// AlgorithmIdentification15
//
// Identification of a cryptographic algorithm and parameters for the MAC computation.
type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

// AuthenticatedData4
//
// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 MAC"`
}

// May be one of ICCD, AGNT, MERC, ACQR, ISSR, TRML
type AuthenticationEntity2Code string

// May be one of TOKA, BIOM, MOBL, OTHR, FPIN, NPIN, PSWD, SCRT, SCNL
type AuthenticationMethod7Code string

// AutomatedTellerMachine1
//
// ATM information.
type AutomatedTellerMachine1 struct {
	Id       Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Id"`
	AddtlId  Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AddtlId,omitempty"`
	SeqNb    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SeqNb,omitempty"`
	BaseCcy  ActiveCurrencyCode              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 BaseCcy"`
	Lctn     PostalAddress17                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Lctn,omitempty"`
	LctnCtgy TransactionEnvironment2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 LctnCtgy,omitempty"`
	Cpblties PointOfInteractionCapabilities5 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Cpblties,omitempty"`
	Eqpmnt   ATMEquipment1                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Eqpmnt,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL
type CardDataReading1Code string

// May be one of ECTL, CICC, MGST, CTLS
type CardDataReading4Code string

// May be one of FFLB, SFLB, NFLB
type CardFallback1Code string

// CardholderAuthentication8
//
// Data related to the authentication of the card and the cardholder.
type CardholderAuthentication8 struct {
	AuthntcnMtd       AuthenticationMethod7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AuthntcnMtd"`
	TknReqd           bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 TknReqd,omitempty"`
	AuthntcnVal       Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AuthntcnVal,omitempty"`
	PrtctdAuthntcnVal ContentInformationType10  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PrtctdAuthntcnVal,omitempty"`
	CrdhldrOnLinePIN  OnLinePIN5                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CrdhldrOnLinePIN,omitempty"`
}

// May be one of NPIN, FCPN, FEPN, FDSG, FBIO, FBIG, PKIS, PCOD
type CardholderVerificationCapability3Code string

// CertificateIssuer1
//
// Certificate issuer name (see X.509).
type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 RltvDstngshdNm"`
}

// ContentInformationType10
//
// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 EnvlpdData"`
}

// ContentInformationType15
//
// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType15 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AuthntcdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type Document struct {
	ATMPINMgmtReq ATMPINManagementRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ATMPINMgmtReq"`
}

// EncapsulatedContent3
//
// Data to authenticate.
type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Cntt,omitempty"`
}

// EncryptedContent3
//
// Encrypted data with an encryption key.
type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

// EnvelopedData4
//
// Encrypted data with encryption key.
type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 NcrptdCntt,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{3}
type Exact3AlphaNumericText string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// GenericIdentification77
//
// Identification of an entity.
type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ShrtNm,omitempty"`
}

// GeographicCoordinates1
//
// Location on the Earth specified by two numbers representing vertical and horizontal position.
type GeographicCoordinates1 struct {
	Lat  Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Lat"`
	Long Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Long"`
}

// GeographicLocation1Choice
//
// Geographic location of the ATM specified by geographic coordinates or UTM coordinates.
type GeographicLocation1Choice struct {
	GeogcCordints GeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 GeogcCordints"`
	UTMCordints   UTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 UTMCordints"`
}

// Header20
//
// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header20 struct {
	MsgFctn    ATMMessageFunction1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 MsgFctn"`
	PrtcolVrsn Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PrtcolVrsn"`
	XchgId     Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 XchgId"`
	CreDtTm    ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CreDtTm"`
	InitgPty   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 InitgPty"`
	RcptPty    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 RcptPty,omitempty"`
	PrcStat    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PrcStat,omitempty"`
	Tracblt    []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Tracblt,omitempty"`
}

// ISODateTime
//
// A particular point in the progression of time defined by a mandatory date and a mandatory time component, expressed in either UTC time format (YYYY-MM-DDThh:mm:ss.sssZ), local time with UTC offset format (YYYY-MM-DDThh:mm:ss.sss+/-hh:mm), or local time format (YYYY-MM-DDThh:mm:ss.sss). These representations are defined in "XML Schema Part 2: Datatypes Second Edition - W3C Recommendation 28 October 2004" which is aligned with ISO 8601.
// Note on the time format:
// 1) beginning / end of calendar day
// 00:00:00 = the beginning of a calendar day
// 24:00:00 = the end of a calendar day
// 2) fractions of second in time format
// Decimal fractions of seconds may be included. In this case, the involved parties shall agree on the maximum number of digits that are allowed.
type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// IssuerAndSerialNumber1
//
// Certificate issuer name and serial number  (see X.509).
type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SrlNb"`
}

// KEK4
//
// Key encryption key (KEK), using previously distributed symmetric key.
type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 NcrptdKey"`
}

// KEKIdentifier2
//
// Identification of a key encryption key (KEK), using previously distributed symmetric key.
type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 DerivtnId,omitempty"`
}

// KeyTransport4
//
// Key encryption key (KEK), encrypted with a previously distributed asymmetric public key.
type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 NcrptdKey"`
}

// Max10000Binary
//
// Specifies a binary string with a maximum length of 10000 binary bytes.
type Max10000Binary []byte

func (t *Max10000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Max100KBinary
//
// Binary data of 100K maximum.
type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max10Text string

// Max140Binary
//
// Specifies a binary string with a maximum length of 140 binary bytes.
type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Max35Binary
//
// Specifies a binary string with a maximum length of 35 binary bytes.
type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max3Text string

// Must be at least 1 items long
type Max45Text string

// Max5000Binary
//
// Specifies a binary string with a maximum length of 5000 binary bytes.
type Max5000Binary []byte

func (t *Max5000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max5000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Max500Binary
//
// Specifies a binary string with a maximum length of 500 binary bytes.
type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max500Text string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// May be one of BALN, CMPA, CMPD, ACMD, DVCC, DIAQ, DIAP, GSTS, INQQ, INQP, KYAQ, KYAP, PINQ, PINP, RJAQ, RJAP, WITV, WITK, WITQ, WITP, INQC, H2AP, H2AQ, TMOP, CSEC, DSEC, SKSC, SSTS
type MessageFunction7Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

// Min5Max16Binary
//
// Specifies a binary string with a minimum length of 5 bytes, and a maximum length of 16 bytes.
type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

// OnLinePIN5
//
// Encrypted personal identification number (PIN) and related information.
type OnLinePIN5 struct {
	NcrptdPINBlck ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 NcrptdPINBlck"`
	PINFrmt       PINFormat4Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PINFrmt"`
	AddtlInpt     Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AddtlInpt,omitempty"`
}

// May be one of ANSI, BNCM, BKSY, DBLD, DBLC, ECI2, ECI3, EMVS, IBM3, ISO0, ISO1, ISO2, ISO3, ISO4, ISO5, VIS2, VIS3
type PINFormat4Code string

// Parameter4
//
// Parameters of the asymmetric encryption algorithm.
type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 MskGnrtrAlgo,omitempty"`
}

// Parameter5
//
// Parameters associated to a mask generator cryptographic function.
type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 DgstAlgo,omitempty"`
}

// Parameter6
//
// Parameters associated to a cryptographic encryption algorithm.
type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 BPddg,omitempty"`
}

// Parameter7
//
// Parameters associated to the MAC algorithm.
type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 BPddg,omitempty"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

// PaymentCard16
//
// Card performing the withdrawal transaction.
type PaymentCard16 struct {
	CardDataNtryMd CardDataReading1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CardDataNtryMd"`
	FllbckInd      CardFallback1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 FllbckInd,omitempty"`
	PrtctdCardData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData13          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PlainCardData,omitempty"`
	CardCtryCd     Max3Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CardCtryCd,omitempty"`
	CardCcyCd      Exact3AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CardCcyCd,omitempty"`
}

// PlainCardData13
//
// Sensible data associated with the payment card performing the transaction.
type PlainCardData13 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PAN,omitempty"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 XpryDt,omitempty"`
	SvcCd     Exact3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SvcCd,omitempty"`
	Trck1     Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Trck1,omitempty"`
	Trck2     Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Trck2,omitempty"`
	Trck3     Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Trck3,omitempty"`
	CrdhldrNm Max45Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CrdhldrNm,omitempty"`
}

// PointOfInteractionCapabilities5
//
// Capabilities of the ATM terminal.
type PointOfInteractionCapabilities5 struct {
	CardRdData       []CardDataReading4Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CardRdData,omitempty"`
	CardWrtData      []CardDataReading4Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CardWrtData,omitempty"`
	Authntcn         []CardholderVerificationCapability3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Authntcn,omitempty"`
	PINLngthCpblties float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PINLngthCpblties,omitempty"`
	ApprvlCdLngth    float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 ApprvlCdLngth,omitempty"`
	MxScrptLngth     float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 MxScrptLngth,omitempty"`
	CardCaptrCpbl    bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CardCaptrCpbl,omitempty"`
}

// PostalAddress17
//
// Information that locates and identifies a specific address, as defined by postal services or in free format text.
type PostalAddress17 struct {
	AdrLine     []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 BldgNb,omitempty"`
	PstCd       Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 PstCd,omitempty"`
	TwnNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 TwnNm"`
	CtrySubDvsn []Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Ctry"`
	GLctn       GeographicLocation1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 GLctn,omitempty"`
}

// Recipient4Choice
//
// Transport key or key encryption key (KEK) for the recipient.
type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 KeyIdr"`
}

// Recipient5Choice
//
// Identification of a cryptographic asymmetric key.
type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 KeyIdr"`
}

// RelativeDistinguishedName1
//
// Relative distinguished name defined by X.500 and X.509.
type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AttrVal"`
}

// TerminalHosting1
//
// Entity hosting the ATM terminal.
type TerminalHosting1 struct {
	Ctgy TransactionEnvironment3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Ctgy,omitempty"`
	Id   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Id,omitempty"`
}

// Traceability4
//
// Identification of partners involved in exchange from the ATM to the issuer, with the relative timestamp of their exchanges.
type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 TracDtTmOut"`
}

// May be one of PRIV, PUBL
type TransactionEnvironment2Code string

// May be one of BRCH, MERC, OTHR
type TransactionEnvironment3Code string

// TransactionIdentifier1
//
// Identification of the transaction in an unambiguous way.
type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 TxRef"`
}

// TransactionVerificationResult5
//
// Result of performed verifications for the transaction.
type TransactionVerificationResult5 struct {
	Mtd         AuthenticationMethod7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Mtd"`
	VrfctnNtty  AuthenticationEntity2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 VrfctnNtty,omitempty"`
	Rslt        Verification1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Rslt,omitempty"`
	AddtlRslt   Max500Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AddtlRslt,omitempty"`
	AuthntcnTkn Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 AuthntcnTkn,omitempty"`
}

// UTMCoordinates1
//
// Location on the Earth specified by the Universal Transverse Mercator coordinate system, using the WGS84 geodesic system.
type UTMCoordinates1 struct {
	UTMZone    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 UTMZone"`
	UTMEstwrd  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 UTMEstwrd"`
	UTMNrthwrd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 UTMNrthwrd"`
}

// May be one of FAIL, MISS, NOVF, PART, SUCC, ERRR
type Verification1Code string

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
