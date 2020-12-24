package types

// <xs:simpleType name="TAmb">
// <xs:annotation>
//   <xs:documentation>Tipo Ambiente</xs:documentation>
// </xs:annotation>
// <xs:restriction base="xs:string">
//   <xs:whiteSpace value="preserve"/>
//   <xs:enumeration value="1"/>
//   <xs:enumeration value="2"/>
// </xs:restriction>
// </xs:simpleType>

// TpAmb Identificação do Ambiente:1 - Produção; 2 - Homologação
type TpAmb string

const (
	// TpAmbProd Ambiente de produção
	TpAmbProd TpAmb = "1"

	// TpAmbHomol Ambiente de Homologação
	TpAmbHomol TpAmb = "2"
)

func (tp TpAmb) String() string {
	switch tp {
	case TpAmbProd:
		return "Produção"
	case TpAmbHomol:
		return "Homologação"
	default:
		return "Inválido"
	}
}
