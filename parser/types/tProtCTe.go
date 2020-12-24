package types

import "encoding/xml"

// <xs:complexType name="TProtCTe">
// <xs:annotation>
// 	<xs:documentation>Tipo Protocolo de status resultado do processamento da CT-e</xs:documentation>
// </xs:annotation>
// <xs:sequence>
// 	<xs:element name="infProt">
// 		<xs:annotation>
// 			<xs:documentation>Dados do protocolo de status</xs:documentation>
// 		</xs:annotation>
// 		<xs:complexType>
// 			<xs:sequence>
// 				<xs:element name="tpAmb" type="TAmb">
// 					<xs:annotation>
// 						<xs:documentation>Identificação do Ambiente:
// 1 - Produção
// 2 - Homologação</xs:documentation>
// 					</xs:annotation>
// 				</xs:element>
// 				<xs:element name="verAplic" type="TVerAplic">
// 					<xs:annotation>
// 						<xs:documentation>Versão do Aplicativo que processou o CT-e</xs:documentation>
// 					</xs:annotation>
// 				</xs:element>
// 				<xs:element name="chCTe" type="TChNFe">
// 					<xs:annotation>
// 						<xs:documentation>Chaves de acesso da CT-e, </xs:documentation>
// 					</xs:annotation>
// 				</xs:element>
// 				<xs:element name="dhRecbto" type="TDateTimeUTC">
// 					<xs:annotation>
// 						<xs:documentation>Data e hora de processamento, no formato AAAA-MM-DDTHH:MM:SS TZD. </xs:documentation>
// 					</xs:annotation>
// 				</xs:element>
// 				<xs:element name="nProt" type="TProt" minOccurs="0">
// 					<xs:annotation>
// 						<xs:documentation>Número do Protocolo de Status do CT-e. </xs:documentation>
// 					</xs:annotation>
// 				</xs:element>
// 				<xs:element name="digVal" type="ds:DigestValueType" minOccurs="0">
// 					<xs:annotation>
// 						<xs:documentation>Digest Value da CT-e processado. Utilizado para conferir a integridade do CT-e original.</xs:documentation>
// 					</xs:annotation>
// 				</xs:element>
// 				<xs:element name="cStat">
// 					<xs:annotation>
// 						<xs:documentation>Código do status do CT-e.</xs:documentation>
// 					</xs:annotation>
// 					<xs:simpleType>
// 						<xs:restriction base="TStat"/>
// 					</xs:simpleType>
// 				</xs:element>
// 				<xs:element name="xMotivo" type="TMotivo">
// 					<xs:annotation>
// 						<xs:documentation>Descrição literal do status do CT-e.</xs:documentation>
// 					</xs:annotation>
// 				</xs:element>
// 			</xs:sequence>
// 			<xs:attribute name="Id" type="xs:ID" use="optional"/>
// 		</xs:complexType>
// 	</xs:element>
// 	<xs:element name="infFisco" minOccurs="0">
// 		<xs:annotation>
// 			<xs:documentation>Mensagem do Fisco</xs:documentation>
// 		</xs:annotation>
// 		<xs:complexType>
// 			<xs:sequence>
// 				<xs:element name="cMsg">
// 					<xs:annotation>
// 						<xs:documentation>Código do status da mensagem do fisco</xs:documentation>
// 					</xs:annotation>
// 					<xs:simpleType>
// 						<xs:restriction base="TStat"/>
// 					</xs:simpleType>
// 				</xs:element>
// 				<xs:element name="xMsg" type="TMotivo">
// 					<xs:annotation>
// 						<xs:documentation>Mensagem do Fisco</xs:documentation>
// 					</xs:annotation>
// 				</xs:element>
// 			</xs:sequence>
// 		</xs:complexType>
// 	</xs:element>
// 	<xs:element ref="ds:Signature" minOccurs="0"/>
// </xs:sequence>
// <xs:attribute name="versao" use="required">
// 	<xs:simpleType>
// 		<xs:restriction base="TVerCTe"/>
// 	</xs:simpleType>
// </xs:attribute>
// </xs:complexType>

// TProtCTe Tipo Protocolo de status resultado do processamento da CT-e
type TProtCTe struct {
	XMLName xml.Name `xml:"protCTe"`
	InfProt InfProt  `xml:"infProt"`
}

// InfProt Dados do protocolo de status
type InfProt struct {
	XMLName xml.Name `xml:"infProt"`
	TpAmb   TpAmb    `xml:"tpAmb"`
}
