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
func Html() *gmlElement { return newGmlElement(elHTML, false) }

// <head { attr }> { children } </head>
func Head() *gmlElement { return newGmlElement(elHead, false) }

// <body { attr }> { children } </body>
func Body() *gmlElement { return newGmlElement(elBody, false) }

// <a { attr }> { children } </a>
func A() *gmlElement { return newGmlElement(elA, false) }

// <article { attr }> { children } </article>
func Article() *gmlElement { return newGmlElement(elArticle, false) }

// <aside { attr }> { children } </aside>
func Aside() *gmlElement { return newGmlElement(elAside, false) }

// <details { attr }> { children } </details>
func Details() *gmlElement { return newGmlElement(elDetails, false) }

// <div { attr }> { children } </div>
func Div() *gmlElement { return newGmlElement(elDiv, false) }

// <header { attr }> { children } </header>
func Header() *gmlElement { return newGmlElement(elHeader, false) }

// <hgroup { attr }> { children } </hgroup>
func Hgroup() *gmlElement { return newGmlElement(elHgroup, false) }

// <h1 { attr }> { children } </h1>
func H1() *gmlElement { return newGmlElement(elH1, false) }
func H2() *gmlElement { return newGmlElement(elH2, false) }
func H3() *gmlElement { return newGmlElement(elH3, false) }
func H4() *gmlElement { return newGmlElement(elH4, false) }
func H5() *gmlElement { return newGmlElement(elH5, false) }
func H6() *gmlElement { return newGmlElement(elH6, false) }

func Footer() *gmlElement { return newGmlElement(elFooter, false) }
func Main() *gmlElement   { return newGmlElement(elMain, false) }
func Nav() *gmlElement    { return newGmlElement(elNav, false) }
func P() *gmlElement      { return newGmlElement(elP, false) }

func Section() *gmlElement { return newGmlElement(elSection, false) }
func Span() *gmlElement    { return newGmlElement(elSpan, false) }
func Summary() *gmlElement { return newGmlElement(elSummary, false) }
func Style() *gmlElement   { return newGmlElement(elStyle, false) }
func Title() *gmlElement   { return newGmlElement(elTitle, false) }

func Button() *gmlElement   { return newGmlElement(elButton, false) }
func Datalist() *gmlElement { return newGmlElement(elDatalist, false) }
func Dialog() *gmlElement   { return newGmlElement(elDialog, false) }
func Fieldset() *gmlElement { return newGmlElement(elFieldset, false) }
func Form() *gmlElement     { return newGmlElement(elForm, false) }
func Label() *gmlElement    { return newGmlElement(elLabel, false) }
func Legend() *gmlElement   { return newGmlElement(elLegend, false) }
func Meter() *gmlElement    { return newGmlElement(elMeter, false) }

func Optgroup() *gmlElement { return newGmlElement(elOptgroup, false) }
func Option() *gmlElement   { return newGmlElement(elOption, false) }
func Select() *gmlElement   { return newGmlElement(elSelect, false) }
func Textarea() *gmlElement { return newGmlElement(elTextarea, false) }
func Output() *gmlElement   { return newGmlElement(elOutput, false) }
func Progress() *gmlElement { return newGmlElement(elProgress, false) }

func Abbr() *gmlElement       { return newGmlElement(elAbbr, false) }
func Address() *gmlElement    { return newGmlElement(elAddress, false) }
func B() *gmlElement          { return newGmlElement(elB, false) }
func Bdi() *gmlElement        { return newGmlElement(elBdi, false) }
func Bdo() *gmlElement        { return newGmlElement(elBdo, false) }
func Blockquote() *gmlElement { return newGmlElement(elBlockquote, false) }
func Cite() *gmlElement       { return newGmlElement(elCite, false) }
func Code() *gmlElement       { return newGmlElement(elCode, false) }
func Data() *gmlElement       { return newGmlElement(elData, false) }
func Del() *gmlElement        { return newGmlElement(elDel, false) }
func Dfn() *gmlElement        { return newGmlElement(elDfn, false) }
func Em() *gmlElement         { return newGmlElement(elEm, false) }
func I() *gmlElement          { return newGmlElement(elI, false) }
func Ins() *gmlElement        { return newGmlElement(elIns, false) }
func Kbd() *gmlElement        { return newGmlElement(elKbd, false) }
func Mark() *gmlElement       { return newGmlElement(elMark, false) }
func Pre() *gmlElement        { return newGmlElement(elPre, false) }
func Q() *gmlElement          { return newGmlElement(elQ, false) }
func Rp() *gmlElement         { return newGmlElement(elRp, false) }
func Rt() *gmlElement         { return newGmlElement(elRt, false) }
func Ruby() *gmlElement       { return newGmlElement(elRuby, false) }
func S() *gmlElement          { return newGmlElement(elS, false) }
func Samp() *gmlElement       { return newGmlElement(elSamp, false) }
func Small() *gmlElement      { return newGmlElement(elSmall, false) }
func Strong() *gmlElement     { return newGmlElement(elStrong, false) }
func Sub() *gmlElement        { return newGmlElement(elSub, false) }
func Sup() *gmlElement        { return newGmlElement(elSup, false) }
func U() *gmlElement          { return newGmlElement(elU, false) }
func Var() *gmlElement        { return newGmlElement(elVar, false) }

func Dd() *gmlElement   { return newGmlElement(elDd, false) }
func Dl() *gmlElement   { return newGmlElement(elDl, false) }
func Dt() *gmlElement   { return newGmlElement(elDt, false) }
func Li() *gmlElement   { return newGmlElement(elLi, false) }
func Ol() *gmlElement   { return newGmlElement(elOl, false) }
func Menu() *gmlElement { return newGmlElement(elMenu, false) }
func Ul() *gmlElement   { return newGmlElement(elUl, false) }

func Caption() *gmlElement  { return newGmlElement(elCaption, false) }
func Colgroup() *gmlElement { return newGmlElement(elColgroup, false) }
func Table() *gmlElement    { return newGmlElement(elTable, false) }
func Tbody() *gmlElement    { return newGmlElement(elTbody, false) }
func Td() *gmlElement       { return newGmlElement(elTd, false) }
func Tfoot() *gmlElement    { return newGmlElement(elTfoot, false) }
func Thead() *gmlElement    { return newGmlElement(elThead, false) }
func Th() *gmlElement       { return newGmlElement(elTh, false) }
func Tr() *gmlElement       { return newGmlElement(elTr, false) }

func Noscript() *gmlElement { return newGmlElement(elNoscript, false) }
func Script() *gmlElement   { return newGmlElement(elScript, false) }
func Template() *gmlElement { return newGmlElement(elTemplate, false) }

func Audio() *gmlElement      { return newGmlElement(elAudio, false) }
func Canvas() *gmlElement     { return newGmlElement(elCanvas, false) }
func Figcaption() *gmlElement { return newGmlElement(elFigcaption, false) }
func Figure() *gmlElement     { return newGmlElement(elFigure, false) }
func Iframe() *gmlElement     { return newGmlElement(elIframe, false) }
func Map() *gmlElement        { return newGmlElement(elMap, false) }
func Object() *gmlElement     { return newGmlElement(elObject, false) }
func Picture() *gmlElement    { return newGmlElement(elPicture, false) }
func Svg() *gmlElement        { return newGmlElement(elSvg, false) }
func Time() *gmlElement       { return newGmlElement(elTime, false) }
func Video() *gmlElement      { return newGmlElement(elVideo, false) }
