package htmlToPdf

var cssMap map[string]string = map[string]string{
	"tr":          `flex-direction: row;`,
	"#htitle":     `margin:0px;font-size:26px;text-align:center;display:flex;flex-direction:row;`,
	"#htitle>div": `flex:1;line-height:44px;font-weight: bold;color:#ff0000;`,
	".dline":      `border-top:0.5px solid #f00;border-bottom:0.5px solid #f00;padding-top:1px;margin:5px 0px`}
