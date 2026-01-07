package gml

const (
	elHTML uint8 = iota
	elHead
	elBody
	elA
	elArticle
	elAside
	elDetails
	elDiv
	elHeader
	elHgroup
	elH1
	elH2
	elH3
	elH4
	elH5
	elH6
	elFooter
	elMain
	elNav
	elP
	elSection
	elSpan
	elSummary
	elStyle
	elTitle
	elButton
	elDatalist
	elDialog
	elFieldset
	elForm
	elLabel
	elLegend
	elMeter
	elOptgroup
	elOption
	elSelect
	elTextarea
	elOutput
	elProgress
	elAbbr
	elAddress
	elB
	elBdi
	elBdo
	elBlockquote
	elCite
	elCode
	elData
	elDel
	elDfn
	elEm
	elI
	elIns
	elKbd
	elMark
	elPre
	elQ
	elRp
	elRt
	elRuby
	elS
	elSamp
	elSmall
	elStrong
	elSub
	elSup
	elU
	elVar
	elDd
	elDl
	elDt
	elLi
	elOl
	elMenu
	elUl
	elCaption
	elColgroup
	elTable
	elTbody
	elTd
	elTfoot
	elThead
	elTh
	elTr
	elNoscript
	elScript
	elTemplate
	elAudio
	elCanvas
	elFigcaption
	elFigure
	elIframe
	elMap
	elObject
	elPicture
	elSvg
	elTime
	elVideo
)

var elementNames = map[uint8][]byte{
	elHTML:       []byte("html"),
	elHead:       []byte("head"),
	elBody:       []byte("body"),
	elA:          []byte("a"),
	elArticle:    []byte("article"),
	elAside:      []byte("aside"),
	elDetails:    []byte("details"),
	elDiv:        []byte("div"),
	elHeader:     []byte("header"),
	elHgroup:     []byte("hgroup"),
	elH1:         []byte("h1"),
	elH2:         []byte("h2"),
	elH3:         []byte("h3"),
	elH4:         []byte("h4"),
	elH5:         []byte("h5"),
	elH6:         []byte("h6"),
	elFooter:     []byte("footer"),
	elMain:       []byte("main"),
	elNav:        []byte("nav"),
	elP:          []byte("p"),
	elSection:    []byte("section"),
	elSpan:       []byte("span"),
	elSummary:    []byte("summary"),
	elStyle:      []byte("style"),
	elTitle:      []byte("title"),
	elButton:     []byte("button"),
	elDatalist:   []byte("datalist"),
	elDialog:     []byte("dialog"),
	elFieldset:   []byte("fieldset"),
	elForm:       []byte("form"),
	elLabel:      []byte("label"),
	elLegend:     []byte("legend"),
	elMeter:      []byte("meter"),
	elOptgroup:   []byte("optgroup"),
	elOption:     []byte("option"),
	elSelect:     []byte("select"),
	elTextarea:   []byte("textarea"),
	elOutput:     []byte("output"),
	elProgress:   []byte("progress"),
	elAbbr:       []byte("abbr"),
	elAddress:    []byte("address"),
	elB:          []byte("b"),
	elBdi:        []byte("bdi"),
	elBdo:        []byte("bdo"),
	elBlockquote: []byte("blockquote"),
	elCite:       []byte("cite"),
	elCode:       []byte("code"),
	elData:       []byte("data"),
	elDel:        []byte("del"),
	elDfn:        []byte("dfn"),
	elEm:         []byte("em"),
	elI:          []byte("i"),
	elIns:        []byte("ins"),
	elKbd:        []byte("kbd"),
	elMark:       []byte("mark"),
	elPre:        []byte("pre"),
	elQ:          []byte("q"),
	elRp:         []byte("rp"),
	elRt:         []byte("rt"),
	elRuby:       []byte("ruby"),
	elS:          []byte("s"),
	elSamp:       []byte("samp"),
	elSmall:      []byte("small"),
	elStrong:     []byte("strong"),
	elSub:        []byte("sub"),
	elSup:        []byte("sup"),
	elU:          []byte("u"),
	elVar:        []byte("var"),
	elDd:         []byte("dd"),
	elDl:         []byte("dl"),
	elDt:         []byte("dt"),
	elLi:         []byte("li"),
	elOl:         []byte("ol"),
	elMenu:       []byte("menu"),
	elUl:         []byte("ul"),
	elCaption:    []byte("caption"),
	elColgroup:   []byte("colgroup"),
	elTable:      []byte("table"),
	elTbody:      []byte("tbody"),
	elTd:         []byte("td"),
	elTfoot:      []byte("tfoot"),
	elThead:      []byte("thead"),
	elTh:         []byte("th"),
	elTr:         []byte("tr"),
	elNoscript:   []byte("noscript"),
	elScript:     []byte("script"),
	elTemplate:   []byte("template"),
	elAudio:      []byte("audio"),
	elCanvas:     []byte("canvas"),
	elFigcaption: []byte("figcaption"),
	elFigure:     []byte("figure"),
	elIframe:     []byte("iframe"),
	elMap:        []byte("map"),
	elObject:     []byte("object"),
	elPicture:    []byte("picture"),
	elSvg:        []byte("svg"),
	elTime:       []byte("time"),
	elVideo:      []byte("video"),
}

