import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Select } from "@/components/ui/select"; // You'll need to create this component
import { Card } from "@/components/ui/card";
import { Job } from "@/types/job";
import JobCard from "./JobCard";

type Status = "APPLIED" | "INTERVIEW" | "REJECTED" | "ACCEPTED" | "ALL";

interface JobListProps {
  jobs: Job[];
  onDelete: (id: string) => void;
}

function JobList({ jobs }: JobListProps) {
  const [searchTerm, setSearchTerm] = useState("");
  const [statusFilter, setStatusFilter] = useState<Status>("ALL");

  const filteredJobs = jobs.filter((job) => {
    const matchesSearch = job.name.toLowerCase().includes(
      searchTerm.toLowerCase(),
    );
    const matchesStatus = statusFilter === "ALL" || job.status === statusFilter;
    return matchesSearch && matchesStatus;
  });

  return (
    <div className="space-y-4">
      <div className="flex gap-4 mb-4">
        <Input
          type="text"
          placeholder="Search jobs..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="max-w-xs"
        />
        <Select
          value={statusFilter}
          onValueChange={(value: Status) => setStatusFilter(value)}
        >
          <option value="ALL">All Status</option>
          <option value="APPLIED">Applied</option>
          <option value="INTERVIEW">Interview</option>
          <option value="REJECTED">Rejected</option>
          <option value="ACCEPTED">Accepted</option>
        </Select>
      </div>

      <div className="grid gap-4">
        {filteredJobs.length > 0 ? (
          filteredJobs.map((job) => (
            <JobCard
              key={job.id}
              job={job}
            />
          ))
        ) : (
          <Card className="p-4 text-center text-gray-500">
            No jobs found matching your criteria
          </Card>
        )}
      </div>
    </div>
  );
}

export default JobList;
