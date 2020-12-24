package parser

import (
	"encoding/xml"
	"io"
	"valdemarceccon/cte_parser/parser/types"
)

// Parse Processa arquivo xml em seu tipo
func Parse(reader io.Reader) (*types.ProcCTE, error) {
	decoder := xml.NewDecoder(reader)
	ret := &types.ProcCTE{}
	err := decoder.Decode(ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
