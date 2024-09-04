import { Title } from "@/components/title"

const AboutPage = () => {
  // Render About Page
  // ! ここでComponentsを呼び出して、Aboutページを組み上げる。
  // ! 独自のHTMLやCSSはできるだけ作らない。
  // ! 独自のHTMLやCSSが必要な場合、ここでComponentを上書きする。
  // ! 独自のHTMLやCSSは、Componentに持ち込まない。

  const pageTitle = "About Page"
  // ! componentsの引数は、State以外極力propsとして渡す。
  // ! componentsに渡す情報は、変数定義してから渡す。

  // ! ページに関するReacthooksは、ここで定義する。
  // ! - 再レンダリング
  // ! - Componentに渡す為に生成した情報のState管理
  // !   ただし、Componentで簡潔できることは、Componentで行う。
  // ! Componentsに関するReacthooksは、components内で定義する。
  return (
    <>
      <Title>{pageTitle}</Title>
    </>
  )
}

export default AboutPage
