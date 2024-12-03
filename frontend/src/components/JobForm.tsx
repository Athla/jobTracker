import { useState } from "react";
import JobCard, { Job } from "./CardJob";
import { Input } from "@/components/ui/input";
import { Select } from "@/components/ui/select"; // You'll need to create this component
import { Card } from "@/components/ui/card";

type Status = "APPLIED" | "INTERVIEW" | "REJECTED" | "ACCEPTED" | "ALL";

interface JobListProps {
  jobs: Job[];
  onDelete: (id: string) => void;
}

function JobList({ jobs, onDelete }: JobListProps) {
  const [searchTerm, setSearchTerm] = useState("");
  const [statusFilter, setStatusFilter] = useState<Status>("ALL");

  const filteredJobs = jobs.filter((job) => {
    const matchesSearch = job.Name.toLowerCase().includes(
      searchTerm.toLowerCase(),
    );
    const matchesStatus = statusFilter === "ALL" || job.Status === statusFilter;
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
              key={job.Id}
              job={job}
              onDelete={(id: string | number) => onDelete(String(id))}
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