// <html { attr }> { children } </html>
func Html() *gmlElement { return newgmlElement(elHTML, false) }

// <head { attr }> { children } </head>
func Head() *gmlElement { return newgmlElement(elHead, false) }

// <body { attr }> { children } </body>
func Body() *gmlElement { return newgmlElement(elBody, false) }

// <a { attr }> { children } </a>
func A() *gmlElement { return newgmlElement(elA, false) }

// <article { attr }> { children } </article>
func Article() *gmlElement { return newgmlElement(elArticle, false) }

// <aside { attr }> { children } </aside>
func Aside() *gmlElement { return newgmlElement(elAside, false) }

// <details { attr }> { children } </details>
func Details() *gmlElement { return newgmlElement(elDetails, false) }

// <div { attr }> { children } </div>
func Div() *gmlElement { return newgmlElement(elDiv, false) }

// <header { attr }> { children } </header>
func Header() *gmlElement { return newgmlElement(elHeader, false) }

// <hgroup { attr }> { children } </hgroup>
func Hgroup() *gmlElement { return newgmlElement(elHgroup, false) }

// <h1 { attr }> { children } </h1>
func H1() *gmlElement { return newgmlElement(elH1, false) }
func H2() *gmlElement { return newgmlElement(elH2, false) }
func H3() *gmlElement { return newgmlElement(elH3, false) }
func H4() *gmlElement { return newgmlElement(elH4, false) }
func H5() *gmlElement { return newgmlElement(elH5, false) }
func H6() *gmlElement { return newgmlElement(elH6, false) }

func Footer() *gmlElement { return newgmlElement(elFooter, false) }
func Main() *gmlElement   { return newgmlElement(elMain, false) }
func Nav() *gmlElement    { return newgmlElement(elNav, false) }
func P() *gmlElement      { return newgmlElement(elP, false) }

func Section() *gmlElement { return newgmlElement(elSection, false) }
func Span() *gmlElement    { return newgmlElement(elSpan, false) }
func Summary() *gmlElement { return newgmlElement(elSummary, false) }
func Style() *gmlElement   { return newgmlElement(elStyle, false) }
func Title() *gmlElement   { return newgmlElement(elTitle, false) }

func Button() *gmlElement   { return newgmlElement(elButton, false) }
func Datalist() *gmlElement { return newgmlElement(elDatalist, false) }
func Dialog() *gmlElement   { return newgmlElement(elDialog, false) }
func Fieldset() *gmlElement { return newgmlElement(elFieldset, false) }
func Form() *gmlElement     { return newgmlElement(elForm, false) }
func Label() *gmlElement    { return newgmlElement(elLabel, false) }
func Legend() *gmlElement   { return newgmlElement(elLegend, false) }
func Meter() *gmlElement    { return newgmlElement(elMeter, false) }

func Optgroup() *gmlElement { return newgmlElement(elOptgroup, false) }
func Option() *gmlElement   { return newgmlElement(elOption, false) }
func Select() *gmlElement   { return newgmlElement(elSelect, false) }
func Textarea() *gmlElement { return newgmlElement(elTextarea, false) }
func Output() *gmlElement   { return newgmlElement(elOutput, false) }
func Progress() *gmlElement { return newgmlElement(elProgress, false) }

