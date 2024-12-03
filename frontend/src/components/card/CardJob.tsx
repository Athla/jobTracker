import React from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useState } from "react";

export interface Job {
  Id: string;
  Name: string;
  Source: string;
  Description: string;
  CreatedAt: string;
}

interface JobCardProps {
  job: Job;
  onDelete: (id: string | number) => void;
}

const JobCard: React.FC<JobCardProps> = ({ job, onDelete }) => {
  const [isDeleting, setIsDeleting] = useState(false);

  const handleDelete = async () => {
    if (!confirm("Are you sure you want to delete this job?")) return;

    setIsDeleting(true);
    try {
      const response = await fetch(`http://localhost:8080/api/jobs/${job.Id}`, {
        method: "DELETE",
      });
      if (!response.ok) {
        throw new Error("Failed to delete job");
      }
      onDelete(job.Id);
    } catch (error) {
      console.error("Error deleting job:", error);
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <Card className="relative">
      <CardHeader>
        <div className="flex justify-between items-start">
          <CardTitle>{job.Name}</CardTitle>
          <Button
            variant="ghost"
            size="sm"
            onClick={handleDelete}
            disabled={isDeleting}
          >
            {isDeleting ? "Deleting..." : "Delete"}
          </Button>
        </div>
        <CardDescription>
          {new Date(job.CreatedAt).toLocaleString()}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="space-y-2">
          <p className="text-sm font-medium">Source: {job.Source}</p>
          <p className="text-sm">{job.Description}</p>
        </div>
      </CardContent>
    </Card>
  );
};
export default JobCard;
