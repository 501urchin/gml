package gml

// <html { attr }> { children } </html>
func Html(attributes []Attr, children ...HtmlElement) Element {
	return newElement("html", false, attributes, children...)
}

// <head { attr }> { children } </head>
func Head(attributes []Attr, children ...HtmlElement) Element {
	return newElement("head", false, attributes, children...)
}

// <body { attr }> { children } </body>
func Body(attributes []Attr, children ...HtmlElement) Element {
	return newElement("body", false, attributes, children...)
}

// <article { attr }> { children } </article>
func Article(attributes []Attr, children ...HtmlElement) Element {
	return newElement("article", false, attributes, children...)
}

// <aside { attr }> { children } </aside>
func Aside(attributes []Attr, children ...HtmlElement) Element {
	return newElement("aside", false, attributes, children...)
}

// <details { attr }> { children } </details>
func Details(attributes []Attr, children ...HtmlElement) Element {
	return newElement("details", false, attributes, children...)
}

// <div { attr }> { children } </div>
func Div(attributes []Attr, children ...HtmlElement) Element {
	return newElement("div", false, attributes, children...)
}

// <header { attr }> { children } </header>
func Header(attributes []Attr, children ...HtmlElement) Element {
	return newElement("header", false, attributes, children...)
}

// <hgroup { attr }> { children } </hgroup>
func Hgroup(attributes []Attr, children ...HtmlElement) Element {
	return newElement("hgroup", false, attributes, children...)
}

// <h1 { attr }> { children } </h1>
func H1(attributes []Attr, children ...HtmlElement) Element {
	return newElement("h1", false, attributes, children...)
}

// <h2 { attr }> { children } </h2>
func H2(attributes []Attr, children ...HtmlElement) Element {
	return newElement("h2", false, attributes, children...)
}

// <h3 { attr }> { children } </h3>
func H3(attributes []Attr, children ...HtmlElement) Element {
	return newElement("h3", false, attributes, children...)
}

// <h4 { attr }> { children } </h4>
func H4(attributes []Attr, children ...HtmlElement) Element {
	return newElement("h4", false, attributes, children...)
}

// <h5 { attr }> { children } </h5>
func H5(attributes []Attr, children ...HtmlElement) Element {
	return newElement("h5", false, attributes, children...)
}

// <h6 { attr }> { children } </h6>
func H6(attributes []Attr, children ...HtmlElement) Element {
	return newElement("h6", false, attributes, children...)
}

// <footer { attr }> { children } </footer>
func Footer(attributes []Attr, children ...HtmlElement) Element {
	return newElement("footer", false, attributes, children...)
}

// <main { attr }> { children } </main>
func Main(attributes []Attr, children ...HtmlElement) Element {
	return newElement("main", false, attributes, children...)
}

// <nav { attr }> { children } </nav>
func Nav(attributes []Attr, children ...HtmlElement) Element {
	return newElement("nav", false, attributes, children...)
}

// <p { attr }> { children } </p>
func P(attributes []Attr, children ...HtmlElement) Element {
	return newElement("p", false, attributes, children...)
}

// <section { attr }> { children } </section>
func Section(attributes []Attr, children ...HtmlElement) Element {
	return newElement("section", false, attributes, children...)
}

// <span { attr }> { children } </span>
func Span(attributes []Attr, children ...HtmlElement) Element {
	return newElement("span", false, attributes, children...)
}

// <summary { attr }> { children } </summary>
func Summary(attributes []Attr, children ...HtmlElement) Element {
	return newElement("summary", false, attributes, children...)
}

// <style { attr }> { children } </style>
func Style(attributes []Attr, children ...HtmlElement) Element {
	return newElement("style", false, attributes, children...)
}

// <title { attr }> { children } </title>
func Title(attributes []Attr, children ...HtmlElement) Element {
	return newElement("title", false, attributes, children...)
}

// <button { attr }> { children } </button>
func Button(attributes []Attr, children ...HtmlElement) Element {
	return newElement("button", false, attributes, children...)
}