func Abbr() *gmlElement       { return newgmlElement(elAbbr, false) }
func Address() *gmlElement    { return newgmlElement(elAddress, false) }
func B() *gmlElement          { return newgmlElement(elB, false) }
func Bdi() *gmlElement        { return newgmlElement(elBdi, false) }
func Bdo() *gmlElement        { return newgmlElement(elBdo, false) }
func Blockquote() *gmlElement { return newgmlElement(elBlockquote, false) }
func Cite() *gmlElement       { return newgmlElement(elCite, false) }
func Code() *gmlElement       { return newgmlElement(elCode, false) }
func Data() *gmlElement       { return newgmlElement(elData, false) }
func Del() *gmlElement        { return newgmlElement(elDel, false) }
func Dfn() *gmlElement        { return newgmlElement(elDfn, false) }
func Em() *gmlElement         { return newgmlElement(elEm, false) }
func I() *gmlElement          { return newgmlElement(elI, false) }
func Ins() *gmlElement        { return newgmlElement(elIns, false) }
func Kbd() *gmlElement        { return newgmlElement(elKbd, false) }
func Mark() *gmlElement       { return newgmlElement(elMark, false) }
func Pre() *gmlElement        { return newgmlElement(elPre, false) }
func Q() *gmlElement          { return newgmlElement(elQ, false) }
func Rp() *gmlElement         { return newgmlElement(elRp, false) }
func Rt() *gmlElement         { return newgmlElement(elRt, false) }
func Ruby() *gmlElement       { return newgmlElement(elRuby, false) }
func S() *gmlElement          { return newgmlElement(elS, false) }
func Samp() *gmlElement       { return newgmlElement(elSamp, false) }
func Small() *gmlElement      { return newgmlElement(elSmall, false) }
func Strong() *gmlElement     { return newgmlElement(elStrong, false) }
func Sub() *gmlElement        { return newgmlElement(elSub, false) }
func Sup() *gmlElement        { return newgmlElement(elSup, false) }
func U() *gmlElement          { return newgmlElement(elU, false) }
func Var() *gmlElement        { return newgmlElement(elVar, false) }

func Dd() *gmlElement   { return newgmlElement(elDd, false) }
func Dl() *gmlElement   { return newgmlElement(elDl, false) }
func Dt() *gmlElement   { return newgmlElement(elDt, false) }
func Li() *gmlElement   { return newgmlElement(elLi, false) }
func Ol() *gmlElement   { return newgmlElement(elOl, false) }
func Menu() *gmlElement { return newgmlElement(elMenu, false) }
func Ul() *gmlElement   { return newgmlElement(elUl, false) }

func Caption() *gmlElement  { return newgmlElement(elCaption, false) }
func Colgroup() *gmlElement { return newgmlElement(elColgroup, false) }
func Table() *gmlElement    { return newgmlElement(elTable, false) }
func Tbody() *gmlElement    { return newgmlElement(elTbody, false) }
func Td() *gmlElement       { return newgmlElement(elTd, false) }
func Tfoot() *gmlElement    { return newgmlElement(elTfoot, false) }
func Thead() *gmlElement    { return newgmlElement(elThead, false) }
func Th() *gmlElement       { return newgmlElement(elTh, false) }
func Tr() *gmlElement       { return newgmlElement(elTr, false) }

func Noscript() *gmlElement { return newgmlElement(elNoscript, false) }
func Script() *gmlElement   { return newgmlElement(elScript, false) }
func Template() *gmlElement { return newgmlElement(elTemplate, false) }

func Audio() *gmlElement      { return newgmlElement(elAudio, false) }
func Canvas() *gmlElement     { return newgmlElement(elCanvas, false) }
func Figcaption() *gmlElement { return newgmlElement(elFigcaption, false) }
func Figure() *gmlElement     { return newgmlElement(elFigure, false) }
func Iframe() *gmlElement     { return newgmlElement(elIframe, false) }
func Map() *gmlElement        { return newgmlElement(elMap, false) }
func Object() *gmlElement     { return newgmlElement(elObject, false) }
func Picture() *gmlElement    { return newgmlElement(elPicture, false) }
func Svg() *gmlElement        { return newgmlElement(elSvg, false) }
func Time() *gmlElement       { return newgmlElement(elTime, false) }
func Video() *gmlElement      { return newgmlElement(elVideo, false) }
