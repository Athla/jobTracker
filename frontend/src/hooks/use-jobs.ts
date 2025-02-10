import { useState, useEffect } from "react"; // Add useEffect
import { JobAPI } from "@/services/api";
import { Job, JobStatus } from "@/types/job";
import { useToast } from "./use-toast";
import { useAuth } from "@/context/AuthContext";

export function useJobs() {
  const [jobs, setJobs] = useState<Job[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const { toast } = useToast();
  const { logout } = useAuth();

  // Add useEffect to fetch jobs when component mounts
  useEffect(() => {
    fetchJobs();
  }, []); // Empty dependency array means this runs once when component mounts

  const updateJob = async (jobId: string, newStatus: JobStatus) => {
    const previousJobs = [...jobs];

    // Optimistic update
    setJobs(
      jobs.map((job) =>
        job.id === jobId ? { ...job, status: newStatus } : job
      )
    );

    try {
      const job = jobs.find((j) => j.id === jobId);
      if (!job) return;

      await JobAPI.updateStatus(jobId, {
        status: newStatus,
        version: job.version,
      });
    } catch (error) {
      // Revert on failure
      setJobs(previousJobs);
      toast({
        title: "Error",
        description: `Failed to update job status due: ${error}`,
        variant: "destructive",
      });
    }
  };

  const createJob = async (newJob: Partial<Job>) => {
    try {
      const createdJob = await JobAPI.create(newJob);
      setJobs((prev) => [...prev, createdJob]);
      toast({
        title: "Success",
        description: "Job created successfully",
      });
    } catch (error) {
      const message =
        error instanceof Error ? error.message : "Failed to create job";
      toast({
        title: "Error",
        description: message,
        variant: "destructive",
      });
      throw error;
    }
  };

  const fetchJobs = async () => {
    try {
      setIsLoading(true);
      const fetchedJobs = await JobAPI.getAll();
      setJobs(fetchedJobs || []); // Ensure we always have an array
      setError(null);
    } catch (e) {
      const error = e instanceof Error ? e : new Error("Failed to fetch jobs");
      setError(error);

      if (error.message === "Unauthorized") {
        logout();
        return;
      }

      toast({
        title: "Error",
        description: error.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  };

  return {
    jobs,
    isLoading,
    error,
    fetchJobs,
    updateJob,
    createJob,
  };
}