// <datalist { attr }> { children } </datalist>
func Datalist(attributes []Attr, children ...HtmlElement) Element {
	return newElement("datalist", false, attributes, children...)
}

// <dialog { attr }> { children } </dialog>
func Dialog(attributes []Attr, children ...HtmlElement) Element {
	return newElement("dialog", false, attributes, children...)
}

// <fieldset { attr }> { children } </fieldset>
func Fieldset(attributes []Attr, children ...HtmlElement) Element {
	return newElement("fieldset", false, attributes, children...)
}

// <form { attr }> { children } </form>
func Form(attributes []Attr, children ...HtmlElement) Element {
	return newElement("form", false, attributes, children...)
}

// <label { attr }> { children } </label>
func Label(attributes []Attr, children ...HtmlElement) Element {
	return newElement("label", false, attributes, children...)
}

// <legend { attr }> { children } </legend>
func Legend(attributes []Attr, children ...HtmlElement) Element {
	return newElement("legend", false, attributes, children...)
}

// <meter { attr }> { children } </meter>
func Meter(attributes []Attr, children ...HtmlElement) Element {
	return newElement("meter", false, attributes, children...)
}

// <optgroup { attr }> { children } </optgroup>
func Optgroup(attributes []Attr, children ...HtmlElement) Element {
	return newElement("optgroup", false, attributes, children...)
}

// <option { attr }> { children } </option>
func Option(attributes []Attr, children ...HtmlElement) Element {
	return newElement("option", false, attributes, children...)
}

// <select { attr }> { children } </select>
func Select(attributes []Attr, children ...HtmlElement) Element {
	return newElement("select", false, attributes, children...)
}

// <textarea { attr }> { children } </textarea>
func Textarea(attributes []Attr, children ...HtmlElement) Element {
	return newElement("textarea", false, attributes, children...)
}

// <output { attr }> { children } </output>
func Output(attributes []Attr, children ...HtmlElement) Element {
	return newElement("output", false, attributes, children...)
}

// <progress { attr }> { children } </progress>
func Progress(attributes []Attr, children ...HtmlElement) Element {
	return newElement("progress", false, attributes, children...)
}

// <abbr { attr }> { children } </abbr>
func Abbr(attributes []Attr, children ...HtmlElement) Element {
	return newElement("abbr", false, attributes, children...)
}

// <address { attr }> { children } </address>
func Address(attributes []Attr, children ...HtmlElement) Element {
	return newElement("address", false, attributes, children...)
}

// <b { attr }> { children } </b>
func B(attributes []Attr, children ...HtmlElement) Element {
	return newElement("b", false, attributes, children...)
}

// <bdi { attr }> { children } </bdi>
func Bdi(attributes []Attr, children ...HtmlElement) Element {
	return newElement("bdi", false, attributes, children...)
}

// <bdo { attr }> { children } </bdo>
func Bdo(attributes []Attr, children ...HtmlElement) Element {
	return newElement("bdo", false, attributes, children...)
}

// <blockquote { attr }> { children } </blockquote>
func Blockquote(attributes []Attr, children ...HtmlElement) Element {
	return newElement("blockquote", false, attributes, children...)
}

// <cite { attr }> { children } </cite>
func Cite(attributes []Attr, children ...HtmlElement) Element {
	return newElement("cite", false, attributes, children...)
}

// <code { attr }> { children } </code>
func Code(attributes []Attr, children ...HtmlElement) Element {
	return newElement("code", false, attributes, children...)
}

// <data { attr }> { children } </data>
func Data(attributes []Attr, children ...HtmlElement) Element {
	return newElement("data", false, attributes, children...)
}

// <del { attr }> { children } </del>
func Del(attributes []Attr, children ...HtmlElement) Element {
	return newElement("del", false, attributes, children...)
}

// <dfn { attr }> { children } </dfn>
func Dfn(attributes []Attr, children ...HtmlElement) Element {
	return newElement("dfn", false, attributes, children...)
}

