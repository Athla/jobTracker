import { motion } from "framer-motion";
import { useState, useEffect } from "react";
import "./App.css";
import JobCard, { Job } from "./components/CardJob";
import CreateJob from "./components/CreateJob";
import DeleteAll from "./components/DeleteAll";
import { Button } from "./components/ui/button";
import Register from "./components/Register";
import Login from "./components/LoginCard"; // Import the Login component

function App() {
  const [data, setData] = useState<Job[]>([]);
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [token, setToken] = useState<string>("");

  const fetchData = () => {
    if (!token) return;

    fetch("http://localhost:8080/api/jobs", {
      headers: {
        Authorization: `${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        interface JobData {
          id: string;
          name: string;
          source: string;
          description: string;
          created_at: string;
        }

        const formattedData = data.map((job: JobData) => ({
          Id: job.id,
          Name: job.name,
          Source: job.source,
          Description: job.description,
          CreatedAt: job.created_at,
        }));
        setData(formattedData);
      })
      .catch((error) => console.error("Error fetching data:", error));
  };

  const handleDelete = (id: string) => {
    setData((prevData) => prevData.filter((job) => job.Id !== id));
  };

  const handleCreate = (newJob: Job) => {
    setData((prevData) => [newJob, ...prevData]);
  };

  const handleLogin = (token: string) => {
    setToken(token);
    setIsLoggedIn(true);
    fetchData();
  };

  const handleLogout = () => {
    setToken("");
    setIsLoggedIn(false);
    setData([]);
  };

  useEffect(() => {
    if (isLoggedIn) {
      fetchData(); // Fetch data immediately on mount
    }
  }, [isLoggedIn]);

  return (
    <div className="flex flex-col gap-4 p-4">
      {!isLoggedIn ? (
        <>
          <Login onLogin={handleLogin} />
          <Register /> {/* Add the Register component */}
        </>
      ) : (
        <>
          <div className="flex gap-2 justify-center">
            <DeleteAll onDeleteAll={fetchData} token={token} />
            <Button
              variant="outline"
              onClick={fetchData}
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            >
              Update jobs!
            </Button>
            <CreateJob onCreate={handleCreate} token={token} />
            <Button
              variant="outline"
              onClick={handleLogout}
              className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
            >
              Logout
            </Button>
          </div>

          {data.length > 0 && (
            <div className="grid grid-cols-3 gap-4">
              {data.map((job: Job, index: number) => (
                <motion.div
                  key={job.Id}
                  initial={{ opacity: 0, y: 0 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{
                    duration: 0.4,
                    delay: index * 0.1, // Creates a stagger effect
                    ease: "easeOut",
                  }}
                >
                  <JobCard job={job} onDelete={handleDelete} token={token} />
                </motion.div>
              ))}
            </div>
          )}
        </>
      )}
    </div>
  );
}

export default App;
