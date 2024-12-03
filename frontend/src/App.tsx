import { useState, useEffect } from "react";
import "./App.css";
import JobCard from "./components/card/CardJob";
import { Job } from "./components/card/CardJob";
import CreateJob from "./components/card/CreateJob";
import DeleteAll from "./components/DeleteAll";

function App() {
  const [message, setMessage] = useState<Job[]>([]);

  const fetchData = () => {
    fetch("http://localhost:8080/api/jobs")
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        const formattedData = data.map((job: any) => ({
          Id: job.id,
          Name: job.name,
          Source: job.source,
          Description: job.description,
          CreatedAt: job.createdat,
        }));
        setMessage(formattedData);
      })
      .catch((error) => console.error("Error fetching data:", error));
  };

  useEffect(() => {
    fetchData(); // Fetch data immediately on mount

    const intervalId = setInterval(() => {
      fetchData(); // Fetch data every 10 seconds
    }, 10000); // 10000 milliseconds = 10 seconds

    return () => clearInterval(intervalId); // Cleanup on unmount
  }, []);

  return (
    <>
      <DeleteAll onDeleteAll={fetchData} />
      <button
        onClick={fetchData}
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        style={{ position: "absolute", top: "20px", left: "20px" }}
      >
        Update jobs!
      </button>
      {message.length > 0 && (
        <div>
          {message.map((job: Job) => (
            <JobCard key={job.Id} job={job} onDelete={() => {}} />
          ))}
        </div>
      )}
      <CreateJob />
    </>
  );
}

export default App;
