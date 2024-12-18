import { useState, useEffect } from "react";
import "./App.css";
import JobCard from "./components/CardJob";
import { Job } from "./components/CardJob";
import CreateJob from "./components/CreateJob";
import DeleteAll from "./components/DeleteAll";
import { Button } from "./components/ui/button";

function App() {
  const [message, setMessage] = useState<Job[]>([]);

  const fetchData = () => {
    fetch("http://localhost:8080/api/jobs")
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        interface JobData {
          id: string;
          name: string;
          source: string;
          description: string;
          createdat: string;
        }

        const formattedData = data.map((job: JobData) => ({
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
    <div className="flex flex-col gap-4 p-4">
      <div className="flex gap-2 justify-center">
        <DeleteAll onDeleteAll={fetchData} />
        <Button
          variant="outline"
          onClick={fetchData}
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        >
          Update jobs!
        </Button>
        <CreateJob />
      </div>

      {message.length > 0 && (
        <div className="grid grid-cols-3 gap-4">
          {" "}
          {/* Changed to 3 columns */}
          {message.map((job: Job) => (
            <JobCard key={job.Id} job={job} onDelete={() => {}} />
          ))}
        </div>
      )}
    </div>
  );
}

export default App;