// <em { attr }> { children } </em>
func Em(attributes []Attr, children ...HtmlElement) Element {
	return newElement("em", false, attributes, children...)
}

// <i { attr }> { children } </i>
func I(attributes []Attr, children ...HtmlElement) Element {
	return newElement("i", false, attributes, children...)
}

// <ins { attr }> { children } </ins>
func Ins(attributes []Attr, children ...HtmlElement) Element {
	return newElement("ins", false, attributes, children...)
}

// <kbd { attr }> { children } </kbd>
func Kbd(attributes []Attr, children ...HtmlElement) Element {
	return newElement("kbd", false, attributes, children...)
}

// <mark { attr }> { children } </mark>
func Mark(attributes []Attr, children ...HtmlElement) Element {
	return newElement("mark", false, attributes, children...)
}

// <pre { attr }> { children } </pre>
func Pre(attributes []Attr, children ...HtmlElement) Element {
	return newElement("pre", false, attributes, children...)
}

// <q { attr }> { children } </q>
func Q(attributes []Attr, children ...HtmlElement) Element {
	return newElement("q", false, attributes, children...)
}

// <rp { attr }> { children } </rp>
func Rp(attributes []Attr, children ...HtmlElement) Element {
	return newElement("rp", false, attributes, children...)
}

// <rt { attr }> { children } </rt>
func Rt(attributes []Attr, children ...HtmlElement) Element {
	return newElement("rt", false, attributes, children...)
}

// <ruby { attr }> { children } </ruby>
func Ruby(attributes []Attr, children ...HtmlElement) Element {
	return newElement("ruby", false, attributes, children...)
}

// <s { attr }> { children } </s>
func S(attributes []Attr, children ...HtmlElement) Element {
	return newElement("s", false, attributes, children...)
}

// <samp { attr }> { children } </samp>
func Samp(attributes []Attr, children ...HtmlElement) Element {
	return newElement("samp", false, attributes, children...)
}

// <small { attr }> { children } </small>
func Small(attributes []Attr, children ...HtmlElement) Element {
	return newElement("small", false, attributes, children...)
}

// <strong { attr }> { children } </strong>
func Strong(attributes []Attr, children ...HtmlElement) Element {
	return newElement("strong", false, attributes, children...)
}

// <sub { attr }> { children } </sub>
func Sub(attributes []Attr, children ...HtmlElement) Element {
	return newElement("sub", false, attributes, children...)
}

// <sup { attr }> { children } </sup>
func Sup(attributes []Attr, children ...HtmlElement) Element {
	return newElement("sup", false, attributes, children...)
}

// <u { attr }> { children } </u>
func U(attributes []Attr, children ...HtmlElement) Element {
	return newElement("u", false, attributes, children...)
}

// <var { attr }> { children } </var>
func Var(attributes []Attr, children ...HtmlElement) Element {
	return newElement("var", false, attributes, children...)
}

// <dd { attr }> { children } </dd>
func Dd(attributes []Attr, children ...HtmlElement) Element {
	return newElement("dd", false, attributes, children...)
}

// <dl { attr }> { children } </dl>
func Dl(attributes []Attr, children ...HtmlElement) Element {
	return newElement("dl", false, attributes, children...)
}

// <dt { attr }> { children } </dt>
func Dt(attributes []Attr, children ...HtmlElement) Element {
	return newElement("dt", false, attributes, children...)
}

// <li { attr }> { children } </li>
func Li(attributes []Attr, children ...HtmlElement) Element {
	return newElement("li", false, attributes, children...)
}

// <ol { attr }> { children } </ol>
func Ol(attributes []Attr, children ...HtmlElement) Element {
	return newElement("ol", false, attributes, children...)
}

// <menu { attr }> { children } </menu>
func Menu(attributes []Attr, children ...HtmlElement) Element {
	return newElement("menu", false, attributes, children...)
}

