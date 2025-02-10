import { useJobs } from "@/hooks/use-jobs";
import { LoadingState } from "./LoadingState";
import Board from "./Board";
import { Button } from "./ui/button";
import { Toaster } from "./ui/toaster";
import { useAuth } from "@/context/AuthContext";
import CreateJobDialog from "./CreateJobDialog";

export default function Dashboard() {
  const { jobs, isLoading, error, updateJob, createJob } = useJobs();
  const { logout } = useAuth();

  if (isLoading) {
    return <LoadingState message="Loading jobs..." />;
  }

  if (error) {
    return (
      <div className="p-4 text-center">
        <p className="text-destructive">{error.message}</p>
        <Button onClick={() => window.location.reload()} className="mt-4">
          Retry
        </Button>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <header className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold">Job Tracker</h1>
          <div className="flex gap-4">
            <CreateJobDialog onJobCreated={createJob} />
            <Button onClick={() => logout()}>Logout</Button>
          </div>
        </header>
        {jobs && jobs.length === 0 ? (
          <div className="text-center py-12">
            <h2 className="text-xl font-semibold text-muted-foreground mb-4">
              No jobs yet
            </h2>
            <p className="text-muted-foreground mb-8">
              Click the "Create Job" button to add your first job application
            </p>
          </div>
        ) : (
          <div className="w-full overflow-x-auto">
            <Board jobs={jobs || []} onJobMove={updateJob} />
          </div>
        )}
      </div>
      <Toaster />
    </div>
  );
}
