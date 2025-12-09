import { useState } from 'react'

function App() {
  const [name, setName] = useState('')
  const [response, setResponse] = useState('')
  const [loading, setLoading] = useState(false)

  const callGreeting = async () => {
    setLoading(true)
    try {
      // Replace with your actual API endpoint
      const res = await fetch('/api/greeting', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name }),
      })
      const data = await res.json()
      setResponse(data.message || JSON.stringify(data))
    } catch (err) {
      setResponse(`Error: ${err}`)
    }
    setLoading(false)
  }

  return (
    <div className="container">
      <header>
        <h1>Worker Dashboard</h1>
        <p className="subtitle">Test your worker functions</p>
      </header>

      <main>
        <section className="card">
          <h2>Greeting Function</h2>
          <div className="input-group">
            <input
              type="text"
              placeholder="Enter your name"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
            <button onClick={callGreeting} disabled={loading || !name}>
              {loading ? 'Loading...' : 'Call Function'}
            </button>
          </div>
          {response && (
            <div className="response">
              <strong>Response:</strong>
              <pre>{response}</pre>
            </div>
          )}
        </section>
      </main>

      <footer>
        <p>Worker Starter Template</p>
      </footer>
    </div>
  )
}

export default App

