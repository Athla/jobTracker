import { useState } from 'react'
import './App.css'
import JobCard from './components/card/cardJob'
import { Job } from './components/card/cardJob'
import CreateJob from './components/card/CreateJob'

function App() {

  const fetchData = () => {
    fetch('http://localhost:8080/api/jobs')
      .then(response => response.json())
      .then(data => {
        console.log(data)
        const formattedData = data.map((job: any) => ({
          Id: job.id,
          Name: job.name,
          Source: job.source,
          Description: job.description,
          CreatedAt: job.createdat,
        }));
        setMessage(formattedData);
      })
      .catch(error => console.error('Error fetching data:', error))
  }
  
  const [message, setMessage] = useState<Job[]>([])

  return (
    <>
      <CreateJob />
      <button onClick={fetchData}>
        Click to fetch from Go server! Click here!
      </button>
      {message.length > 0 && (
        <div>
          <h2>Server Response:</h2>
          {message.map((job: Job) => (
            <JobCard key={job.Id} job={job} />
          ))}
        </div>
      )}
    </>
  )
}

export default App
