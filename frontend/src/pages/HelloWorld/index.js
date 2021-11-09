import React, { useState, useEffect } from 'react'

import { backendURL } from '../../constants'

const HelloWorld = () => {
  const [isLoaded, setIsLoaded] = useState(false)
  const [randomString, setRandomString] = useState("")

  useEffect(() => {
    fetch(`${backendURL}/random_string`)
    .then(res => res.text())
    .then(result => {
      setRandomString(result)
      setIsLoaded(true)
    })
  }, [])

  return (
    <div>
      <h1>Hello world</h1>
      {isLoaded ? 
        <p>{randomString}</p>
        : <p>Loading ....</p>
      }
    </div>
  )
}

export default HelloWorld