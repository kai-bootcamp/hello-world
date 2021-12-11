import type { NextPage } from 'next'
import Head from 'next/head'
import { useState } from 'react'
import styles from '../styles/Home.module.css'

type Message = {
    string: string
}

const Home: NextPage = () => {
  const [data, setData] = useState<Message>({string: ''})

  const fetchRandomString = async () => {
    const res = await fetch('/api/strings/random')
    const resData: Message = await res.json()
    setData(resData)
  }

  return (
    <div className={styles.container}>
      <Head>
        <title>Bootcamp</title>
        <meta name="description" content="Bootcamp" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          HELLO WORLD
        </h1>

        <div className={styles.grid}>
          <button onClick={fetchRandomString} className={styles.card}>
            Get random string
          </button>
          <p>{data.string}</p>
        </div>
      </main>

      <footer className={styles.footer}>
        KAI Bootcamp
      </footer>
    </div>
  )
}

export default Home
