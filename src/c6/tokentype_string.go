// generated by stringer -type=TokenType; DO NOT EDIT

package c6

import "fmt"

const _TokenType_name = "T_SPACET_COMMENT_LINET_COMMENT_BLOCKT_SEMICOLONT_COMMAT_ID_SELECTORT_CLASS_SELECTORT_TAGNAME_SELECTORT_UNIVERSAL_SELECTORT_PARENT_SELECTORT_PSEUDO_SELECTORT_AND_SELECTORT_PLUST_GTT_BRACE_STARTT_BRACE_ENDT_LANG_CODET_BRACKET_LEFTT_ATTRIBUTE_NAMET_BRACKET_RIGHTT_EQUALT_TILDE_EQUALT_PIPE_EQUALT_VARIABLET_IMPORTT_CHARSETT_QQ_STRINGT_Q_STRINGT_UNQUOTE_STRINGT_PAREN_STARTT_PAREN_ENDT_CONSTANTT_INTEGERT_FLOATT_UNIT_PXT_UNIT_PTT_UNIT_EMT_UNIT_DEGT_UNIT_PERCENTT_PROPERTY_NAMET_PROPERTY_VALUET_HEX_COLORT_COLONT_INTERPOLATION_STARTT_INTERPOLATION_END"

var _TokenType_index = [...]uint16{0, 7, 21, 36, 47, 54, 67, 83, 101, 121, 138, 155, 169, 175, 179, 192, 203, 214, 228, 244, 259, 266, 279, 291, 301, 309, 318, 329, 339, 355, 368, 379, 389, 398, 405, 414, 423, 432, 442, 456, 471, 487, 498, 505, 526, 545}

func (i TokenType) String() string {
	if i < 0 || i+1 >= TokenType(len(_TokenType_index)) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
