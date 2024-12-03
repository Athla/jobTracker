import React from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Trash2 } from "lucide-react";

import { useState } from "react";

export interface Job {
  Id: string;
  Name: string;
  Source: string;
  Description: string;
  CreatedAt: string;
  Status?: "APPLIED" | "INTERVIEW" | "REJECTED" | "ACCEPTED";
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

  const getStatusColor = (status?: string) => {
    switch (status) {
      case "APPLIED":
        return "bg-blue-100 text-blue-800";
      case "INTERVIEW":
        return "bg-yellow-100 text-yellow-800";
      case "REJECTED":
        return "bg-red-100 text-red-800";
      case "ACCEPTED":
        return "bg-green-100 text-green-800";
      default:
        return "bg-gray-100 text-gray-800";
    }
  };

  return (
    <Card className="relative">
      <CardHeader>
        <div className="flex justify-between items-start">
          <div>
            <CardTitle>{job.Name}</CardTitle>
            {job.Status && (
              <span
                className={`inline-block px-2 py-1 text-xs font-semibold rounded-full mt-2 ${getStatusColor(job.Status)}`}
              >
                {job.Status}
              </span>
            )}
          </div>
          <Button
            variant="ghost"
            size="icon"
            onClick={handleDelete}
            disabled={isDeleting}
            className="h-8 w-8 text-muted-foreground hover:text-destructive hover:bg-destructive/10"
          >
            <Trash2 className="h-4 w-4" />
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