// <ul { attr }> { children } </ul>
func Ul(attributes []Attr, children ...HtmlElement) Element {
	return newElement("ul", false, attributes, children...)
}

// <caption { attr }> { children } </caption>
func Caption(attributes []Attr, children ...HtmlElement) Element {
	return newElement("caption", false, attributes, children...)
}

// <colgroup { attr }> { children } </colgroup>
func Colgroup(attributes []Attr, children ...HtmlElement) Element {
	return newElement("colgroup", false, attributes, children...)
}

// <table { attr }> { children } </table>
func Table(attributes []Attr, children ...HtmlElement) Element {
	return newElement("table", false, attributes, children...)
}

// <tbody { attr }> { children } </tbody>
func Tbody(attributes []Attr, children ...HtmlElement) Element {
	return newElement("tbody", false, attributes, children...)
}

// <td { attr }> { children } </td>
func Td(attributes []Attr, children ...HtmlElement) Element {
	return newElement("td", false, attributes, children...)
}

// <tfoot { attr }> { children } </tfoot>
func Tfoot(attributes []Attr, children ...HtmlElement) Element {
	return newElement("tfoot", false, attributes, children...)
}

// <thead { attr }> { children } </thead>
func Thead(attributes []Attr, children ...HtmlElement) Element {
	return newElement("thead", false, attributes, children...)
}

// <th { attr }> { children } </th>
func Th(attributes []Attr, children ...HtmlElement) Element {
	return newElement("th", false, attributes, children...)
}

// <tr { attr }> { children } </tr>
func Tr(attributes []Attr, children ...HtmlElement) Element {
	return newElement("tr", false, attributes, children...)
}

// <noscript { attr }> { children } </noscript>
func Noscript(attributes []Attr, children ...HtmlElement) Element {
	return newElement("noscript", false, attributes, children...)
}

// <script { attr }> { children } </script>
func Script(attributes []Attr, children ...HtmlElement) Element {
	return newElement("script", false, attributes, children...)
}

// <template { attr }> { children } </template>
func Template(attributes []Attr, children ...HtmlElement) Element {
	return newElement("template", false, attributes, children...)
}

// <audio { attr }> { children } </audio>
func Audio(attributes []Attr, children ...HtmlElement) Element {
	return newElement("audio", false, attributes, children...)
}

// <canvas { attr }> { children } </canvas>
func Canvas(attributes []Attr, children ...HtmlElement) Element {
	return newElement("canvas", false, attributes, children...)
}

// <figcaption { attr }> { children } </figcaption>
func Figcaption(attributes []Attr, children ...HtmlElement) Element {
	return newElement("figcaption", false, attributes, children...)
}

// <figure { attr }> { children } </figure>
func Figure(attributes []Attr, children ...HtmlElement) Element {
	return newElement("figure", false, attributes, children...)
}

// <iframe { attr }> { children } </iframe>
func Iframe(attributes []Attr, children ...HtmlElement) Element {
	return newElement("iframe", false, attributes, children...)
}

// <map { attr }> { children } </map>
func Map(attributes []Attr, children ...HtmlElement) Element {
	return newElement("map", false, attributes, children...)
}

// <object { attr }> { children } </object>
func Object(attributes []Attr, children ...HtmlElement) Element {
	return newElement("object", false, attributes, children...)
}

// <picture { attr }> { children } </picture>
func Picture(attributes []Attr, children ...HtmlElement) Element {
	return newElement("picture", false, attributes, children...)
}

// <svg { attr }> { children } </svg>
func Svg(attributes []Attr, children ...HtmlElement) Element {
	return newElement("svg", false, attributes, children...)
}

// <time { attr }> { children } </time>
func Time(attributes []Attr, children ...HtmlElement) Element {
	return newElement("time", false, attributes, children...)
}

// <video { attr }> { children } </video>
func Video(attributes []Attr, children ...HtmlElement) Element {
	return newElement("video", false, attributes, children...)
}
