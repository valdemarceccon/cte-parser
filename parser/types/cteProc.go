package types

import (
	"encoding/xml"
)

// <xs:schema xmlns:ds="http://www.w3.org/2000/09/xmldsig#" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns="http://www.portalfiscal.inf.br/cte" targetNamespace="http://www.portalfiscal.inf.br/cte" elementFormDefault="qualified" attributeFormDefault="unqualified">
// 	<xs:include schemaLocation="cteTiposBasico_v3.00.xsd"/>
// 	<xs:element name="cteProc">
// 		<xs:annotation>
// 			<xs:documentation> CT-e processado</xs:documentation>
// 		</xs:annotation>
// 		<xs:complexType>
// 			<xs:sequence>
// 				<xs:element name="CTe" type="TCTe"/>
// 				<xs:element name="protCTe" type="TProtCTe"/>
// 			</xs:sequence>
// 			<xs:attribute name="versao" type="TVerCTe" use="required"/>
// 			<xs:attribute name="ipTransmissor" type="TIPv4" use="optional"/>
// 		</xs:complexType>
// 	</xs:element>
// </xs:schema>

//ProcCTE CT-e processado
type ProcCTE struct {
	XMLName       xml.Name   `xml:"cteProc"`
	Versao        TVersaoCTe `xml:"versao,attr"`
	IPTransmissor TIPv4      `xml:"ipTransmissor,attr"`
	ProtCTe       TProtCTe   `xml:"protCTe"`
}

// TVersaoCTe Tipo Versão do CT-e - 3.00
type TVersaoCTe string

// TIPv4 Tipo IP versão 4
type TIPv4 string
