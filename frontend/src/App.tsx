import { useState } from 'react'

function App() {
  const [name, setName] = useState('')
  const [response, setResponse] = useState('')
  const [loading, setLoading] = useState(false)
  const [isError, setIsError] = useState(false)

  const callGreeting = async () => {
    setLoading(true)
    setIsError(false)
    try {
      const res = await fetch('/api/greeting', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name }),
      })
      const data = await res.json()
      setResponse(data.message || JSON.stringify(data, null, 2))
    } catch (err) {
      setIsError(true)
      setResponse(`Error: ${err}`)
    }
    setLoading(false)
  }

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && name && !loading) {
      callGreeting()
    }
  }

  return (
    <div className="container">
      <header>
        <h1>Worker <span>Dashboard</span></h1>
        <p className="subtitle">Test your dibbla worker functions</p>
      </header>

      <main>
        <section className="card">
          <h2>Greeting Function</h2>
          <div className="input-group">
            <input
              type="text"
              placeholder="Enter your name..."
              value={name}
              onChange={(e) => setName(e.target.value)}
              onKeyDown={handleKeyDown}
            />
            <button onClick={callGreeting} disabled={loading || !name}>
              {loading ? (
                <span className="loading-dots">Calling</span>
              ) : (
                'Call Function'
              )}
            </button>
          </div>
          {response && (
            <div className={`response ${isError ? 'error' : ''}`}>
              <strong>Response</strong>
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
